


package v20210315

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MongoDBResourceMongoDBCollection struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput                                  `pulumi:"location"`
	Name     pulumi.StringOutput                                     `pulumi:"name"`
	Options  MongoDBCollectionGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource MongoDBCollectionGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type     pulumi.StringOutput                                     `pulumi:"type"`
}


func NewMongoDBResourceMongoDBCollection(ctx *pulumi.Context,
	name string, args *MongoDBResourceMongoDBCollectionArgs, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoDBCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210315:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150401:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150408:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20151106:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160319:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160331:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20190801:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20191212:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200301:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200401:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200901:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210115:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210301preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210401preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210415:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210515:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210615:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210701preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20211015:MongoDBResourceMongoDBCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource MongoDBResourceMongoDBCollection
	err := ctx.RegisterResource("azure-native:documentdb/v20210315:MongoDBResourceMongoDBCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMongoDBResourceMongoDBCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoDBResourceMongoDBCollectionState, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoDBCollection, error) {
	var resource MongoDBResourceMongoDBCollection
	err := ctx.ReadResource("azure-native:documentdb/v20210315:MongoDBResourceMongoDBCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mongoDBResourceMongoDBCollectionState struct {
}

type MongoDBResourceMongoDBCollectionState struct {
}

func (MongoDBResourceMongoDBCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoDBCollectionState)(nil)).Elem()
}

type mongoDBResourceMongoDBCollectionArgs struct {
	AccountName       string                    `pulumi:"accountName"`
	CollectionName    *string                   `pulumi:"collectionName"`
	DatabaseName      string                    `pulumi:"databaseName"`
	Location          *string                   `pulumi:"location"`
	Options           *CreateUpdateOptions      `pulumi:"options"`
	Resource          MongoDBCollectionResource `pulumi:"resource"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type MongoDBResourceMongoDBCollectionArgs struct {
	AccountName       pulumi.StringInput
	CollectionName    pulumi.StringPtrInput
	DatabaseName      pulumi.StringInput
	Location          pulumi.StringPtrInput
	Options           CreateUpdateOptionsPtrInput
	Resource          MongoDBCollectionResourceInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (MongoDBResourceMongoDBCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoDBCollectionArgs)(nil)).Elem()
}

type MongoDBResourceMongoDBCollectionInput interface {
	pulumi.Input

	ToMongoDBResourceMongoDBCollectionOutput() MongoDBResourceMongoDBCollectionOutput
	ToMongoDBResourceMongoDBCollectionOutputWithContext(ctx context.Context) MongoDBResourceMongoDBCollectionOutput
}

func (*MongoDBResourceMongoDBCollection) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBResourceMongoDBCollection)(nil))
}

func (i *MongoDBResourceMongoDBCollection) ToMongoDBResourceMongoDBCollectionOutput() MongoDBResourceMongoDBCollectionOutput {
	return i.ToMongoDBResourceMongoDBCollectionOutputWithContext(context.Background())
}

func (i *MongoDBResourceMongoDBCollection) ToMongoDBResourceMongoDBCollectionOutputWithContext(ctx context.Context) MongoDBResourceMongoDBCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBResourceMongoDBCollectionOutput)
}

type MongoDBResourceMongoDBCollectionOutput struct{ *pulumi.OutputState }

func (MongoDBResourceMongoDBCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBResourceMongoDBCollection)(nil))
}

func (o MongoDBResourceMongoDBCollectionOutput) ToMongoDBResourceMongoDBCollectionOutput() MongoDBResourceMongoDBCollectionOutput {
	return o
}

func (o MongoDBResourceMongoDBCollectionOutput) ToMongoDBResourceMongoDBCollectionOutputWithContext(ctx context.Context) MongoDBResourceMongoDBCollectionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MongoDBResourceMongoDBCollectionInput)(nil)).Elem(), &MongoDBResourceMongoDBCollection{})
	pulumi.RegisterOutputType(MongoDBResourceMongoDBCollectionOutput{})
}
