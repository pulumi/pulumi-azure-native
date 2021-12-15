


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualApplianceSite(ctx *pulumi.Context, args *LookupVirtualApplianceSiteArgs, opts ...pulumi.InvokeOption) (*LookupVirtualApplianceSiteResult, error) {
	var rv LookupVirtualApplianceSiteResult
	err := ctx.Invoke("azure-native:network/v20200801:getVirtualApplianceSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualApplianceSiteArgs struct {
	NetworkVirtualApplianceName string `pulumi:"networkVirtualApplianceName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	SiteName                    string `pulumi:"siteName"`
}


type LookupVirtualApplianceSiteResult struct {
	AddressPrefix     *string                            `pulumi:"addressPrefix"`
	Etag              string                             `pulumi:"etag"`
	Id                *string                            `pulumi:"id"`
	Name              *string                            `pulumi:"name"`
	O365Policy        *Office365PolicyPropertiesResponse `pulumi:"o365Policy"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	Type              string                             `pulumi:"type"`
}
