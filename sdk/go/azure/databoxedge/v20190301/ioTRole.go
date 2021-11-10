


package v20190301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IoTRole struct {
	pulumi.CustomResourceState

	HostPlatform         pulumi.StringOutput              `pulumi:"hostPlatform"`
	IoTDeviceDetails     IoTDeviceInfoResponseOutput      `pulumi:"ioTDeviceDetails"`
	IoTEdgeDeviceDetails IoTDeviceInfoResponseOutput      `pulumi:"ioTEdgeDeviceDetails"`
	Kind                 pulumi.StringOutput              `pulumi:"kind"`
	Name                 pulumi.StringOutput              `pulumi:"name"`
	RoleStatus           pulumi.StringOutput              `pulumi:"roleStatus"`
	ShareMappings        MountPointMapResponseArrayOutput `pulumi:"shareMappings"`
	Type                 pulumi.StringOutput              `pulumi:"type"`
}


func NewIoTRole(ctx *pulumi.Context,
	name string, args *IoTRoleArgs, opts ...pulumi.ResourceOption) (*IoTRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.HostPlatform == nil {
		return nil, errors.New("invalid value for required argument 'HostPlatform'")
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
	if args.RoleStatus == nil {
		return nil, errors.New("invalid value for required argument 'RoleStatus'")
	}
	args.Kind = pulumi.String("IOT")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:IoTRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:IoTRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:IoTRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:IoTRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:IoTRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:IoTRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:IoTRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:IoTRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:IoTRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:IoTRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:IoTRole"),
		},
	})
	opts = append(opts, aliases)
	var resource IoTRole
	err := ctx.RegisterResource("azure-native:databoxedge/v20190301:IoTRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIoTRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IoTRoleState, opts ...pulumi.ResourceOption) (*IoTRole, error) {
	var resource IoTRole
	err := ctx.ReadResource("azure-native:databoxedge/v20190301:IoTRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ioTRoleState struct {
}

type IoTRoleState struct {
}

func (IoTRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTRoleState)(nil)).Elem()
}

type ioTRoleArgs struct {
	DeviceName           string          `pulumi:"deviceName"`
	HostPlatform         string          `pulumi:"hostPlatform"`
	IoTDeviceDetails     IoTDeviceInfo   `pulumi:"ioTDeviceDetails"`
	IoTEdgeDeviceDetails IoTDeviceInfo   `pulumi:"ioTEdgeDeviceDetails"`
	Kind                 string          `pulumi:"kind"`
	Name                 *string         `pulumi:"name"`
	ResourceGroupName    string          `pulumi:"resourceGroupName"`
	RoleStatus           string          `pulumi:"roleStatus"`
	ShareMappings        []MountPointMap `pulumi:"shareMappings"`
}


type IoTRoleArgs struct {
	DeviceName           pulumi.StringInput
	HostPlatform         pulumi.StringInput
	IoTDeviceDetails     IoTDeviceInfoInput
	IoTEdgeDeviceDetails IoTDeviceInfoInput
	Kind                 pulumi.StringInput
	Name                 pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	RoleStatus           pulumi.StringInput
	ShareMappings        MountPointMapArrayInput
}

func (IoTRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTRoleArgs)(nil)).Elem()
}

type IoTRoleInput interface {
	pulumi.Input

	ToIoTRoleOutput() IoTRoleOutput
	ToIoTRoleOutputWithContext(ctx context.Context) IoTRoleOutput
}

func (*IoTRole) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTRole)(nil))
}

func (i *IoTRole) ToIoTRoleOutput() IoTRoleOutput {
	return i.ToIoTRoleOutputWithContext(context.Background())
}

func (i *IoTRole) ToIoTRoleOutputWithContext(ctx context.Context) IoTRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTRoleOutput)
}

type IoTRoleOutput struct{ *pulumi.OutputState }

func (IoTRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTRole)(nil))
}

func (o IoTRoleOutput) ToIoTRoleOutput() IoTRoleOutput {
	return o
}

func (o IoTRoleOutput) ToIoTRoleOutputWithContext(ctx context.Context) IoTRoleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IoTRoleOutput{})
}
