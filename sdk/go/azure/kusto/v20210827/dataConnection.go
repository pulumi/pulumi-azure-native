


package v20210827

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DataConnection struct {
	pulumi.CustomResourceState

	Kind     pulumi.StringOutput    `pulumi:"kind"`
	Location pulumi.StringPtrOutput `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewDataConnection(ctx *pulumi.Context,
	name string, args *DataConnectionArgs, opts ...pulumi.ResourceOption) (*DataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:kusto/v20210827:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20190121:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20190515:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20190907:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20191109:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20200215:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20200614:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20200918:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20210101:DataConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource DataConnection
	err := ctx.RegisterResource("azure-native:kusto/v20210827:DataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataConnectionState, opts ...pulumi.ResourceOption) (*DataConnection, error) {
	var resource DataConnection
	err := ctx.ReadResource("azure-native:kusto/v20210827:DataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataConnectionState struct {
}

type DataConnectionState struct {
}

func (DataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectionState)(nil)).Elem()
}

type dataConnectionArgs struct {
	ClusterName        string  `pulumi:"clusterName"`
	DataConnectionName *string `pulumi:"dataConnectionName"`
	DatabaseName       string  `pulumi:"databaseName"`
	Kind               string  `pulumi:"kind"`
	Location           *string `pulumi:"location"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type DataConnectionArgs struct {
	ClusterName        pulumi.StringInput
	DataConnectionName pulumi.StringPtrInput
	DatabaseName       pulumi.StringInput
	Kind               pulumi.StringInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
}

func (DataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectionArgs)(nil)).Elem()
}

type DataConnectionInput interface {
	pulumi.Input

	ToDataConnectionOutput() DataConnectionOutput
	ToDataConnectionOutputWithContext(ctx context.Context) DataConnectionOutput
}

func (*DataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnection)(nil))
}

func (i *DataConnection) ToDataConnectionOutput() DataConnectionOutput {
	return i.ToDataConnectionOutputWithContext(context.Background())
}

func (i *DataConnection) ToDataConnectionOutputWithContext(ctx context.Context) DataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectionOutput)
}

type DataConnectionOutput struct{ *pulumi.OutputState }

func (DataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnection)(nil))
}

func (o DataConnectionOutput) ToDataConnectionOutput() DataConnectionOutput {
	return o
}

func (o DataConnectionOutput) ToDataConnectionOutputWithContext(ctx context.Context) DataConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataConnectionOutput{})
}
