


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualWan(ctx *pulumi.Context, args *LookupVirtualWanArgs, opts ...pulumi.InvokeOption) (*LookupVirtualWanResult, error) {
	var rv LookupVirtualWanResult
	err := ctx.Invoke("azure-native:network:getVirtualWan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualWanArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualWANName    string `pulumi:"virtualWANName"`
}


type LookupVirtualWanResult struct {
	AllowBranchToBranchTraffic     *bool                 `pulumi:"allowBranchToBranchTraffic"`
	AllowVnetToVnetTraffic         *bool                 `pulumi:"allowVnetToVnetTraffic"`
	DisableVpnEncryption           *bool                 `pulumi:"disableVpnEncryption"`
	Etag                           string                `pulumi:"etag"`
	Id                             *string               `pulumi:"id"`
	Location                       string                `pulumi:"location"`
	Name                           string                `pulumi:"name"`
	Office365LocalBreakoutCategory string                `pulumi:"office365LocalBreakoutCategory"`
	ProvisioningState              string                `pulumi:"provisioningState"`
	Tags                           map[string]string     `pulumi:"tags"`
	Type                           string                `pulumi:"type"`
	VirtualHubs                    []SubResourceResponse `pulumi:"virtualHubs"`
	VpnSites                       []SubResourceResponse `pulumi:"vpnSites"`
}
