


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IoTAddon struct {
	pulumi.CustomResourceState

	HostPlatform         pulumi.StringOutput         `pulumi:"hostPlatform"`
	HostPlatformType     pulumi.StringOutput         `pulumi:"hostPlatformType"`
	IoTDeviceDetails     IoTDeviceInfoResponseOutput `pulumi:"ioTDeviceDetails"`
	IoTEdgeDeviceDetails IoTDeviceInfoResponseOutput `pulumi:"ioTEdgeDeviceDetails"`
	Kind                 pulumi.StringOutput         `pulumi:"kind"`
	Name                 pulumi.StringOutput         `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput         `pulumi:"provisioningState"`
	SystemData           SystemDataResponseOutput    `pulumi:"systemData"`
	Type                 pulumi.StringOutput         `pulumi:"type"`
	Version              pulumi.StringOutput         `pulumi:"version"`
}


func NewIoTAddon(ctx *pulumi.Context,
	name string, args *IoTAddonArgs, opts ...pulumi.ResourceOption) (*IoTAddon, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.IoTDeviceDetails == nil {
		return nil, errors.New("invalid value for required argument 'IoTDeviceDetails'")
	}
	if args.IoTEdgeDeviceDetails == nil {
		return nil, errors.New("invalid value for required argument 'IoTEdgeDeviceDetails'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	args.Kind = pulumi.String("IotEdge")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:IoTAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:IoTAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:IoTAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:IoTAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:IoTAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:IoTAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:IoTAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:IoTAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:IoTAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:IoTAddon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:IoTAddon"),
		},
	})
	opts = append(opts, aliases)
	var resource IoTAddon
	err := ctx.RegisterResource("azure-native:databoxedge/v20201201:IoTAddon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIoTAddon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IoTAddonState, opts ...pulumi.ResourceOption) (*IoTAddon, error) {
	var resource IoTAddon
	err := ctx.ReadResource("azure-native:databoxedge/v20201201:IoTAddon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ioTAddonState struct {
}

type IoTAddonState struct {
}

func (IoTAddonState) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTAddonState)(nil)).Elem()
}

type ioTAddonArgs struct {
	AddonName            *string       `pulumi:"addonName"`
	DeviceName           string        `pulumi:"deviceName"`
	IoTDeviceDetails     IoTDeviceInfo `pulumi:"ioTDeviceDetails"`
	IoTEdgeDeviceDetails IoTDeviceInfo `pulumi:"ioTEdgeDeviceDetails"`
	Kind                 string        `pulumi:"kind"`
	ResourceGroupName    string        `pulumi:"resourceGroupName"`
	RoleName             string        `pulumi:"roleName"`
}


type IoTAddonArgs struct {
	AddonName            pulumi.StringPtrInput
	DeviceName           pulumi.StringInput
	IoTDeviceDetails     IoTDeviceInfoInput
	IoTEdgeDeviceDetails IoTDeviceInfoInput
	Kind                 pulumi.StringInput
	ResourceGroupName    pulumi.StringInput
	RoleName             pulumi.StringInput
}

func (IoTAddonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTAddonArgs)(nil)).Elem()
}

type IoTAddonInput interface {
	pulumi.Input

	ToIoTAddonOutput() IoTAddonOutput
	ToIoTAddonOutputWithContext(ctx context.Context) IoTAddonOutput
}

func (*IoTAddon) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTAddon)(nil)).Elem()
}

func (i *IoTAddon) ToIoTAddonOutput() IoTAddonOutput {
	return i.ToIoTAddonOutputWithContext(context.Background())
}

func (i *IoTAddon) ToIoTAddonOutputWithContext(ctx context.Context) IoTAddonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTAddonOutput)
}

type IoTAddonOutput struct{ *pulumi.OutputState }

func (IoTAddonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTAddon)(nil)).Elem()
}

func (o IoTAddonOutput) ToIoTAddonOutput() IoTAddonOutput {
	return o
}

func (o IoTAddonOutput) ToIoTAddonOutputWithContext(ctx context.Context) IoTAddonOutput {
	return o
}

func (o IoTAddonOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTAddon) pulumi.StringOutput { return v.HostPlatform }).(pulumi.StringOutput)
}

func (o IoTAddonOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTAddon) pulumi.StringOutput { return v.HostPlatformType }).(pulumi.StringOutput)
}

func (o IoTAddonOutput) IoTDeviceDetails() IoTDeviceInfoResponseOutput {
	return o.ApplyT(func(v *IoTAddon) IoTDeviceInfoResponseOutput { return v.IoTDeviceDetails }).(IoTDeviceInfoResponseOutput)
}

func (o IoTAddonOutput) IoTEdgeDeviceDetails() IoTDeviceInfoResponseOutput {
	return o.ApplyT(func(v *IoTAddon) IoTDeviceInfoResponseOutput { return v.IoTEdgeDeviceDetails }).(IoTDeviceInfoResponseOutput)
}

func (o IoTAddonOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTAddon) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o IoTAddonOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTAddon) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IoTAddonOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTAddon) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o IoTAddonOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IoTAddon) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o IoTAddonOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTAddon) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o IoTAddonOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTAddon) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IoTAddonOutput{})
}
