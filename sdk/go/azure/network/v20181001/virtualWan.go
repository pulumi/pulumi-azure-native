


package v20181001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualWan struct {
	pulumi.CustomResourceState

	AllowBranchToBranchTraffic     pulumi.BoolPtrOutput                         `pulumi:"allowBranchToBranchTraffic"`
	AllowVnetToVnetTraffic         pulumi.BoolPtrOutput                         `pulumi:"allowVnetToVnetTraffic"`
	DisableVpnEncryption           pulumi.BoolPtrOutput                         `pulumi:"disableVpnEncryption"`
	Etag                           pulumi.StringOutput                          `pulumi:"etag"`
	Location                       pulumi.StringOutput                          `pulumi:"location"`
	Name                           pulumi.StringOutput                          `pulumi:"name"`
	Office365LocalBreakoutCategory pulumi.StringOutput                          `pulumi:"office365LocalBreakoutCategory"`
	P2SVpnServerConfigurations     P2SVpnServerConfigurationResponseArrayOutput `pulumi:"p2SVpnServerConfigurations"`
	ProvisioningState              pulumi.StringOutput                          `pulumi:"provisioningState"`
	SecurityProviderName           pulumi.StringPtrOutput                       `pulumi:"securityProviderName"`
	Tags                           pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                           pulumi.StringOutput                          `pulumi:"type"`
	VirtualHubs                    SubResourceResponseArrayOutput               `pulumi:"virtualHubs"`
	VpnSites                       SubResourceResponseArrayOutput               `pulumi:"vpnSites"`
}


func NewVirtualWan(ctx *pulumi.Context,
	name string, args *VirtualWanArgs, opts ...pulumi.ResourceOption) (*VirtualWan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualWan"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualWan
	err := ctx.RegisterResource("azure-native:network/v20181001:VirtualWan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualWan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualWanState, opts ...pulumi.ResourceOption) (*VirtualWan, error) {
	var resource VirtualWan
	err := ctx.ReadResource("azure-native:network/v20181001:VirtualWan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualWanState struct {
}

type VirtualWanState struct {
}

func (VirtualWanState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWanState)(nil)).Elem()
}

type virtualWanArgs struct {
	AllowBranchToBranchTraffic *bool                           `pulumi:"allowBranchToBranchTraffic"`
	AllowVnetToVnetTraffic     *bool                           `pulumi:"allowVnetToVnetTraffic"`
	DisableVpnEncryption       *bool                           `pulumi:"disableVpnEncryption"`
	Id                         *string                         `pulumi:"id"`
	Location                   *string                         `pulumi:"location"`
	P2SVpnServerConfigurations []P2SVpnServerConfigurationType `pulumi:"p2SVpnServerConfigurations"`
	ResourceGroupName          string                          `pulumi:"resourceGroupName"`
	SecurityProviderName       *string                         `pulumi:"securityProviderName"`
	Tags                       map[string]string               `pulumi:"tags"`
	VirtualWANName             *string                         `pulumi:"virtualWANName"`
}


type VirtualWanArgs struct {
	AllowBranchToBranchTraffic pulumi.BoolPtrInput
	AllowVnetToVnetTraffic     pulumi.BoolPtrInput
	DisableVpnEncryption       pulumi.BoolPtrInput
	Id                         pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	P2SVpnServerConfigurations P2SVpnServerConfigurationTypeArrayInput
	ResourceGroupName          pulumi.StringInput
	SecurityProviderName       pulumi.StringPtrInput
	Tags                       pulumi.StringMapInput
	VirtualWANName             pulumi.StringPtrInput
}

func (VirtualWanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWanArgs)(nil)).Elem()
}

type VirtualWanInput interface {
	pulumi.Input

	ToVirtualWanOutput() VirtualWanOutput
	ToVirtualWanOutputWithContext(ctx context.Context) VirtualWanOutput
}

func (*VirtualWan) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualWan)(nil)).Elem()
}

func (i *VirtualWan) ToVirtualWanOutput() VirtualWanOutput {
	return i.ToVirtualWanOutputWithContext(context.Background())
}

func (i *VirtualWan) ToVirtualWanOutputWithContext(ctx context.Context) VirtualWanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualWanOutput)
}

type VirtualWanOutput struct{ *pulumi.OutputState }

func (VirtualWanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualWan)(nil)).Elem()
}

func (o VirtualWanOutput) ToVirtualWanOutput() VirtualWanOutput {
	return o
}

func (o VirtualWanOutput) ToVirtualWanOutputWithContext(ctx context.Context) VirtualWanOutput {
	return o
}

func (o VirtualWanOutput) AllowBranchToBranchTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.BoolPtrOutput { return v.AllowBranchToBranchTraffic }).(pulumi.BoolPtrOutput)
}

func (o VirtualWanOutput) AllowVnetToVnetTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.BoolPtrOutput { return v.AllowVnetToVnetTraffic }).(pulumi.BoolPtrOutput)
}

func (o VirtualWanOutput) DisableVpnEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.BoolPtrOutput { return v.DisableVpnEncryption }).(pulumi.BoolPtrOutput)
}

func (o VirtualWanOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualWanOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualWanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualWanOutput) Office365LocalBreakoutCategory() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.Office365LocalBreakoutCategory }).(pulumi.StringOutput)
}

func (o VirtualWanOutput) P2SVpnServerConfigurations() P2SVpnServerConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualWan) P2SVpnServerConfigurationResponseArrayOutput { return v.P2SVpnServerConfigurations }).(P2SVpnServerConfigurationResponseArrayOutput)
}

func (o VirtualWanOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualWanOutput) SecurityProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringPtrOutput { return v.SecurityProviderName }).(pulumi.StringPtrOutput)
}

func (o VirtualWanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualWanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualWanOutput) VirtualHubs() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualWan) SubResourceResponseArrayOutput { return v.VirtualHubs }).(SubResourceResponseArrayOutput)
}

func (o VirtualWanOutput) VpnSites() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualWan) SubResourceResponseArrayOutput { return v.VpnSites }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualWanOutput{})
}
