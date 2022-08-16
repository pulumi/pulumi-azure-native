


package v20200501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Device struct {
	pulumi.CustomResourceState

	ConfiguredRoleTypes     pulumi.StringArrayOutput `pulumi:"configuredRoleTypes"`
	Culture                 pulumi.StringOutput      `pulumi:"culture"`
	DataBoxEdgeDeviceStatus pulumi.StringPtrOutput   `pulumi:"dataBoxEdgeDeviceStatus"`
	Description             pulumi.StringPtrOutput   `pulumi:"description"`
	DeviceHcsVersion        pulumi.StringOutput      `pulumi:"deviceHcsVersion"`
	DeviceLocalCapacity     pulumi.Float64Output     `pulumi:"deviceLocalCapacity"`
	DeviceModel             pulumi.StringOutput      `pulumi:"deviceModel"`
	DeviceSoftwareVersion   pulumi.StringOutput      `pulumi:"deviceSoftwareVersion"`
	DeviceType              pulumi.StringOutput      `pulumi:"deviceType"`
	Etag                    pulumi.StringPtrOutput   `pulumi:"etag"`
	FriendlyName            pulumi.StringPtrOutput   `pulumi:"friendlyName"`
	Location                pulumi.StringOutput      `pulumi:"location"`
	ModelDescription        pulumi.StringPtrOutput   `pulumi:"modelDescription"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	NodeCount               pulumi.IntOutput         `pulumi:"nodeCount"`
	SerialNumber            pulumi.StringOutput      `pulumi:"serialNumber"`
	Sku                     SkuResponsePtrOutput     `pulumi:"sku"`
	Tags                    pulumi.StringMapOutput   `pulumi:"tags"`
	TimeZone                pulumi.StringOutput      `pulumi:"timeZone"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
}


func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:Device"),
		},
	})
	opts = append(opts, aliases)
	var resource Device
	err := ctx.RegisterResource("azure-native:databoxedge/v20200501preview:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("azure-native:databoxedge/v20200501preview:Device", name, id, state, &resource, opts...)
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
	DataBoxEdgeDeviceStatus *string           `pulumi:"dataBoxEdgeDeviceStatus"`
	Description             *string           `pulumi:"description"`
	DeviceName              *string           `pulumi:"deviceName"`
	FriendlyName            *string           `pulumi:"friendlyName"`
	Location                *string           `pulumi:"location"`
	ModelDescription        *string           `pulumi:"modelDescription"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	Sku                     *Sku              `pulumi:"sku"`
	Tags                    map[string]string `pulumi:"tags"`
}


type DeviceArgs struct {
	DataBoxEdgeDeviceStatus pulumi.StringPtrInput
	Description             pulumi.StringPtrInput
	DeviceName              pulumi.StringPtrInput
	FriendlyName            pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	ModelDescription        pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Sku                     SkuPtrInput
	Tags                    pulumi.StringMapInput
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

func (o DeviceOutput) ConfiguredRoleTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Device) pulumi.StringArrayOutput { return v.ConfiguredRoleTypes }).(pulumi.StringArrayOutput)
}

func (o DeviceOutput) Culture() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Culture }).(pulumi.StringOutput)
}

func (o DeviceOutput) DataBoxEdgeDeviceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.DataBoxEdgeDeviceStatus }).(pulumi.StringPtrOutput)
}

func (o DeviceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DeviceOutput) DeviceHcsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceHcsVersion }).(pulumi.StringOutput)
}

func (o DeviceOutput) DeviceLocalCapacity() pulumi.Float64Output {
	return o.ApplyT(func(v *Device) pulumi.Float64Output { return v.DeviceLocalCapacity }).(pulumi.Float64Output)
}

func (o DeviceOutput) DeviceModel() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceModel }).(pulumi.StringOutput)
}

func (o DeviceOutput) DeviceSoftwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceSoftwareVersion }).(pulumi.StringOutput)
}

func (o DeviceOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceType }).(pulumi.StringOutput)
}

func (o DeviceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DeviceOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o DeviceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DeviceOutput) ModelDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.ModelDescription }).(pulumi.StringPtrOutput)
}

func (o DeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeviceOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Device) pulumi.IntOutput { return v.NodeCount }).(pulumi.IntOutput)
}

func (o DeviceOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o DeviceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Device) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o DeviceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Device) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DeviceOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

func (o DeviceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeviceOutput{})
}
