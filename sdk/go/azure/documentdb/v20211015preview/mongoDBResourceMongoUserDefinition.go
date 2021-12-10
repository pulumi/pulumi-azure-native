


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MongoDBResourceMongoUserDefinition struct {
	pulumi.CustomResourceState

	CustomData   pulumi.StringPtrOutput  `pulumi:"customData"`
	DatabaseName pulumi.StringPtrOutput  `pulumi:"databaseName"`
	Mechanisms   pulumi.StringPtrOutput  `pulumi:"mechanisms"`
	Name         pulumi.StringOutput     `pulumi:"name"`
	Password     pulumi.StringPtrOutput  `pulumi:"password"`
	Roles        RoleResponseArrayOutput `pulumi:"roles"`
	Type         pulumi.StringOutput     `pulumi:"type"`
	UserName     pulumi.StringPtrOutput  `pulumi:"userName"`
}


func NewMongoDBResourceMongoUserDefinition(ctx *pulumi.Context,
	name string, args *MongoDBResourceMongoUserDefinitionArgs, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoUserDefinition, error) {
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
			Type: pulumi.String("azure-native:documentdb:MongoDBResourceMongoUserDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource MongoDBResourceMongoUserDefinition
	err := ctx.RegisterResource("azure-native:documentdb/v20211015preview:MongoDBResourceMongoUserDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMongoDBResourceMongoUserDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoDBResourceMongoUserDefinitionState, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoUserDefinition, error) {
	var resource MongoDBResourceMongoUserDefinition
	err := ctx.ReadResource("azure-native:documentdb/v20211015preview:MongoDBResourceMongoUserDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mongoDBResourceMongoUserDefinitionState struct {
}

type MongoDBResourceMongoUserDefinitionState struct {
}

func (MongoDBResourceMongoUserDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoUserDefinitionState)(nil)).Elem()
}

type mongoDBResourceMongoUserDefinitionArgs struct {
	AccountName           string  `pulumi:"accountName"`
	CustomData            *string `pulumi:"customData"`
	DatabaseName          *string `pulumi:"databaseName"`
	Mechanisms            *string `pulumi:"mechanisms"`
	MongoUserDefinitionId *string `pulumi:"mongoUserDefinitionId"`
	Password              *string `pulumi:"password"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	Roles                 []Role  `pulumi:"roles"`
	UserName              *string `pulumi:"userName"`
}


type MongoDBResourceMongoUserDefinitionArgs struct {
	AccountName           pulumi.StringInput
	CustomData            pulumi.StringPtrInput
	DatabaseName          pulumi.StringPtrInput
	Mechanisms            pulumi.StringPtrInput
	MongoUserDefinitionId pulumi.StringPtrInput
	Password              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Roles                 RoleArrayInput
	UserName              pulumi.StringPtrInput
}

func (MongoDBResourceMongoUserDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoUserDefinitionArgs)(nil)).Elem()
}

type MongoDBResourceMongoUserDefinitionInput interface {
	pulumi.Input

	ToMongoDBResourceMongoUserDefinitionOutput() MongoDBResourceMongoUserDefinitionOutput
	ToMongoDBResourceMongoUserDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoUserDefinitionOutput
}

func (*MongoDBResourceMongoUserDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBResourceMongoUserDefinition)(nil))
}

func (i *MongoDBResourceMongoUserDefinition) ToMongoDBResourceMongoUserDefinitionOutput() MongoDBResourceMongoUserDefinitionOutput {
	return i.ToMongoDBResourceMongoUserDefinitionOutputWithContext(context.Background())
}

func (i *MongoDBResourceMongoUserDefinition) ToMongoDBResourceMongoUserDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoUserDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBResourceMongoUserDefinitionOutput)
}

type MongoDBResourceMongoUserDefinitionOutput struct{ *pulumi.OutputState }

func (MongoDBResourceMongoUserDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBResourceMongoUserDefinition)(nil))
}

func (o MongoDBResourceMongoUserDefinitionOutput) ToMongoDBResourceMongoUserDefinitionOutput() MongoDBResourceMongoUserDefinitionOutput {
	return o
}

func (o MongoDBResourceMongoUserDefinitionOutput) ToMongoDBResourceMongoUserDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoUserDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MongoDBResourceMongoUserDefinitionOutput{})
}
