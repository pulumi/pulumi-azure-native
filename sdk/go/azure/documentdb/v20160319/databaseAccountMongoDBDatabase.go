


package v20160319

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DatabaseAccountMongoDBDatabase struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewDatabaseAccountMongoDBDatabase(ctx *pulumi.Context,
	name string, args *DatabaseAccountMongoDBDatabaseArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:DatabaseAccountMongoDBDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:DatabaseAccountMongoDBDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountMongoDBDatabase
	err := ctx.RegisterResource("azure-native:documentdb/v20160319:DatabaseAccountMongoDBDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccountMongoDBDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountMongoDBDatabaseState, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBDatabase, error) {
	var resource DatabaseAccountMongoDBDatabase
	err := ctx.ReadResource("azure-native:documentdb/v20160319:DatabaseAccountMongoDBDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAccountMongoDBDatabaseState struct {
}

type DatabaseAccountMongoDBDatabaseState struct {
}

func (DatabaseAccountMongoDBDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBDatabaseState)(nil)).Elem()
}

type databaseAccountMongoDBDatabaseArgs struct {
	AccountName       string                  `pulumi:"accountName"`
	DatabaseName      *string                 `pulumi:"databaseName"`
	Options           map[string]string       `pulumi:"options"`
	Resource          MongoDBDatabaseResource `pulumi:"resource"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
}


type DatabaseAccountMongoDBDatabaseArgs struct {
	AccountName       pulumi.StringInput
	DatabaseName      pulumi.StringPtrInput
	Options           pulumi.StringMapInput
	Resource          MongoDBDatabaseResourceInput
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountMongoDBDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBDatabaseArgs)(nil)).Elem()
}

type DatabaseAccountMongoDBDatabaseInput interface {
	pulumi.Input

	ToDatabaseAccountMongoDBDatabaseOutput() DatabaseAccountMongoDBDatabaseOutput
	ToDatabaseAccountMongoDBDatabaseOutputWithContext(ctx context.Context) DatabaseAccountMongoDBDatabaseOutput
}

func (*DatabaseAccountMongoDBDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountMongoDBDatabase)(nil)).Elem()
}

func (i *DatabaseAccountMongoDBDatabase) ToDatabaseAccountMongoDBDatabaseOutput() DatabaseAccountMongoDBDatabaseOutput {
	return i.ToDatabaseAccountMongoDBDatabaseOutputWithContext(context.Background())
}

func (i *DatabaseAccountMongoDBDatabase) ToDatabaseAccountMongoDBDatabaseOutputWithContext(ctx context.Context) DatabaseAccountMongoDBDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountMongoDBDatabaseOutput)
}

type DatabaseAccountMongoDBDatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseAccountMongoDBDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountMongoDBDatabase)(nil)).Elem()
}

func (o DatabaseAccountMongoDBDatabaseOutput) ToDatabaseAccountMongoDBDatabaseOutput() DatabaseAccountMongoDBDatabaseOutput {
	return o
}

func (o DatabaseAccountMongoDBDatabaseOutput) ToDatabaseAccountMongoDBDatabaseOutputWithContext(ctx context.Context) DatabaseAccountMongoDBDatabaseOutput {
	return o
}

func (o DatabaseAccountMongoDBDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountMongoDBDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseAccountMongoDBDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatabaseAccountMongoDBDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountMongoDBDatabaseOutput{})
}
