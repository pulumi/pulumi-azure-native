


package v20221101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PacketCoreControlPlane struct {
	pulumi.CustomResourceState

	ControlPlaneAccessInterface InterfacePropertiesResponseOutput                 `pulumi:"controlPlaneAccessInterface"`
	CoreNetworkTechnology       pulumi.StringPtrOutput                            `pulumi:"coreNetworkTechnology"`
	Identity                    ManagedServiceIdentityResponsePtrOutput           `pulumi:"identity"`
	Installation                InstallationResponseOutput                        `pulumi:"installation"`
	InteropSettings             pulumi.AnyOutput                                  `pulumi:"interopSettings"`
	LocalDiagnosticsAccess      LocalDiagnosticsAccessConfigurationResponseOutput `pulumi:"localDiagnosticsAccess"`
	Location                    pulumi.StringOutput                               `pulumi:"location"`
	Name                        pulumi.StringOutput                               `pulumi:"name"`
	Platform                    PlatformConfigurationResponseOutput               `pulumi:"platform"`
	ProvisioningState           pulumi.StringOutput                               `pulumi:"provisioningState"`
	RollbackVersion             pulumi.StringOutput                               `pulumi:"rollbackVersion"`
	Sites                       SiteResourceIdResponseArrayOutput                 `pulumi:"sites"`
	Sku                         pulumi.StringOutput                               `pulumi:"sku"`
	SystemData                  SystemDataResponseOutput                          `pulumi:"systemData"`
	Tags                        pulumi.StringMapOutput                            `pulumi:"tags"`
	Type                        pulumi.StringOutput                               `pulumi:"type"`
	UeMtu                       pulumi.IntPtrOutput                               `pulumi:"ueMtu"`
	Version                     pulumi.StringPtrOutput                            `pulumi:"version"`
}


func NewPacketCoreControlPlane(ctx *pulumi.Context,
	name string, args *PacketCoreControlPlaneArgs, opts ...pulumi.ResourceOption) (*PacketCoreControlPlane, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneAccessInterface == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneAccessInterface'")
	}
	if args.LocalDiagnosticsAccess == nil {
		return nil, errors.New("invalid value for required argument 'LocalDiagnosticsAccess'")
	}
	if args.Platform == nil {
		return nil, errors.New("invalid value for required argument 'Platform'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sites == nil {
		return nil, errors.New("invalid value for required argument 'Sites'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if isZero(args.UeMtu) {
		args.UeMtu = pulumi.IntPtr(1440)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:PacketCoreControlPlane"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:PacketCoreControlPlane"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:PacketCoreControlPlane"),
		},
	})
	opts = append(opts, aliases)
	var resource PacketCoreControlPlane
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20221101:PacketCoreControlPlane", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPacketCoreControlPlane(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PacketCoreControlPlaneState, opts ...pulumi.ResourceOption) (*PacketCoreControlPlane, error) {
	var resource PacketCoreControlPlane
	err := ctx.ReadResource("azure-native:mobilenetwork/v20221101:PacketCoreControlPlane", name, id, state, &resource, opts...)
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
	ControlPlaneAccessInterface InterfaceProperties                 `pulumi:"controlPlaneAccessInterface"`
	CoreNetworkTechnology       *string                             `pulumi:"coreNetworkTechnology"`
	Identity                    *ManagedServiceIdentity             `pulumi:"identity"`
	InteropSettings             interface{}                         `pulumi:"interopSettings"`
	LocalDiagnosticsAccess      LocalDiagnosticsAccessConfiguration `pulumi:"localDiagnosticsAccess"`
	Location                    *string                             `pulumi:"location"`
	PacketCoreControlPlaneName  *string                             `pulumi:"packetCoreControlPlaneName"`
	Platform                    PlatformConfiguration               `pulumi:"platform"`
	ResourceGroupName           string                              `pulumi:"resourceGroupName"`
	Sites                       []SiteResourceId                    `pulumi:"sites"`
	Sku                         string                              `pulumi:"sku"`
	Tags                        map[string]string                   `pulumi:"tags"`
	UeMtu                       *int                                `pulumi:"ueMtu"`
	Version                     *string                             `pulumi:"version"`
}


type PacketCoreControlPlaneArgs struct {
	ControlPlaneAccessInterface InterfacePropertiesInput
	CoreNetworkTechnology       pulumi.StringPtrInput
	Identity                    ManagedServiceIdentityPtrInput
	InteropSettings             pulumi.Input
	LocalDiagnosticsAccess      LocalDiagnosticsAccessConfigurationInput
	Location                    pulumi.StringPtrInput
	PacketCoreControlPlaneName  pulumi.StringPtrInput
	Platform                    PlatformConfigurationInput
	ResourceGroupName           pulumi.StringInput
	Sites                       SiteResourceIdArrayInput
	Sku                         pulumi.StringInput
	Tags                        pulumi.StringMapInput
	UeMtu                       pulumi.IntPtrInput
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

func (o PacketCoreControlPlaneOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o PacketCoreControlPlaneOutput) Installation() InstallationResponseOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) InstallationResponseOutput { return v.Installation }).(InstallationResponseOutput)
}

func (o PacketCoreControlPlaneOutput) InteropSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.AnyOutput { return v.InteropSettings }).(pulumi.AnyOutput)
}

func (o PacketCoreControlPlaneOutput) LocalDiagnosticsAccess() LocalDiagnosticsAccessConfigurationResponseOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) LocalDiagnosticsAccessConfigurationResponseOutput {
		return v.LocalDiagnosticsAccess
	}).(LocalDiagnosticsAccessConfigurationResponseOutput)
}

func (o PacketCoreControlPlaneOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PacketCoreControlPlaneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PacketCoreControlPlaneOutput) Platform() PlatformConfigurationResponseOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) PlatformConfigurationResponseOutput { return v.Platform }).(PlatformConfigurationResponseOutput)
}

func (o PacketCoreControlPlaneOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PacketCoreControlPlaneOutput) RollbackVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.RollbackVersion }).(pulumi.StringOutput)
}

func (o PacketCoreControlPlaneOutput) Sites() SiteResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) SiteResourceIdResponseArrayOutput { return v.Sites }).(SiteResourceIdResponseArrayOutput)
}

func (o PacketCoreControlPlaneOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.Sku }).(pulumi.StringOutput)
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

func (o PacketCoreControlPlaneOutput) UeMtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.IntPtrOutput { return v.UeMtu }).(pulumi.IntPtrOutput)
}

func (o PacketCoreControlPlaneOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PacketCoreControlPlaneOutput{})
}
