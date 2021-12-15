


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CassandraDataCenter struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties DataCenterResourceResponsePropertiesOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewCassandraDataCenter(ctx *pulumi.Context,
	name string, args *CassandraDataCenterArgs, opts ...pulumi.ResourceOption) (*CassandraDataCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:CassandraDataCenter"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:CassandraDataCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource CassandraDataCenter
	err := ctx.RegisterResource("azure-native:documentdb/v20210301preview:CassandraDataCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCassandraDataCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraDataCenterState, opts ...pulumi.ResourceOption) (*CassandraDataCenter, error) {
	var resource CassandraDataCenter
	err := ctx.ReadResource("azure-native:documentdb/v20210301preview:CassandraDataCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cassandraDataCenterState struct {
}

type CassandraDataCenterState struct {
}

func (CassandraDataCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraDataCenterState)(nil)).Elem()
}

type cassandraDataCenterArgs struct {
	ClusterName       string                        `pulumi:"clusterName"`
	DataCenterName    *string                       `pulumi:"dataCenterName"`
	Properties        *DataCenterResourceProperties `pulumi:"properties"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
}


type CassandraDataCenterArgs struct {
	ClusterName       pulumi.StringInput
	DataCenterName    pulumi.StringPtrInput
	Properties        DataCenterResourcePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
}

func (CassandraDataCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraDataCenterArgs)(nil)).Elem()
}

type CassandraDataCenterInput interface {
	pulumi.Input

	ToCassandraDataCenterOutput() CassandraDataCenterOutput
	ToCassandraDataCenterOutputWithContext(ctx context.Context) CassandraDataCenterOutput
}

func (*CassandraDataCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraDataCenter)(nil)).Elem()
}

func (i *CassandraDataCenter) ToCassandraDataCenterOutput() CassandraDataCenterOutput {
	return i.ToCassandraDataCenterOutputWithContext(context.Background())
}

func (i *CassandraDataCenter) ToCassandraDataCenterOutputWithContext(ctx context.Context) CassandraDataCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraDataCenterOutput)
}

type CassandraDataCenterOutput struct{ *pulumi.OutputState }

func (CassandraDataCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraDataCenter)(nil)).Elem()
}

func (o CassandraDataCenterOutput) ToCassandraDataCenterOutput() CassandraDataCenterOutput {
	return o
}

func (o CassandraDataCenterOutput) ToCassandraDataCenterOutputWithContext(ctx context.Context) CassandraDataCenterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CassandraDataCenterOutput{})
}
