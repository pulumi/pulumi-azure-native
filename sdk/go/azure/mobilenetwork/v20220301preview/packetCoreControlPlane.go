


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PacketCoreControlPlane struct {
	pulumi.CustomResourceState

	ControlPlaneAccessInterface InterfacePropertiesResponseOutput         `pulumi:"controlPlaneAccessInterface"`
	CoreNetworkTechnology       pulumi.StringPtrOutput                    `pulumi:"coreNetworkTechnology"`
	CreatedAt                   pulumi.StringPtrOutput                    `pulumi:"createdAt"`
	CreatedBy                   pulumi.StringPtrOutput                    `pulumi:"createdBy"`
	CreatedByType               pulumi.StringPtrOutput                    `pulumi:"createdByType"`
	CustomLocation              CustomLocationResourceIdResponsePtrOutput `pulumi:"customLocation"`
	LastModifiedAt              pulumi.StringPtrOutput                    `pulumi:"lastModifiedAt"`
	LastModifiedBy              pulumi.StringPtrOutput                    `pulumi:"lastModifiedBy"`
	LastModifiedByType          pulumi.StringPtrOutput                    `pulumi:"lastModifiedByType"`
	Location                    pulumi.StringOutput                       `pulumi:"location"`
	MobileNetwork               MobileNetworkResourceIdResponseOutput     `pulumi:"mobileNetwork"`
	Name                        pulumi.StringOutput                       `pulumi:"name"`
	ProvisioningState           pulumi.StringOutput                       `pulumi:"provisioningState"`
	SystemData                  SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags                        pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                        pulumi.StringOutput                       `pulumi:"type"`
	Version                     pulumi.StringPtrOutput                    `pulumi:"version"`
}


func NewPacketCoreControlPlane(ctx *pulumi.Context,
	name string, args *PacketCoreControlPlaneArgs, opts ...pulumi.ResourceOption) (*PacketCoreControlPlane, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneAccessInterface == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneAccessInterface'")
	}
	if args.MobileNetwork == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetwork'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:PacketCoreControlPlane"),
		},
	})
	opts = append(opts, aliases)
	var resource PacketCoreControlPlane
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220301preview:PacketCoreControlPlane", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPacketCoreControlPlane(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PacketCoreControlPlaneState, opts ...pulumi.ResourceOption) (*PacketCoreControlPlane, error) {
	var resource PacketCoreControlPlane
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220301preview:PacketCoreControlPlane", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type packetCoreControlPlaneState struct {
}

type PacketCoreControlPlaneState struct {
}

func (PacketCoreControlPlaneState) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCoreControlPlaneState)(nil)).Elem()
}

type packetCoreControlPlaneArgs struct {
	ControlPlaneAccessInterface InterfaceProperties       `pulumi:"controlPlaneAccessInterface"`
	CoreNetworkTechnology       *string                   `pulumi:"coreNetworkTechnology"`
	CreatedAt                   *string                   `pulumi:"createdAt"`
	CreatedBy                   *string                   `pulumi:"createdBy"`
	CreatedByType               *string                   `pulumi:"createdByType"`
	CustomLocation              *CustomLocationResourceId `pulumi:"customLocation"`
	LastModifiedAt              *string                   `pulumi:"lastModifiedAt"`
	LastModifiedBy              *string                   `pulumi:"lastModifiedBy"`
	LastModifiedByType          *string                   `pulumi:"lastModifiedByType"`
	Location                    *string                   `pulumi:"location"`
	MobileNetwork               MobileNetworkResourceId   `pulumi:"mobileNetwork"`
	PacketCoreControlPlaneName  *string                   `pulumi:"packetCoreControlPlaneName"`
	ResourceGroupName           string                    `pulumi:"resourceGroupName"`
	Tags                        map[string]string         `pulumi:"tags"`
	Version                     *string                   `pulumi:"version"`
}


type PacketCoreControlPlaneArgs struct {
	ControlPlaneAccessInterface InterfacePropertiesInput
	CoreNetworkTechnology       pulumi.StringPtrInput
	CreatedAt                   pulumi.StringPtrInput
	CreatedBy                   pulumi.StringPtrInput
	CreatedByType               pulumi.StringPtrInput
	CustomLocation              CustomLocationResourceIdPtrInput
	LastModifiedAt              pulumi.StringPtrInput
	LastModifiedBy              pulumi.StringPtrInput
	LastModifiedByType          pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	MobileNetwork               MobileNetworkResourceIdInput
	PacketCoreControlPlaneName  pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	Tags                        pulumi.StringMapInput
	Version                     pulumi.StringPtrInput
}

func (PacketCoreControlPlaneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCoreControlPlaneArgs)(nil)).Elem()
}

type PacketCoreControlPlaneInput interface {
	pulumi.Input

	ToPacketCoreControlPlaneOutput() PacketCoreControlPlaneOutput
	ToPacketCoreControlPlaneOutputWithContext(ctx context.Context) PacketCoreControlPlaneOutput
}

func (*PacketCoreControlPlane) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCoreControlPlane)(nil)).Elem()
}

func (i *PacketCoreControlPlane) ToPacketCoreControlPlaneOutput() PacketCoreControlPlaneOutput {
	return i.ToPacketCoreControlPlaneOutputWithContext(context.Background())
}

func (i *PacketCoreControlPlane) ToPacketCoreControlPlaneOutputWithContext(ctx context.Context) PacketCoreControlPlaneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCoreControlPlaneOutput)
}

type PacketCoreControlPlaneOutput struct{ *pulumi.OutputState }

func (PacketCoreControlPlaneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCoreControlPlane)(nil)).Elem()
}

func (o PacketCoreControlPlaneOutput) ToPacketCoreControlPlaneOutput() PacketCoreControlPlaneOutput {
	return o
}

func (o PacketCoreControlPlaneOutput) ToPacketCoreControlPlaneOutputWithContext(ctx context.Context) PacketCoreControlPlaneOutput {
	return o
}

func (o PacketCoreControlPlaneOutput) ControlPlaneAccessInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) InterfacePropertiesResponseOutput {
		return v.ControlPlaneAccessInterface
	}).(InterfacePropertiesResponseOutput)
}

func (o PacketCoreControlPlaneOutput) CoreNetworkTechnology() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.CoreNetworkTechnology }).(pulumi.StringPtrOutput)
}

func (o PacketCoreControlPlaneOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o PacketCoreControlPlaneOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o PacketCoreControlPlaneOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o PacketCoreControlPlaneOutput) CustomLocation() CustomLocationResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) CustomLocationResourceIdResponsePtrOutput { return v.CustomLocation }).(CustomLocationResourceIdResponsePtrOutput)
}

func (o PacketCoreControlPlaneOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o PacketCoreControlPlaneOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o PacketCoreControlPlaneOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o PacketCoreControlPlaneOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PacketCoreControlPlaneOutput) MobileNetwork() MobileNetworkResourceIdResponseOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) MobileNetworkResourceIdResponseOutput { return v.MobileNetwork }).(MobileNetworkResourceIdResponseOutput)
}

func (o PacketCoreControlPlaneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PacketCoreControlPlaneOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PacketCoreControlPlaneOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PacketCoreControlPlaneOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PacketCoreControlPlaneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PacketCoreControlPlaneOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PacketCoreControlPlaneOutput{})
}
