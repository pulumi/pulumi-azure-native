


package v20210115

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppVnetConnectionSlot(ctx *pulumi.Context, args *LookupWebAppVnetConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppVnetConnectionSlotResult, error) {
	var rv LookupWebAppVnetConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20210115:getWebAppVnetConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppVnetConnectionSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
	VnetName          string `pulumi:"vnetName"`
}


type LookupWebAppVnetConnectionSlotResult struct {
	CertBlob       *string             `pulumi:"certBlob"`
	CertThumbprint string              `pulumi:"certThumbprint"`
	DnsServers     *string             `pulumi:"dnsServers"`
	Id             string              `pulumi:"id"`
	IsSwift        *bool               `pulumi:"isSwift"`
	Kind           *string             `pulumi:"kind"`
	Name           string              `pulumi:"name"`
	ResyncRequired bool                `pulumi:"resyncRequired"`
	Routes         []VnetRouteResponse `pulumi:"routes"`
	Type           string              `pulumi:"type"`
	VnetResourceId *string             `pulumi:"vnetResourceId"`
}
