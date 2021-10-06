


package v20160801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppVnetConnection(ctx *pulumi.Context, args *LookupWebAppVnetConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppVnetConnectionResult, error) {
	var rv LookupWebAppVnetConnectionResult
	err := ctx.Invoke("azure-native:web/v20160801:getWebAppVnetConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppVnetConnectionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VnetName          string `pulumi:"vnetName"`
}


type LookupWebAppVnetConnectionResult struct {
	CertBlob       *string             `pulumi:"certBlob"`
	CertThumbprint string              `pulumi:"certThumbprint"`
	DnsServers     *string             `pulumi:"dnsServers"`
	Id             string              `pulumi:"id"`
	Kind           *string             `pulumi:"kind"`
	Name           string              `pulumi:"name"`
	ResyncRequired bool                `pulumi:"resyncRequired"`
	Routes         []VnetRouteResponse `pulumi:"routes"`
	Type           string              `pulumi:"type"`
	VnetResourceId *string             `pulumi:"vnetResourceId"`
}
