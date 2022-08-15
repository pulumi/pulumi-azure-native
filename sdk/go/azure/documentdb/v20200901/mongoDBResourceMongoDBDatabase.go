


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type MongoDBResourceMongoDBDatabase struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput                                `pulumi:"location"`
	Name     pulumi.StringOutput                                   `pulumi:"name"`
	Options  MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource MongoDBDatabaseGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                                `pulumi:"tags"`
	Type     pulumi.StringOutput                                   `pulumi:"type"`
}


func NewMongoDBResourceMongoDBDatabase(ctx *pulumi.Context,
	name string, args *MongoDBResourceMongoDBDatabaseArgs, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoDBDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:MongoDBResourceMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:MongoDBResourceMongoDBDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource MongoDBResourceMongoDBDatabase
	err := ctx.RegisterResource("azure-native:documentdb/v20200901:MongoDBResourceMongoDBDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMongoDBResourceMongoDBDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoDBResourceMongoDBDatabaseState, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoDBDatabase, error) {
	var resource MongoDBResourceMongoDBDatabase
	err := ctx.ReadResource("azure-native:documentdb/v20200901:MongoDBResourceMongoDBDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mongoDBResourceMongoDBDatabaseState struct {
}

type MongoDBResourceMongoDBDatabaseState struct {
}

func (MongoDBResourceMongoDBDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoDBDatabaseState)(nil)).Elem()
}

type mongoDBResourceMongoDBDatabaseArgs struct {
	AccountName       string                  `pulumi:"accountName"`
	DatabaseName      *string                 `pulumi:"databaseName"`
	Location          *string                 `pulumi:"location"`
	Options           *CreateUpdateOptions    `pulumi:"options"`
	Resource          MongoDBDatabaseResource `pulumi:"resource"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Tags              map[string]string       `pulumi:"tags"`
}


type MongoDBResourceMongoDBDatabaseArgs struct {
	AccountName       pulumi.StringInput
	DatabaseName      pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Options           CreateUpdateOptionsPtrInput
	Resource          MongoDBDatabaseResourceInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (MongoDBResourceMongoDBDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoDBDatabaseArgs)(nil)).Elem()
}

type MongoDBResourceMongoDBDatabaseInput interface {
	pulumi.Input

	ToMongoDBResourceMongoDBDatabaseOutput() MongoDBResourceMongoDBDatabaseOutput
	ToMongoDBResourceMongoDBDatabaseOutputWithContext(ctx context.Context) MongoDBResourceMongoDBDatabaseOutput
}

func (*MongoDBResourceMongoDBDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBResourceMongoDBDatabase)(nil)).Elem()
}

func (i *MongoDBResourceMongoDBDatabase) ToMongoDBResourceMongoDBDatabaseOutput() MongoDBResourceMongoDBDatabaseOutput {
	return i.ToMongoDBResourceMongoDBDatabaseOutputWithContext(context.Background())
}

func (i *MongoDBResourceMongoDBDatabase) ToMongoDBResourceMongoDBDatabaseOutputWithContext(ctx context.Context) MongoDBResourceMongoDBDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBResourceMongoDBDatabaseOutput)
}

type MongoDBResourceMongoDBDatabaseOutput struct{ *pulumi.OutputState }

func (MongoDBResourceMongoDBDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBResourceMongoDBDatabase)(nil)).Elem()
}

func (o MongoDBResourceMongoDBDatabaseOutput) ToMongoDBResourceMongoDBDatabaseOutput() MongoDBResourceMongoDBDatabaseOutput {
	return o
}

func (o MongoDBResourceMongoDBDatabaseOutput) ToMongoDBResourceMongoDBDatabaseOutputWithContext(ctx context.Context) MongoDBResourceMongoDBDatabaseOutput {
	return o
}

func (o MongoDBResourceMongoDBDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o MongoDBResourceMongoDBDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MongoDBResourceMongoDBDatabaseOutput) Options() MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBDatabase) MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
		return v.Options
	}).(MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o MongoDBResourceMongoDBDatabaseOutput) Resource() MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBDatabase) MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(MongoDBDatabaseGetPropertiesResponseResourcePtrOutput)
}

func (o MongoDBResourceMongoDBDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MongoDBResourceMongoDBDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoDBDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MongoDBResourceMongoDBDatabaseOutput{})
}
