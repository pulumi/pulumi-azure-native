


package v20200801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PacketCapture struct {
	pulumi.CustomResourceState

	BytesToCapturePerPacket pulumi.Float64PtrOutput                    `pulumi:"bytesToCapturePerPacket"`
	Etag                    pulumi.StringOutput                        `pulumi:"etag"`
	Filters                 PacketCaptureFilterResponseArrayOutput     `pulumi:"filters"`
	Name                    pulumi.StringOutput                        `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput                        `pulumi:"provisioningState"`
	StorageLocation         PacketCaptureStorageLocationResponseOutput `pulumi:"storageLocation"`
	Target                  pulumi.StringOutput                        `pulumi:"target"`
	TimeLimitInSeconds      pulumi.IntPtrOutput                        `pulumi:"timeLimitInSeconds"`
	TotalBytesPerSession    pulumi.Float64PtrOutput                    `pulumi:"totalBytesPerSession"`
}


func NewPacketCapture(ctx *pulumi.Context,
	name string, args *PacketCaptureArgs, opts ...pulumi.ResourceOption) (*PacketCapture, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkWatcherName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkWatcherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageLocation == nil {
		return nil, errors.New("invalid value for required argument 'StorageLocation'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	if isZero(args.BytesToCapturePerPacket) {
		args.BytesToCapturePerPacket = pulumi.Float64Ptr(0.0)
	}
	if isZero(args.TimeLimitInSeconds) {
		args.TimeLimitInSeconds = pulumi.IntPtr(18000)
	}
	if isZero(args.TotalBytesPerSession) {
		args.TotalBytesPerSession = pulumi.Float64Ptr(1073741824.0)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:PacketCapture"),
		},
	})
	opts = append(opts, aliases)
	var resource PacketCapture
	err := ctx.RegisterResource("azure-native:network/v20200801:PacketCapture", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPacketCapture(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PacketCaptureState, opts ...pulumi.ResourceOption) (*PacketCapture, error) {
	var resource PacketCapture
	err := ctx.ReadResource("azure-native:network/v20200801:PacketCapture", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type packetCaptureState struct {
}

type PacketCaptureState struct {
}

func (PacketCaptureState) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCaptureState)(nil)).Elem()
}

type packetCaptureArgs struct {
	BytesToCapturePerPacket *float64                     `pulumi:"bytesToCapturePerPacket"`
	Filters                 []PacketCaptureFilter        `pulumi:"filters"`
	NetworkWatcherName      string                       `pulumi:"networkWatcherName"`
	PacketCaptureName       *string                      `pulumi:"packetCaptureName"`
	ResourceGroupName       string                       `pulumi:"resourceGroupName"`
	StorageLocation         PacketCaptureStorageLocation `pulumi:"storageLocation"`
	Target                  string                       `pulumi:"target"`
	TimeLimitInSeconds      *int                         `pulumi:"timeLimitInSeconds"`
	TotalBytesPerSession    *float64                     `pulumi:"totalBytesPerSession"`
}


type PacketCaptureArgs struct {
	BytesToCapturePerPacket pulumi.Float64PtrInput
	Filters                 PacketCaptureFilterArrayInput
	NetworkWatcherName      pulumi.StringInput
	PacketCaptureName       pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	StorageLocation         PacketCaptureStorageLocationInput
	Target                  pulumi.StringInput
	TimeLimitInSeconds      pulumi.IntPtrInput
	TotalBytesPerSession    pulumi.Float64PtrInput
}

func (PacketCaptureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCaptureArgs)(nil)).Elem()
}

type PacketCaptureInput interface {
	pulumi.Input

	ToPacketCaptureOutput() PacketCaptureOutput
	ToPacketCaptureOutputWithContext(ctx context.Context) PacketCaptureOutput
}

func (*PacketCapture) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCapture)(nil)).Elem()
}

func (i *PacketCapture) ToPacketCaptureOutput() PacketCaptureOutput {
	return i.ToPacketCaptureOutputWithContext(context.Background())
}

func (i *PacketCapture) ToPacketCaptureOutputWithContext(ctx context.Context) PacketCaptureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureOutput)
}

type PacketCaptureOutput struct{ *pulumi.OutputState }

func (PacketCaptureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCapture)(nil)).Elem()
}

func (o PacketCaptureOutput) ToPacketCaptureOutput() PacketCaptureOutput {
	return o
}

func (o PacketCaptureOutput) ToPacketCaptureOutputWithContext(ctx context.Context) PacketCaptureOutput {
	return o
}

func (o PacketCaptureOutput) BytesToCapturePerPacket() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PacketCapture) pulumi.Float64PtrOutput { return v.BytesToCapturePerPacket }).(pulumi.Float64PtrOutput)
}

func (o PacketCaptureOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCapture) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PacketCaptureOutput) Filters() PacketCaptureFilterResponseArrayOutput {
	return o.ApplyT(func(v *PacketCapture) PacketCaptureFilterResponseArrayOutput { return v.Filters }).(PacketCaptureFilterResponseArrayOutput)
}

func (o PacketCaptureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCapture) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PacketCaptureOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCapture) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PacketCaptureOutput) StorageLocation() PacketCaptureStorageLocationResponseOutput {
	return o.ApplyT(func(v *PacketCapture) PacketCaptureStorageLocationResponseOutput { return v.StorageLocation }).(PacketCaptureStorageLocationResponseOutput)
}

func (o PacketCaptureOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCapture) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

func (o PacketCaptureOutput) TimeLimitInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PacketCapture) pulumi.IntPtrOutput { return v.TimeLimitInSeconds }).(pulumi.IntPtrOutput)
}

func (o PacketCaptureOutput) TotalBytesPerSession() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PacketCapture) pulumi.Float64PtrOutput { return v.TotalBytesPerSession }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PacketCaptureOutput{})
}
