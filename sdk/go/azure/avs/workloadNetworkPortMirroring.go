


package avs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadNetworkPortMirroring struct {
	pulumi.CustomResourceState

	Destination       pulumi.StringPtrOutput  `pulumi:"destination"`
	Direction         pulumi.StringPtrOutput  `pulumi:"direction"`
	DisplayName       pulumi.StringPtrOutput  `pulumi:"displayName"`
	Name              pulumi.StringOutput     `pulumi:"name"`
	ProvisioningState pulumi.StringOutput     `pulumi:"provisioningState"`
	Revision          pulumi.Float64PtrOutput `pulumi:"revision"`
	Source            pulumi.StringPtrOutput  `pulumi:"source"`
	Status            pulumi.StringOutput     `pulumi:"status"`
	Type              pulumi.StringOutput     `pulumi:"type"`
}


func NewWorkloadNetworkPortMirroring(ctx *pulumi.Context,
	name string, args *WorkloadNetworkPortMirroringArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkPortMirroring, error) {
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
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkPortMirroring"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkPortMirroring
	err := ctx.RegisterResource("azure-native:avs:WorkloadNetworkPortMirroring", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadNetworkPortMirroring(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkPortMirroringState, opts ...pulumi.ResourceOption) (*WorkloadNetworkPortMirroring, error) {
	var resource WorkloadNetworkPortMirroring
	err := ctx.ReadResource("azure-native:avs:WorkloadNetworkPortMirroring", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workloadNetworkPortMirroringState struct {
}

type WorkloadNetworkPortMirroringState struct {
}

func (WorkloadNetworkPortMirroringState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkPortMirroringState)(nil)).Elem()
}

type workloadNetworkPortMirroringArgs struct {
	Destination       *string  `pulumi:"destination"`
	Direction         *string  `pulumi:"direction"`
	DisplayName       *string  `pulumi:"displayName"`
	PortMirroringId   *string  `pulumi:"portMirroringId"`
	PrivateCloudName  string   `pulumi:"privateCloudName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Revision          *float64 `pulumi:"revision"`
	Source            *string  `pulumi:"source"`
}


type WorkloadNetworkPortMirroringArgs struct {
	Destination       pulumi.StringPtrInput
	Direction         pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	PortMirroringId   pulumi.StringPtrInput
	PrivateCloudName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Revision          pulumi.Float64PtrInput
	Source            pulumi.StringPtrInput
}

func (WorkloadNetworkPortMirroringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkPortMirroringArgs)(nil)).Elem()
}

type WorkloadNetworkPortMirroringInput interface {
	pulumi.Input

	ToWorkloadNetworkPortMirroringOutput() WorkloadNetworkPortMirroringOutput
	ToWorkloadNetworkPortMirroringOutputWithContext(ctx context.Context) WorkloadNetworkPortMirroringOutput
}

func (*WorkloadNetworkPortMirroring) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkPortMirroring)(nil)).Elem()
}

func (i *WorkloadNetworkPortMirroring) ToWorkloadNetworkPortMirroringOutput() WorkloadNetworkPortMirroringOutput {
	return i.ToWorkloadNetworkPortMirroringOutputWithContext(context.Background())
}

func (i *WorkloadNetworkPortMirroring) ToWorkloadNetworkPortMirroringOutputWithContext(ctx context.Context) WorkloadNetworkPortMirroringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkPortMirroringOutput)
}

type WorkloadNetworkPortMirroringOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkPortMirroringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkPortMirroring)(nil)).Elem()
}

func (o WorkloadNetworkPortMirroringOutput) ToWorkloadNetworkPortMirroringOutput() WorkloadNetworkPortMirroringOutput {
	return o
}

func (o WorkloadNetworkPortMirroringOutput) ToWorkloadNetworkPortMirroringOutputWithContext(ctx context.Context) WorkloadNetworkPortMirroringOutput {
	return o
}

func (o WorkloadNetworkPortMirroringOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkPortMirroringOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkPortMirroringOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkPortMirroringOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkloadNetworkPortMirroringOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkloadNetworkPortMirroringOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkPortMirroringOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkPortMirroringOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o WorkloadNetworkPortMirroringOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkPortMirroringOutput{})
}
