


package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVpnSite(ctx *pulumi.Context, args *LookupVpnSiteArgs, opts ...pulumi.InvokeOption) (*LookupVpnSiteResult, error) {
	var rv LookupVpnSiteResult
	err := ctx.Invoke("azure-native:network/v20190701:getVpnSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnSiteArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VpnSiteName       string `pulumi:"vpnSiteName"`
}


type LookupVpnSiteResult struct {
	AddressSpace      *AddressSpaceResponse     `pulumi:"addressSpace"`
	BgpProperties     *BgpSettingsResponse      `pulumi:"bgpProperties"`
	DeviceProperties  *DevicePropertiesResponse `pulumi:"deviceProperties"`
	Etag              string                    `pulumi:"etag"`
	Id                *string                   `pulumi:"id"`
	IpAddress         *string                   `pulumi:"ipAddress"`
	IsSecuritySite    *bool                     `pulumi:"isSecuritySite"`
	Location          string                    `pulumi:"location"`
	Name              string                    `pulumi:"name"`
	ProvisioningState string                    `pulumi:"provisioningState"`
	SiteKey           *string                   `pulumi:"siteKey"`
	Tags              map[string]string         `pulumi:"tags"`
	Type              string                    `pulumi:"type"`
	VirtualWan        *SubResourceResponse      `pulumi:"virtualWan"`
	VpnSiteLinks      []VpnSiteLinkResponse     `pulumi:"vpnSiteLinks"`
}
