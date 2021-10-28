


package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnSite struct {
	pulumi.CustomResourceState

	AddressSpace      AddressSpaceResponsePtrOutput     `pulumi:"addressSpace"`
	BgpProperties     BgpSettingsResponsePtrOutput      `pulumi:"bgpProperties"`
	DeviceProperties  DevicePropertiesResponsePtrOutput `pulumi:"deviceProperties"`
	Etag              pulumi.StringOutput               `pulumi:"etag"`
	IpAddress         pulumi.StringPtrOutput            `pulumi:"ipAddress"`
	IsSecuritySite    pulumi.BoolPtrOutput              `pulumi:"isSecuritySite"`
	Location          pulumi.StringOutput               `pulumi:"location"`
	Name              pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState pulumi.StringOutput               `pulumi:"provisioningState"`
	SiteKey           pulumi.StringPtrOutput            `pulumi:"siteKey"`
	Tags              pulumi.StringMapOutput            `pulumi:"tags"`
	Type              pulumi.StringOutput               `pulumi:"type"`
	VirtualWan        SubResourceResponsePtrOutput      `pulumi:"virtualWan"`
	VpnSiteLinks      VpnSiteLinkResponseArrayOutput    `pulumi:"vpnSiteLinks"`
}


func NewVpnSite(ctx *pulumi.Context,
	name string, args *VpnSiteArgs, opts ...pulumi.ResourceOption) (*VpnSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:VpnSite"),
		},
	})
	opts = append(opts, aliases)
	var resource VpnSite
	err := ctx.RegisterResource("azure-native:network/v20191101:VpnSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVpnSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnSiteState, opts ...pulumi.ResourceOption) (*VpnSite, error) {
	var resource VpnSite
	err := ctx.ReadResource("azure-native:network/v20191101:VpnSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vpnSiteState struct {
}

type VpnSiteState struct {
}

func (VpnSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSiteState)(nil)).Elem()
}

type vpnSiteArgs struct {
	AddressSpace      *AddressSpace     `pulumi:"addressSpace"`
	BgpProperties     *BgpSettings      `pulumi:"bgpProperties"`
	DeviceProperties  *DeviceProperties `pulumi:"deviceProperties"`
	Id                *string           `pulumi:"id"`
	IpAddress         *string           `pulumi:"ipAddress"`
	IsSecuritySite    *bool             `pulumi:"isSecuritySite"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SiteKey           *string           `pulumi:"siteKey"`
	Tags              map[string]string `pulumi:"tags"`
	VirtualWan        *SubResource      `pulumi:"virtualWan"`
	VpnSiteLinks      []VpnSiteLink     `pulumi:"vpnSiteLinks"`
	VpnSiteName       *string           `pulumi:"vpnSiteName"`
}


type VpnSiteArgs struct {
	AddressSpace      AddressSpacePtrInput
	BgpProperties     BgpSettingsPtrInput
	DeviceProperties  DevicePropertiesPtrInput
	Id                pulumi.StringPtrInput
	IpAddress         pulumi.StringPtrInput
	IsSecuritySite    pulumi.BoolPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SiteKey           pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	VirtualWan        SubResourcePtrInput
	VpnSiteLinks      VpnSiteLinkArrayInput
	VpnSiteName       pulumi.StringPtrInput
}

func (VpnSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSiteArgs)(nil)).Elem()
}

type VpnSiteInput interface {
	pulumi.Input

	ToVpnSiteOutput() VpnSiteOutput
	ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput
}

func (*VpnSite) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSite)(nil))
}

func (i *VpnSite) ToVpnSiteOutput() VpnSiteOutput {
	return i.ToVpnSiteOutputWithContext(context.Background())
}

func (i *VpnSite) ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteOutput)
}

type VpnSiteOutput struct{ *pulumi.OutputState }

func (VpnSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSite)(nil))
}

func (o VpnSiteOutput) ToVpnSiteOutput() VpnSiteOutput {
	return o
}

func (o VpnSiteOutput) ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VpnSiteOutput{})
}
