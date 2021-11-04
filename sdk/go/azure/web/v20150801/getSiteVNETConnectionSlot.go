


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSiteVNETConnectionSlot(ctx *pulumi.Context, args *LookupSiteVNETConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteVNETConnectionSlotResult, error) {
	var rv LookupSiteVNETConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteVNETConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteVNETConnectionSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
	VnetName          string `pulumi:"vnetName"`
}


type LookupSiteVNETConnectionSlotResult struct {
	CertBlob       *string             `pulumi:"certBlob"`
	CertThumbprint *string             `pulumi:"certThumbprint"`
	DnsServers     *string             `pulumi:"dnsServers"`
	Id             *string             `pulumi:"id"`
	Kind           *string             `pulumi:"kind"`
	Location       string              `pulumi:"location"`
	Name           *string             `pulumi:"name"`
	ResyncRequired *bool               `pulumi:"resyncRequired"`
	Routes         []VnetRouteResponse `pulumi:"routes"`
	Tags           map[string]string   `pulumi:"tags"`
	Type           *string             `pulumi:"type"`
	VnetResourceId *string             `pulumi:"vnetResourceId"`
}
