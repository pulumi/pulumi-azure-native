


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadNetworkDhcp struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput `pulumi:"name"`
	Properties pulumi.AnyOutput    `pulumi:"properties"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewWorkloadNetworkDhcp(ctx *pulumi.Context,
	name string, args *WorkloadNetworkDhcpArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkDhcp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:avs/v20211201:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20200717preview:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210101preview:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210601:WorkloadNetworkDhcp"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkDhcp
	err := ctx.RegisterResource("azure-native:avs/v20211201:WorkloadNetworkDhcp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadNetworkDhcp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkDhcpState, opts ...pulumi.ResourceOption) (*WorkloadNetworkDhcp, error) {
	var resource WorkloadNetworkDhcp
	err := ctx.ReadResource("azure-native:avs/v20211201:WorkloadNetworkDhcp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workloadNetworkDhcpState struct {
}

type WorkloadNetworkDhcpState struct {
}

func (WorkloadNetworkDhcpState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDhcpState)(nil)).Elem()
}

type workloadNetworkDhcpArgs struct {
	DhcpId            *string     `pulumi:"dhcpId"`
	PrivateCloudName  string      `pulumi:"privateCloudName"`
	Properties        interface{} `pulumi:"properties"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
}


type WorkloadNetworkDhcpArgs struct {
	DhcpId            pulumi.StringPtrInput
	PrivateCloudName  pulumi.StringInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
}

func (WorkloadNetworkDhcpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDhcpArgs)(nil)).Elem()
}

type WorkloadNetworkDhcpInput interface {
	pulumi.Input

	ToWorkloadNetworkDhcpOutput() WorkloadNetworkDhcpOutput
	ToWorkloadNetworkDhcpOutputWithContext(ctx context.Context) WorkloadNetworkDhcpOutput
}

func (*WorkloadNetworkDhcp) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDhcp)(nil))
}

func (i *WorkloadNetworkDhcp) ToWorkloadNetworkDhcpOutput() WorkloadNetworkDhcpOutput {
	return i.ToWorkloadNetworkDhcpOutputWithContext(context.Background())
}

func (i *WorkloadNetworkDhcp) ToWorkloadNetworkDhcpOutputWithContext(ctx context.Context) WorkloadNetworkDhcpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDhcpOutput)
}

type WorkloadNetworkDhcpOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDhcpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDhcp)(nil))
}

func (o WorkloadNetworkDhcpOutput) ToWorkloadNetworkDhcpOutput() WorkloadNetworkDhcpOutput {
	return o
}

func (o WorkloadNetworkDhcpOutput) ToWorkloadNetworkDhcpOutputWithContext(ctx context.Context) WorkloadNetworkDhcpOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkDhcpOutput{})
}
