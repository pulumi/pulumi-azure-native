


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadNetworkSegment struct {
	pulumi.CustomResourceState

	ConnectedGateway  pulumi.StringPtrOutput                           `pulumi:"connectedGateway"`
	DisplayName       pulumi.StringPtrOutput                           `pulumi:"displayName"`
	Name              pulumi.StringOutput                              `pulumi:"name"`
	PortVif           WorkloadNetworkSegmentPortVifResponseArrayOutput `pulumi:"portVif"`
	ProvisioningState pulumi.StringOutput                              `pulumi:"provisioningState"`
	Revision          pulumi.Float64PtrOutput                          `pulumi:"revision"`
	Status            pulumi.StringOutput                              `pulumi:"status"`
	Subnet            WorkloadNetworkSegmentSubnetResponsePtrOutput    `pulumi:"subnet"`
	Type              pulumi.StringOutput                              `pulumi:"type"`
}


func NewWorkloadNetworkSegment(ctx *pulumi.Context,
	name string, args *WorkloadNetworkSegmentArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkSegment, error) {
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
			Type: pulumi.String("azure-native:avs:WorkloadNetworkSegment"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkSegment"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkSegment"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkSegment"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkSegment
	err := ctx.RegisterResource("azure-native:avs/v20211201:WorkloadNetworkSegment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadNetworkSegment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkSegmentState, opts ...pulumi.ResourceOption) (*WorkloadNetworkSegment, error) {
	var resource WorkloadNetworkSegment
	err := ctx.ReadResource("azure-native:avs/v20211201:WorkloadNetworkSegment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workloadNetworkSegmentState struct {
}

type WorkloadNetworkSegmentState struct {
}

func (WorkloadNetworkSegmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkSegmentState)(nil)).Elem()
}

type workloadNetworkSegmentArgs struct {
	ConnectedGateway  *string                       `pulumi:"connectedGateway"`
	DisplayName       *string                       `pulumi:"displayName"`
	PrivateCloudName  string                        `pulumi:"privateCloudName"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	Revision          *float64                      `pulumi:"revision"`
	SegmentId         *string                       `pulumi:"segmentId"`
	Subnet            *WorkloadNetworkSegmentSubnet `pulumi:"subnet"`
}


type WorkloadNetworkSegmentArgs struct {
	ConnectedGateway  pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	PrivateCloudName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Revision          pulumi.Float64PtrInput
	SegmentId         pulumi.StringPtrInput
	Subnet            WorkloadNetworkSegmentSubnetPtrInput
}

func (WorkloadNetworkSegmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkSegmentArgs)(nil)).Elem()
}

type WorkloadNetworkSegmentInput interface {
	pulumi.Input

	ToWorkloadNetworkSegmentOutput() WorkloadNetworkSegmentOutput
	ToWorkloadNetworkSegmentOutputWithContext(ctx context.Context) WorkloadNetworkSegmentOutput
}

func (*WorkloadNetworkSegment) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegment)(nil)).Elem()
}

func (i *WorkloadNetworkSegment) ToWorkloadNetworkSegmentOutput() WorkloadNetworkSegmentOutput {
	return i.ToWorkloadNetworkSegmentOutputWithContext(context.Background())
}

func (i *WorkloadNetworkSegment) ToWorkloadNetworkSegmentOutputWithContext(ctx context.Context) WorkloadNetworkSegmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentOutput)
}

type WorkloadNetworkSegmentOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegment)(nil)).Elem()
}

func (o WorkloadNetworkSegmentOutput) ToWorkloadNetworkSegmentOutput() WorkloadNetworkSegmentOutput {
	return o
}

func (o WorkloadNetworkSegmentOutput) ToWorkloadNetworkSegmentOutputWithContext(ctx context.Context) WorkloadNetworkSegmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkSegmentOutput{})
}
