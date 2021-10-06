


package v20161201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPublicIPAddress(ctx *pulumi.Context, args *LookupPublicIPAddressArgs, opts ...pulumi.InvokeOption) (*LookupPublicIPAddressResult, error) {
	var rv LookupPublicIPAddressResult
	err := ctx.Invoke("azure-native:network/v20161201:getPublicIPAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicIPAddressArgs struct {
	Expand              *string `pulumi:"expand"`
	PublicIpAddressName string  `pulumi:"publicIpAddressName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type LookupPublicIPAddressResult struct {
	DnsSettings              *PublicIPAddressDnsSettingsResponse `pulumi:"dnsSettings"`
	Etag                     *string                             `pulumi:"etag"`
	Id                       *string                             `pulumi:"id"`
	IdleTimeoutInMinutes     *int                                `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                             `pulumi:"ipAddress"`
	IpConfiguration          IPConfigurationResponse             `pulumi:"ipConfiguration"`
	Location                 *string                             `pulumi:"location"`
	Name                     string                              `pulumi:"name"`
	ProvisioningState        *string                             `pulumi:"provisioningState"`
	PublicIPAddressVersion   *string                             `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                             `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             *string                             `pulumi:"resourceGuid"`
	Tags                     map[string]string                   `pulumi:"tags"`
	Type                     string                              `pulumi:"type"`
}
