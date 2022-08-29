


package v20210501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Device struct {
	pulumi.CustomResourceState

	DeviceType        pulumi.StringOutput            `pulumi:"deviceType"`
	Location          pulumi.StringOutput            `pulumi:"location"`
	Name              pulumi.StringOutput            `pulumi:"name"`
	NetworkFunctions  SubResourceResponseArrayOutput `pulumi:"networkFunctions"`
	ProvisioningState pulumi.StringOutput            `pulumi:"provisioningState"`
	Status            pulumi.StringOutput            `pulumi:"status"`
	SystemData        SystemDataResponseOutput       `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput         `pulumi:"tags"`
	Type              pulumi.StringOutput            `pulumi:"type"`
}


func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceType == nil {
		return nil, errors.New("invalid value for required argument 'DeviceType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridnetwork:Device"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20200101preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20220101preview:Device"),
		},
	})
	opts = append(opts, aliases)
	var resource Device
	err := ctx.RegisterResource("azure-native:hybridnetwork/v20210501:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("azure-native:hybridnetwork/v20210501:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type deviceState struct {
}

type DeviceState struct {
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	DeviceName        *string           `pulumi:"deviceName"`
	DeviceType        string            `pulumi:"deviceType"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type DeviceArgs struct {
	DeviceName        pulumi.StringPtrInput
	DeviceType        pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}

type DeviceInput interface {
	pulumi.Input

	ToDeviceOutput() DeviceOutput
	ToDeviceOutputWithContext(ctx context.Context) DeviceOutput
}

func (*Device) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (i *Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i *Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

func (o DeviceOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceType }).(pulumi.StringOutput)
}

func (o DeviceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeviceOutput) NetworkFunctions() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *Device) SubResourceResponseArrayOutput { return v.NetworkFunctions }).(SubResourceResponseArrayOutput)
}

func (o DeviceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DeviceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o DeviceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Device) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DeviceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Device) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DeviceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeviceOutput{})
}
