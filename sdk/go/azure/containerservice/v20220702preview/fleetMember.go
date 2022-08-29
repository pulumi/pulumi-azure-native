


package v20220702preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FleetMember struct {
	pulumi.CustomResourceState

	ClusterResourceId pulumi.StringPtrOutput   `pulumi:"clusterResourceId"`
	Etag              pulumi.StringOutput      `pulumi:"etag"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewFleetMember(ctx *pulumi.Context,
	name string, args *FleetMemberArgs, opts ...pulumi.ResourceOption) (*FleetMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FleetName == nil {
		return nil, errors.New("invalid value for required argument 'FleetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:FleetMember"),
		},
	})
	opts = append(opts, aliases)
	var resource FleetMember
	err := ctx.RegisterResource("azure-native:containerservice/v20220702preview:FleetMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFleetMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetMemberState, opts ...pulumi.ResourceOption) (*FleetMember, error) {
	var resource FleetMember
	err := ctx.ReadResource("azure-native:containerservice/v20220702preview:FleetMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fleetMemberState struct {
}

type FleetMemberState struct {
}

func (FleetMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetMemberState)(nil)).Elem()
}

type fleetMemberArgs struct {
	ClusterResourceId *string `pulumi:"clusterResourceId"`
	FleetMemberName   *string `pulumi:"fleetMemberName"`
	FleetName         string  `pulumi:"fleetName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type FleetMemberArgs struct {
	ClusterResourceId pulumi.StringPtrInput
	FleetMemberName   pulumi.StringPtrInput
	FleetName         pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (FleetMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetMemberArgs)(nil)).Elem()
}

type FleetMemberInput interface {
	pulumi.Input

	ToFleetMemberOutput() FleetMemberOutput
	ToFleetMemberOutputWithContext(ctx context.Context) FleetMemberOutput
}

func (*FleetMember) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetMember)(nil)).Elem()
}

func (i *FleetMember) ToFleetMemberOutput() FleetMemberOutput {
	return i.ToFleetMemberOutputWithContext(context.Background())
}

func (i *FleetMember) ToFleetMemberOutputWithContext(ctx context.Context) FleetMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetMemberOutput)
}

type FleetMemberOutput struct{ *pulumi.OutputState }

func (FleetMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetMember)(nil)).Elem()
}

func (o FleetMemberOutput) ToFleetMemberOutput() FleetMemberOutput {
	return o
}

func (o FleetMemberOutput) ToFleetMemberOutputWithContext(ctx context.Context) FleetMemberOutput {
	return o
}

func (o FleetMemberOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringPtrOutput { return v.ClusterResourceId }).(pulumi.StringPtrOutput)
}

func (o FleetMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o FleetMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FleetMemberOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FleetMemberOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FleetMember) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o FleetMemberOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMember) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FleetMemberOutput{})
}
