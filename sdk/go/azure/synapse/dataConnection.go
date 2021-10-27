


package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DataConnection struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Location   pulumi.StringPtrOutput   `pulumi:"location"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewDataConnection(ctx *pulumi.Context,
	name string, args *DataConnectionArgs, opts ...pulumi.ResourceOption) (*DataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KustoPoolName == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:synapse:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210401preview:DataConnection"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:DataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210601preview:DataConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource DataConnection
	err := ctx.RegisterResource("azure-native:synapse:DataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataConnectionState, opts ...pulumi.ResourceOption) (*DataConnection, error) {
	var resource DataConnection
	err := ctx.ReadResource("azure-native:synapse:DataConnection", name, id, state, &resource, opts...)
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
	DataConnectionName *string `pulumi:"dataConnectionName"`
	DatabaseName       string  `pulumi:"databaseName"`
	Kind               string  `pulumi:"kind"`
	KustoPoolName      string  `pulumi:"kustoPoolName"`
	Location           *string `pulumi:"location"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	WorkspaceName      string  `pulumi:"workspaceName"`
}


type DataConnectionArgs struct {
	DataConnectionName pulumi.StringPtrInput
	DatabaseName       pulumi.StringInput
	Kind               pulumi.StringInput
	KustoPoolName      pulumi.StringInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	WorkspaceName      pulumi.StringInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*DataConnectionInput)(nil)).Elem(), &DataConnection{})
	pulumi.RegisterOutputType(DataConnectionOutput{})
}
