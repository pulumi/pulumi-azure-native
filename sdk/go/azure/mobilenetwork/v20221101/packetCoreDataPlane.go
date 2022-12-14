


package v20221101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PacketCoreDataPlane struct {
	pulumi.CustomResourceState

	Location                 pulumi.StringOutput               `pulumi:"location"`
	Name                     pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput               `pulumi:"provisioningState"`
	SystemData               SystemDataResponseOutput          `pulumi:"systemData"`
	Tags                     pulumi.StringMapOutput            `pulumi:"tags"`
	Type                     pulumi.StringOutput               `pulumi:"type"`
	UserPlaneAccessInterface InterfacePropertiesResponseOutput `pulumi:"userPlaneAccessInterface"`
}


func NewPacketCoreDataPlane(ctx *pulumi.Context,
	name string, args *PacketCoreDataPlaneArgs, opts ...pulumi.ResourceOption) (*PacketCoreDataPlane, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PacketCoreControlPlaneName == nil {
		return nil, errors.New("invalid value for required argument 'PacketCoreControlPlaneName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserPlaneAccessInterface == nil {
		return nil, errors.New("invalid value for required argument 'UserPlaneAccessInterface'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:PacketCoreDataPlane"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:PacketCoreDataPlane"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:PacketCoreDataPlane"),
		},
	})
	opts = append(opts, aliases)
	var resource PacketCoreDataPlane
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20221101:PacketCoreDataPlane", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPacketCoreDataPlane(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PacketCoreDataPlaneState, opts ...pulumi.ResourceOption) (*PacketCoreDataPlane, error) {
	var resource PacketCoreDataPlane
	err := ctx.ReadResource("azure-native:mobilenetwork/v20221101:PacketCoreDataPlane", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type packetCoreDataPlaneState struct {
}

type PacketCoreDataPlaneState struct {
}

func (PacketCoreDataPlaneState) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCoreDataPlaneState)(nil)).Elem()
}

type packetCoreDataPlaneArgs struct {
	Location                   *string             `pulumi:"location"`
	PacketCoreControlPlaneName string              `pulumi:"packetCoreControlPlaneName"`
	PacketCoreDataPlaneName    *string             `pulumi:"packetCoreDataPlaneName"`
	ResourceGroupName          string              `pulumi:"resourceGroupName"`
	Tags                       map[string]string   `pulumi:"tags"`
	UserPlaneAccessInterface   InterfaceProperties `pulumi:"userPlaneAccessInterface"`
}


type PacketCoreDataPlaneArgs struct {
	Location                   pulumi.StringPtrInput
	PacketCoreControlPlaneName pulumi.StringInput
	PacketCoreDataPlaneName    pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	Tags                       pulumi.StringMapInput
	UserPlaneAccessInterface   InterfacePropertiesInput
}

func (PacketCoreDataPlaneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCoreDataPlaneArgs)(nil)).Elem()
}

type PacketCoreDataPlaneInput interface {
	pulumi.Input

	ToPacketCoreDataPlaneOutput() PacketCoreDataPlaneOutput
	ToPacketCoreDataPlaneOutputWithContext(ctx context.Context) PacketCoreDataPlaneOutput
}

func (*PacketCoreDataPlane) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCoreDataPlane)(nil)).Elem()
}

func (i *PacketCoreDataPlane) ToPacketCoreDataPlaneOutput() PacketCoreDataPlaneOutput {
	return i.ToPacketCoreDataPlaneOutputWithContext(context.Background())
}

func (i *PacketCoreDataPlane) ToPacketCoreDataPlaneOutputWithContext(ctx context.Context) PacketCoreDataPlaneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCoreDataPlaneOutput)
}

type PacketCoreDataPlaneOutput struct{ *pulumi.OutputState }

func (PacketCoreDataPlaneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCoreDataPlane)(nil)).Elem()
}

func (o PacketCoreDataPlaneOutput) ToPacketCoreDataPlaneOutput() PacketCoreDataPlaneOutput {
	return o
}

func (o PacketCoreDataPlaneOutput) ToPacketCoreDataPlaneOutputWithContext(ctx context.Context) PacketCoreDataPlaneOutput {
	return o
}

func (o PacketCoreDataPlaneOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreDataPlane) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PacketCoreDataPlaneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreDataPlane) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PacketCoreDataPlaneOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreDataPlane) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PacketCoreDataPlaneOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PacketCoreDataPlane) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PacketCoreDataPlaneOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PacketCoreDataPlane) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PacketCoreDataPlaneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreDataPlane) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PacketCoreDataPlaneOutput) UserPlaneAccessInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v *PacketCoreDataPlane) InterfacePropertiesResponseOutput { return v.UserPlaneAccessInterface }).(InterfacePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(PacketCoreDataPlaneOutput{})
}
