


package v20201101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPublicIPAddress(ctx *pulumi.Context, args *LookupPublicIPAddressArgs, opts ...pulumi.InvokeOption) (*LookupPublicIPAddressResult, error) {
	var rv LookupPublicIPAddressResult
	err := ctx.Invoke("azure-native:network/v20201101:getPublicIPAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPublicIPAddressArgs struct {
	Expand              *string `pulumi:"expand"`
	PublicIpAddressName string  `pulumi:"publicIpAddressName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type LookupPublicIPAddressResult struct {
	DdosSettings             *DdosSettingsResponse               `pulumi:"ddosSettings"`
	DnsSettings              *PublicIPAddressDnsSettingsResponse `pulumi:"dnsSettings"`
	Etag                     string                              `pulumi:"etag"`
	ExtendedLocation         *ExtendedLocationResponse           `pulumi:"extendedLocation"`
	Id                       *string                             `pulumi:"id"`
	IdleTimeoutInMinutes     *int                                `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                             `pulumi:"ipAddress"`
	IpConfiguration          IPConfigurationResponse             `pulumi:"ipConfiguration"`
	IpTags                   []IpTagResponse                     `pulumi:"ipTags"`
	LinkedPublicIPAddress    *PublicIPAddressResponse            `pulumi:"linkedPublicIPAddress"`
	Location                 *string                             `pulumi:"location"`
	MigrationPhase           *string                             `pulumi:"migrationPhase"`
	Name                     string                              `pulumi:"name"`
	NatGateway               *NatGatewayResponse                 `pulumi:"natGateway"`
	ProvisioningState        string                              `pulumi:"provisioningState"`
	PublicIPAddressVersion   *string                             `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                             `pulumi:"publicIPAllocationMethod"`
	PublicIPPrefix           *SubResourceResponse                `pulumi:"publicIPPrefix"`
	ResourceGuid             string                              `pulumi:"resourceGuid"`
	ServicePublicIPAddress   *PublicIPAddressResponse            `pulumi:"servicePublicIPAddress"`
	Sku                      *PublicIPAddressSkuResponse         `pulumi:"sku"`
	Tags                     map[string]string                   `pulumi:"tags"`
	Type                     string                              `pulumi:"type"`
	Zones                    []string                            `pulumi:"zones"`
}


func (val *LookupPublicIPAddressResult) Defaults() *LookupPublicIPAddressResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.IpConfiguration = *tmp.IpConfiguration.Defaults()

	tmp.LinkedPublicIPAddress = tmp.LinkedPublicIPAddress.Defaults()

	tmp.ServicePublicIPAddress = tmp.ServicePublicIPAddress.Defaults()

	return &tmp
}
