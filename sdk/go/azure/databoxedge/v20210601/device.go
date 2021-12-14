


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
	DataBoxEdgeDeviceStatus pulumi.StringPtrOutput            `pulumi:"dataBoxEdgeDeviceStatus"`
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
	Kind                    pulumi.StringPtrOutput            `pulumi:"kind"`
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
	DataBoxEdgeDeviceStatus *string           `pulumi:"dataBoxEdgeDeviceStatus"`
	DataResidency           *DataResidency    `pulumi:"dataResidency"`
	DeviceName              *string           `pulumi:"deviceName"`
	Etag                    *string           `pulumi:"etag"`
	Identity                *ResourceIdentity `pulumi:"identity"`
	Kind                    *string           `pulumi:"kind"`
	Location                *string           `pulumi:"location"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	Sku                     *Sku              `pulumi:"sku"`
	Tags                    map[string]string `pulumi:"tags"`
}


type DeviceArgs struct {
	DataBoxEdgeDeviceStatus pulumi.StringPtrInput
	DataResidency           DataResidencyPtrInput
	DeviceName              pulumi.StringPtrInput
	Etag                    pulumi.StringPtrInput
	Identity                ResourceIdentityPtrInput
	Kind                    pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
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

func init() {
	pulumi.RegisterOutputType(DeviceOutput{})
}
