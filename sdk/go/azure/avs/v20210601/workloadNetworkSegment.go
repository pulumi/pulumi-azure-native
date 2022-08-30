


package v20210601

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
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkSegment"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkSegment
	err := ctx.RegisterResource("azure-native:avs/v20210601:WorkloadNetworkSegment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadNetworkSegment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkSegmentState, opts ...pulumi.ResourceOption) (*WorkloadNetworkSegment, error) {
	var resource WorkloadNetworkSegment
	err := ctx.ReadResource("azure-native:avs/v20210601:WorkloadNetworkSegment", name, id, state, &resource, opts...)
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

func (o WorkloadNetworkSegmentOutput) ConnectedGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringPtrOutput { return v.ConnectedGateway }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkSegmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkSegmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkloadNetworkSegmentOutput) PortVif() WorkloadNetworkSegmentPortVifResponseArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) WorkloadNetworkSegmentPortVifResponseArrayOutput { return v.PortVif }).(WorkloadNetworkSegmentPortVifResponseArrayOutput)
}

func (o WorkloadNetworkSegmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkloadNetworkSegmentOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkSegmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o WorkloadNetworkSegmentOutput) Subnet() WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) WorkloadNetworkSegmentSubnetResponsePtrOutput { return v.Subnet }).(WorkloadNetworkSegmentSubnetResponsePtrOutput)
}

func (o WorkloadNetworkSegmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkSegmentOutput{})
}
