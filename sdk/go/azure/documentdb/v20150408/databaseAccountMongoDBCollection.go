


package v20150408

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DatabaseAccountMongoDBCollection struct {
	pulumi.CustomResourceState

	Indexes  MongoIndexResponseArrayOutput `pulumi:"indexes"`
	Location pulumi.StringPtrOutput        `pulumi:"location"`
	Name     pulumi.StringOutput           `pulumi:"name"`
	ShardKey pulumi.StringMapOutput        `pulumi:"shardKey"`
	Tags     pulumi.StringMapOutput        `pulumi:"tags"`
	Type     pulumi.StringOutput           `pulumi:"type"`
}


func NewDatabaseAccountMongoDBCollection(ctx *pulumi.Context,
	name string, args *DatabaseAccountMongoDBCollectionArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
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
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:DatabaseAccountMongoDBCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountMongoDBCollection
	err := ctx.RegisterResource("azure-native:documentdb/v20150408:DatabaseAccountMongoDBCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccountMongoDBCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountMongoDBCollectionState, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBCollection, error) {
	var resource DatabaseAccountMongoDBCollection
	err := ctx.ReadResource("azure-native:documentdb/v20150408:DatabaseAccountMongoDBCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAccountMongoDBCollectionState struct {
}

type DatabaseAccountMongoDBCollectionState struct {
}

func (DatabaseAccountMongoDBCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBCollectionState)(nil)).Elem()
}

type databaseAccountMongoDBCollectionArgs struct {
	AccountName       string                    `pulumi:"accountName"`
	CollectionName    *string                   `pulumi:"collectionName"`
	DatabaseName      string                    `pulumi:"databaseName"`
	Options           map[string]string         `pulumi:"options"`
	Resource          MongoDBCollectionResource `pulumi:"resource"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
}


type DatabaseAccountMongoDBCollectionArgs struct {
	AccountName       pulumi.StringInput
	CollectionName    pulumi.StringPtrInput
	DatabaseName      pulumi.StringInput
	Options           pulumi.StringMapInput
	Resource          MongoDBCollectionResourceInput
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountMongoDBCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBCollectionArgs)(nil)).Elem()
}

type DatabaseAccountMongoDBCollectionInput interface {
	pulumi.Input

	ToDatabaseAccountMongoDBCollectionOutput() DatabaseAccountMongoDBCollectionOutput
	ToDatabaseAccountMongoDBCollectionOutputWithContext(ctx context.Context) DatabaseAccountMongoDBCollectionOutput
}

func (*DatabaseAccountMongoDBCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountMongoDBCollection)(nil)).Elem()
}

func (i *DatabaseAccountMongoDBCollection) ToDatabaseAccountMongoDBCollectionOutput() DatabaseAccountMongoDBCollectionOutput {
	return i.ToDatabaseAccountMongoDBCollectionOutputWithContext(context.Background())
}

func (i *DatabaseAccountMongoDBCollection) ToDatabaseAccountMongoDBCollectionOutputWithContext(ctx context.Context) DatabaseAccountMongoDBCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountMongoDBCollectionOutput)
}

type DatabaseAccountMongoDBCollectionOutput struct{ *pulumi.OutputState }

func (DatabaseAccountMongoDBCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountMongoDBCollection)(nil)).Elem()
}

func (o DatabaseAccountMongoDBCollectionOutput) ToDatabaseAccountMongoDBCollectionOutput() DatabaseAccountMongoDBCollectionOutput {
	return o
}

func (o DatabaseAccountMongoDBCollectionOutput) ToDatabaseAccountMongoDBCollectionOutputWithContext(ctx context.Context) DatabaseAccountMongoDBCollectionOutput {
	return o
}

func (o DatabaseAccountMongoDBCollectionOutput) Indexes() MongoIndexResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBCollection) MongoIndexResponseArrayOutput { return v.Indexes }).(MongoIndexResponseArrayOutput)
}

func (o DatabaseAccountMongoDBCollectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBCollection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountMongoDBCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseAccountMongoDBCollectionOutput) ShardKey() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBCollection) pulumi.StringMapOutput { return v.ShardKey }).(pulumi.StringMapOutput)
}

func (o DatabaseAccountMongoDBCollectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBCollection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatabaseAccountMongoDBCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountMongoDBCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountMongoDBCollectionOutput{})
}
