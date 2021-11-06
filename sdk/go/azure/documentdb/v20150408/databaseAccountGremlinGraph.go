


package v20150408

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseAccountGremlinGraph struct {
	pulumi.CustomResourceState

	ConflictResolutionPolicy ConflictResolutionPolicyResponsePtrOutput `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               pulumi.IntPtrOutput                       `pulumi:"defaultTtl"`
	Etag                     pulumi.StringPtrOutput                    `pulumi:"etag"`
	IndexingPolicy           IndexingPolicyResponsePtrOutput           `pulumi:"indexingPolicy"`
	Location                 pulumi.StringPtrOutput                    `pulumi:"location"`
	Name                     pulumi.StringOutput                       `pulumi:"name"`
	PartitionKey             ContainerPartitionKeyResponsePtrOutput    `pulumi:"partitionKey"`
	Rid                      pulumi.StringPtrOutput                    `pulumi:"rid"`
	Tags                     pulumi.StringMapOutput                    `pulumi:"tags"`
	Ts                       pulumi.AnyOutput                          `pulumi:"ts"`
	Type                     pulumi.StringOutput                       `pulumi:"type"`
	UniqueKeyPolicy          UniqueKeyPolicyResponsePtrOutput          `pulumi:"uniqueKeyPolicy"`
}


func NewDatabaseAccountGremlinGraph(ctx *pulumi.Context,
	name string, args *DatabaseAccountGremlinGraphArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountGremlinGraph, error) {
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
			Type: pulumi.String("azure-nextgen:documentdb/v20150408:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150401:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20151106:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160319:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160331:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20190801:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20191212:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200301:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200401:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200901:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210115:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210301preview:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210315:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210401preview:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210415:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210515:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210615:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210701preview:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20211015:DatabaseAccountGremlinGraph"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountGremlinGraph
	err := ctx.RegisterResource("azure-native:documentdb/v20150408:DatabaseAccountGremlinGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccountGremlinGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountGremlinGraphState, opts ...pulumi.ResourceOption) (*DatabaseAccountGremlinGraph, error) {
	var resource DatabaseAccountGremlinGraph
	err := ctx.ReadResource("azure-native:documentdb/v20150408:DatabaseAccountGremlinGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAccountGremlinGraphState struct {
}

type DatabaseAccountGremlinGraphState struct {
}

func (DatabaseAccountGremlinGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountGremlinGraphState)(nil)).Elem()
}

type databaseAccountGremlinGraphArgs struct {
	AccountName       string               `pulumi:"accountName"`
	DatabaseName      string               `pulumi:"databaseName"`
	GraphName         *string              `pulumi:"graphName"`
	Options           map[string]string    `pulumi:"options"`
	Resource          GremlinGraphResource `pulumi:"resource"`
	ResourceGroupName string               `pulumi:"resourceGroupName"`
}


type DatabaseAccountGremlinGraphArgs struct {
	AccountName       pulumi.StringInput
	DatabaseName      pulumi.StringInput
	GraphName         pulumi.StringPtrInput
	Options           pulumi.StringMapInput
	Resource          GremlinGraphResourceInput
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountGremlinGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountGremlinGraphArgs)(nil)).Elem()
}

type DatabaseAccountGremlinGraphInput interface {
	pulumi.Input

	ToDatabaseAccountGremlinGraphOutput() DatabaseAccountGremlinGraphOutput
	ToDatabaseAccountGremlinGraphOutputWithContext(ctx context.Context) DatabaseAccountGremlinGraphOutput
}

func (*DatabaseAccountGremlinGraph) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountGremlinGraph)(nil))
}

func (i *DatabaseAccountGremlinGraph) ToDatabaseAccountGremlinGraphOutput() DatabaseAccountGremlinGraphOutput {
	return i.ToDatabaseAccountGremlinGraphOutputWithContext(context.Background())
}

func (i *DatabaseAccountGremlinGraph) ToDatabaseAccountGremlinGraphOutputWithContext(ctx context.Context) DatabaseAccountGremlinGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountGremlinGraphOutput)
}

type DatabaseAccountGremlinGraphOutput struct{ *pulumi.OutputState }

func (DatabaseAccountGremlinGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountGremlinGraph)(nil))
}

func (o DatabaseAccountGremlinGraphOutput) ToDatabaseAccountGremlinGraphOutput() DatabaseAccountGremlinGraphOutput {
	return o
}

func (o DatabaseAccountGremlinGraphOutput) ToDatabaseAccountGremlinGraphOutputWithContext(ctx context.Context) DatabaseAccountGremlinGraphOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountGremlinGraphOutput{})
}
