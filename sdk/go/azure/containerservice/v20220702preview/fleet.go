


package v20220702preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Fleet struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput              `pulumi:"etag"`
	HubProfile        FleetHubProfileResponsePtrOutput `pulumi:"hubProfile"`
	Location          pulumi.StringOutput              `pulumi:"location"`
	Name              pulumi.StringOutput              `pulumi:"name"`
	ProvisioningState pulumi.StringOutput              `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput         `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput           `pulumi:"tags"`
	Type              pulumi.StringOutput              `pulumi:"type"`
}


func NewFleet(ctx *pulumi.Context,
	name string, args *FleetArgs, opts ...pulumi.ResourceOption) (*Fleet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:Fleet"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:Fleet"),
		},
	})
	opts = append(opts, aliases)
	var resource Fleet
	err := ctx.RegisterResource("azure-native:containerservice/v20220702preview:Fleet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFleet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetState, opts ...pulumi.ResourceOption) (*Fleet, error) {
	var resource Fleet
	err := ctx.ReadResource("azure-native:containerservice/v20220702preview:Fleet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fleetState struct {
}

type FleetState struct {
}

func (FleetState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetState)(nil)).Elem()
}

type fleetArgs struct {
	FleetName         *string           `pulumi:"fleetName"`
	HubProfile        *FleetHubProfile  `pulumi:"hubProfile"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type FleetArgs struct {
	FleetName         pulumi.StringPtrInput
	HubProfile        FleetHubProfilePtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (FleetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetArgs)(nil)).Elem()
}

type FleetInput interface {
	pulumi.Input

	ToFleetOutput() FleetOutput
	ToFleetOutputWithContext(ctx context.Context) FleetOutput
}

func (*Fleet) ElementType() reflect.Type {
	return reflect.TypeOf((**Fleet)(nil)).Elem()
}

func (i *Fleet) ToFleetOutput() FleetOutput {
	return i.ToFleetOutputWithContext(context.Background())
}

func (i *Fleet) ToFleetOutputWithContext(ctx context.Context) FleetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetOutput)
}

type FleetOutput struct{ *pulumi.OutputState }

func (FleetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fleet)(nil)).Elem()
}

func (o FleetOutput) ToFleetOutput() FleetOutput {
	return o
}

func (o FleetOutput) ToFleetOutputWithContext(ctx context.Context) FleetOutput {
	return o
}

func (o FleetOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o FleetOutput) HubProfile() FleetHubProfileResponsePtrOutput {
	return o.ApplyT(func(v *Fleet) FleetHubProfileResponsePtrOutput { return v.HubProfile }).(FleetHubProfileResponsePtrOutput)
}

func (o FleetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o FleetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FleetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FleetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Fleet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o FleetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FleetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FleetOutput{})
}
