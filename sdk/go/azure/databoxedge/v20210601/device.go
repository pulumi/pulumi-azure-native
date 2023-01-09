


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Device struct {
	pulumi.CustomResourceState

	ConfiguredRoleTypes     pulumi.StringArrayOutput          `pulumi:"configuredRoleTypes"`
	Culture                 pulumi.StringOutput               `pulumi:"culture"`
	DataBoxEdgeDeviceStatus pulumi.StringOutput               `pulumi:"dataBoxEdgeDeviceStatus"`
	DataResidency           DataResidencyResponsePtrOutput    `pulumi:"dataResidency"`
	Description             pulumi.StringOutput               `pulumi:"description"`
	DeviceHcsVersion        pulumi.StringOutput               `pulumi:"deviceHcsVersion"`
	DeviceLocalCapacity     pulumi.Float64Output              `pulumi:"deviceLocalCapacity"`
	DeviceModel             pulumi.StringOutput               `pulumi:"deviceModel"`
	DeviceSoftwareVersion   pulumi.StringOutput               `pulumi:"deviceSoftwareVersion"`
	DeviceType              pulumi.StringOutput               `pulumi:"deviceType"`
	EdgeProfile             EdgeProfileResponseOutput         `pulumi:"edgeProfile"`
	Etag                    pulumi.StringPtrOutput            `pulumi:"etag"`
	FriendlyName            pulumi.StringOutput               `pulumi:"friendlyName"`
	Identity                ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind                    pulumi.StringOutput               `pulumi:"kind"`
	Location                pulumi.StringOutput               `pulumi:"location"`
	ModelDescription        pulumi.StringOutput               `pulumi:"modelDescription"`
	Name                    pulumi.StringOutput               `pulumi:"name"`
	NodeCount               pulumi.IntOutput                  `pulumi:"nodeCount"`
	ResourceMoveDetails     ResourceMoveDetailsResponseOutput `pulumi:"resourceMoveDetails"`
	SerialNumber            pulumi.StringOutput               `pulumi:"serialNumber"`
	Sku                     SkuResponsePtrOutput              `pulumi:"sku"`
	SystemData              SystemDataResponseOutput          `pulumi:"systemData"`
	Tags                    pulumi.StringMapOutput            `pulumi:"tags"`
	TimeZone                pulumi.StringOutput               `pulumi:"timeZone"`
	Type                    pulumi.StringOutput               `pulumi:"type"`
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
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:Device"),
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
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:Device"),
		},
	})
	opts = append(opts, aliases)
	var resource Device
	err := ctx.RegisterResource("azure-native:databoxedge/v20210601:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("azure-native:databoxedge/v20210601:Device", name, id, state, &resource, opts...)
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
	DataResidency     *DataResidency    `pulumi:"dataResidency"`
	DeviceName        *string           `pulumi:"deviceName"`
	Identity          *ResourceIdentity `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type DeviceArgs struct {
	DataResidency     DataResidencyPtrInput
	DeviceName        pulumi.StringPtrInput
	Identity          ResourceIdentityPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuPtrInput
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

func (o DeviceOutput) ConfiguredRoleTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Device) pulumi.StringArrayOutput { return v.ConfiguredRoleTypes }).(pulumi.StringArrayOutput)
}

func (o DeviceOutput) Culture() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Culture }).(pulumi.StringOutput)
}

func (o DeviceOutput) DataBoxEdgeDeviceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DataBoxEdgeDeviceStatus }).(pulumi.StringOutput)
}

func (o DeviceOutput) DataResidency() DataResidencyResponsePtrOutput {
	return o.ApplyT(func(v *Device) DataResidencyResponsePtrOutput { return v.DataResidency }).(DataResidencyResponsePtrOutput)
}

func (o DeviceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
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

func (o DeviceOutput) EdgeProfile() EdgeProfileResponseOutput {
	return o.ApplyT(func(v *Device) EdgeProfileResponseOutput { return v.EdgeProfile }).(EdgeProfileResponseOutput)
}

func (o DeviceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DeviceOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o DeviceOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Device) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o DeviceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DeviceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DeviceOutput) ModelDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.ModelDescription }).(pulumi.StringOutput)
}

func (o DeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeviceOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Device) pulumi.IntOutput { return v.NodeCount }).(pulumi.IntOutput)
}

func (o DeviceOutput) ResourceMoveDetails() ResourceMoveDetailsResponseOutput {
	return o.ApplyT(func(v *Device) ResourceMoveDetailsResponseOutput { return v.ResourceMoveDetails }).(ResourceMoveDetailsResponseOutput)
}

func (o DeviceOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o DeviceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Device) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o DeviceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Device) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
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
