


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type KustoPoolDataConnection struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Location   pulumi.StringPtrOutput   `pulumi:"location"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewKustoPoolDataConnection(ctx *pulumi.Context,
	name string, args *KustoPoolDataConnectionArgs, opts ...pulumi.ResourceOption) (*KustoPoolDataConnection, error) {
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
			Type: pulumi.String("azure-native:synapse:KustoPoolDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:KustoPoolDataConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoPoolDataConnection
	err := ctx.RegisterResource("azure-native:synapse/v20210601preview:KustoPoolDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoPoolDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoPoolDataConnectionState, opts ...pulumi.ResourceOption) (*KustoPoolDataConnection, error) {
	var resource KustoPoolDataConnection
	err := ctx.ReadResource("azure-native:synapse/v20210601preview:KustoPoolDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoPoolDataConnectionState struct {
}

type KustoPoolDataConnectionState struct {
}

func (KustoPoolDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolDataConnectionState)(nil)).Elem()
}

type kustoPoolDataConnectionArgs struct {
	DataConnectionName *string `pulumi:"dataConnectionName"`
	DatabaseName       string  `pulumi:"databaseName"`
	Kind               string  `pulumi:"kind"`
	KustoPoolName      string  `pulumi:"kustoPoolName"`
	Location           *string `pulumi:"location"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	WorkspaceName      string  `pulumi:"workspaceName"`
}


type KustoPoolDataConnectionArgs struct {
	DataConnectionName pulumi.StringPtrInput
	DatabaseName       pulumi.StringInput
	Kind               pulumi.StringInput
	KustoPoolName      pulumi.StringInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	WorkspaceName      pulumi.StringInput
}

func (KustoPoolDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolDataConnectionArgs)(nil)).Elem()
}

type KustoPoolDataConnectionInput interface {
	pulumi.Input

	ToKustoPoolDataConnectionOutput() KustoPoolDataConnectionOutput
	ToKustoPoolDataConnectionOutputWithContext(ctx context.Context) KustoPoolDataConnectionOutput
}

func (*KustoPoolDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolDataConnection)(nil)).Elem()
}

func (i *KustoPoolDataConnection) ToKustoPoolDataConnectionOutput() KustoPoolDataConnectionOutput {
	return i.ToKustoPoolDataConnectionOutputWithContext(context.Background())
}

func (i *KustoPoolDataConnection) ToKustoPoolDataConnectionOutputWithContext(ctx context.Context) KustoPoolDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoPoolDataConnectionOutput)
}

type KustoPoolDataConnectionOutput struct{ *pulumi.OutputState }

func (KustoPoolDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolDataConnection)(nil)).Elem()
}

func (o KustoPoolDataConnectionOutput) ToKustoPoolDataConnectionOutput() KustoPoolDataConnectionOutput {
	return o
}

func (o KustoPoolDataConnectionOutput) ToKustoPoolDataConnectionOutputWithContext(ctx context.Context) KustoPoolDataConnectionOutput {
	return o
}

func (o KustoPoolDataConnectionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDataConnection) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o KustoPoolDataConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KustoPoolDataConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o KustoPoolDataConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDataConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KustoPoolDataConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoPoolDataConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o KustoPoolDataConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDataConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoPoolDataConnectionOutput{})
}
