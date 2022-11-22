


package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkVirtualAppliance struct {
	pulumi.CustomResourceState

	BootStrapConfigurationBlob pulumi.StringArrayOutput                         `pulumi:"bootStrapConfigurationBlob"`
	CloudInitConfigurationBlob pulumi.StringArrayOutput                         `pulumi:"cloudInitConfigurationBlob"`
	Etag                       pulumi.StringOutput                              `pulumi:"etag"`
	Identity                   ManagedServiceIdentityResponsePtrOutput          `pulumi:"identity"`
	Location                   pulumi.StringPtrOutput                           `pulumi:"location"`
	Name                       pulumi.StringOutput                              `pulumi:"name"`
	ProvisioningState          pulumi.StringOutput                              `pulumi:"provisioningState"`
	Sku                        VirtualApplianceSkuPropertiesResponsePtrOutput   `pulumi:"sku"`
	Tags                       pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                       pulumi.StringOutput                              `pulumi:"type"`
	VirtualApplianceAsn        pulumi.Float64PtrOutput                          `pulumi:"virtualApplianceAsn"`
	VirtualApplianceNics       VirtualApplianceNicPropertiesResponseArrayOutput `pulumi:"virtualApplianceNics"`
	VirtualHub                 SubResourceResponsePtrOutput                     `pulumi:"virtualHub"`
}


func NewNetworkVirtualAppliance(ctx *pulumi.Context,
	name string, args *NetworkVirtualApplianceArgs, opts ...pulumi.ResourceOption) (*NetworkVirtualAppliance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:NetworkVirtualAppliance"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkVirtualAppliance
	err := ctx.RegisterResource("azure-native:network/v20191201:NetworkVirtualAppliance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkVirtualAppliance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkVirtualApplianceState, opts ...pulumi.ResourceOption) (*NetworkVirtualAppliance, error) {
	var resource NetworkVirtualAppliance
	err := ctx.ReadResource("azure-native:network/v20191201:NetworkVirtualAppliance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkVirtualApplianceState struct {
}

type NetworkVirtualApplianceState struct {
}

func (NetworkVirtualApplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceState)(nil)).Elem()
}

type networkVirtualApplianceArgs struct {
	BootStrapConfigurationBlob  []string                       `pulumi:"bootStrapConfigurationBlob"`
	CloudInitConfigurationBlob  []string                       `pulumi:"cloudInitConfigurationBlob"`
	Id                          *string                        `pulumi:"id"`
	Identity                    *ManagedServiceIdentity        `pulumi:"identity"`
	Location                    *string                        `pulumi:"location"`
	NetworkVirtualApplianceName *string                        `pulumi:"networkVirtualApplianceName"`
	ResourceGroupName           string                         `pulumi:"resourceGroupName"`
	Sku                         *VirtualApplianceSkuProperties `pulumi:"sku"`
	Tags                        map[string]string              `pulumi:"tags"`
	VirtualApplianceAsn         *float64                       `pulumi:"virtualApplianceAsn"`
	VirtualHub                  *SubResource                   `pulumi:"virtualHub"`
}


type NetworkVirtualApplianceArgs struct {
	BootStrapConfigurationBlob  pulumi.StringArrayInput
	CloudInitConfigurationBlob  pulumi.StringArrayInput
	Id                          pulumi.StringPtrInput
	Identity                    ManagedServiceIdentityPtrInput
	Location                    pulumi.StringPtrInput
	NetworkVirtualApplianceName pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	Sku                         VirtualApplianceSkuPropertiesPtrInput
	Tags                        pulumi.StringMapInput
	VirtualApplianceAsn         pulumi.Float64PtrInput
	VirtualHub                  SubResourcePtrInput
}

func (NetworkVirtualApplianceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceArgs)(nil)).Elem()
}

type NetworkVirtualApplianceInput interface {
	pulumi.Input

	ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput
	ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput
}

func (*NetworkVirtualAppliance) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkVirtualAppliance)(nil)).Elem()
}

func (i *NetworkVirtualAppliance) ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput {
	return i.ToNetworkVirtualApplianceOutputWithContext(context.Background())
}

func (i *NetworkVirtualAppliance) ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkVirtualApplianceOutput)
}

type NetworkVirtualApplianceOutput struct{ *pulumi.OutputState }

func (NetworkVirtualApplianceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkVirtualAppliance)(nil)).Elem()
}

func (o NetworkVirtualApplianceOutput) ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput {
	return o
}

func (o NetworkVirtualApplianceOutput) ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput {
	return o
}

func (o NetworkVirtualApplianceOutput) BootStrapConfigurationBlob() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringArrayOutput { return v.BootStrapConfigurationBlob }).(pulumi.StringArrayOutput)
}

func (o NetworkVirtualApplianceOutput) CloudInitConfigurationBlob() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringArrayOutput { return v.CloudInitConfigurationBlob }).(pulumi.StringArrayOutput)
}

func (o NetworkVirtualApplianceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NetworkVirtualApplianceOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o NetworkVirtualApplianceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkVirtualApplianceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkVirtualApplianceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NetworkVirtualApplianceOutput) Sku() VirtualApplianceSkuPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) VirtualApplianceSkuPropertiesResponsePtrOutput { return v.Sku }).(VirtualApplianceSkuPropertiesResponsePtrOutput)
}

func (o NetworkVirtualApplianceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkVirtualApplianceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NetworkVirtualApplianceOutput) VirtualApplianceAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.Float64PtrOutput { return v.VirtualApplianceAsn }).(pulumi.Float64PtrOutput)
}

func (o NetworkVirtualApplianceOutput) VirtualApplianceNics() VirtualApplianceNicPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) VirtualApplianceNicPropertiesResponseArrayOutput {
		return v.VirtualApplianceNics
	}).(VirtualApplianceNicPropertiesResponseArrayOutput)
}

func (o NetworkVirtualApplianceOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) SubResourceResponsePtrOutput { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkVirtualApplianceOutput{})
}
