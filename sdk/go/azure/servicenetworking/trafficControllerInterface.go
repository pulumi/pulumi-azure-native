


package servicenetworking

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TrafficControllerInterface struct {
	pulumi.CustomResourceState

	Associations           ResourceIDResponseArrayOutput `pulumi:"associations"`
	ConfigurationEndpoints pulumi.StringArrayOutput      `pulumi:"configurationEndpoints"`
	Frontends              ResourceIDResponseArrayOutput `pulumi:"frontends"`
	Location               pulumi.StringOutput           `pulumi:"location"`
	Name                   pulumi.StringOutput           `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput           `pulumi:"provisioningState"`
	SystemData             SystemDataResponseOutput      `pulumi:"systemData"`
	Tags                   pulumi.StringMapOutput        `pulumi:"tags"`
	Type                   pulumi.StringOutput           `pulumi:"type"`
}


func NewTrafficControllerInterface(ctx *pulumi.Context,
	name string, args *TrafficControllerInterfaceArgs, opts ...pulumi.ResourceOption) (*TrafficControllerInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicenetworking/v20221001preview:TrafficControllerInterface"),
		},
	})
	opts = append(opts, aliases)
	var resource TrafficControllerInterface
	err := ctx.RegisterResource("azure-native:servicenetworking:TrafficControllerInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTrafficControllerInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficControllerInterfaceState, opts ...pulumi.ResourceOption) (*TrafficControllerInterface, error) {
	var resource TrafficControllerInterface
	err := ctx.ReadResource("azure-native:servicenetworking:TrafficControllerInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type trafficControllerInterfaceState struct {
}

type TrafficControllerInterfaceState struct {
}

func (TrafficControllerInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficControllerInterfaceState)(nil)).Elem()
}

type trafficControllerInterfaceArgs struct {
	Location              *string           `pulumi:"location"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Tags                  map[string]string `pulumi:"tags"`
	TrafficControllerName *string           `pulumi:"trafficControllerName"`
}


type TrafficControllerInterfaceArgs struct {
	Location              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
	TrafficControllerName pulumi.StringPtrInput
}

func (TrafficControllerInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficControllerInterfaceArgs)(nil)).Elem()
}

type TrafficControllerInterfaceInput interface {
	pulumi.Input

	ToTrafficControllerInterfaceOutput() TrafficControllerInterfaceOutput
	ToTrafficControllerInterfaceOutputWithContext(ctx context.Context) TrafficControllerInterfaceOutput
}

func (*TrafficControllerInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficControllerInterface)(nil)).Elem()
}

func (i *TrafficControllerInterface) ToTrafficControllerInterfaceOutput() TrafficControllerInterfaceOutput {
	return i.ToTrafficControllerInterfaceOutputWithContext(context.Background())
}

func (i *TrafficControllerInterface) ToTrafficControllerInterfaceOutputWithContext(ctx context.Context) TrafficControllerInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficControllerInterfaceOutput)
}

type TrafficControllerInterfaceOutput struct{ *pulumi.OutputState }

func (TrafficControllerInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficControllerInterface)(nil)).Elem()
}

func (o TrafficControllerInterfaceOutput) ToTrafficControllerInterfaceOutput() TrafficControllerInterfaceOutput {
	return o
}

func (o TrafficControllerInterfaceOutput) ToTrafficControllerInterfaceOutputWithContext(ctx context.Context) TrafficControllerInterfaceOutput {
	return o
}

func (o TrafficControllerInterfaceOutput) Associations() ResourceIDResponseArrayOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) ResourceIDResponseArrayOutput { return v.Associations }).(ResourceIDResponseArrayOutput)
}

func (o TrafficControllerInterfaceOutput) ConfigurationEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringArrayOutput { return v.ConfigurationEndpoints }).(pulumi.StringArrayOutput)
}

func (o TrafficControllerInterfaceOutput) Frontends() ResourceIDResponseArrayOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) ResourceIDResponseArrayOutput { return v.Frontends }).(ResourceIDResponseArrayOutput)
}

func (o TrafficControllerInterfaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o TrafficControllerInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TrafficControllerInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o TrafficControllerInterfaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o TrafficControllerInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o TrafficControllerInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficControllerInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrafficControllerInterfaceOutput{})
}
