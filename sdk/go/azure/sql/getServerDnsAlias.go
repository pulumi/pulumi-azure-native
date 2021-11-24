


package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerDnsAlias(ctx *pulumi.Context, args *LookupServerDnsAliasArgs, opts ...pulumi.InvokeOption) (*LookupServerDnsAliasResult, error) {
	var rv LookupServerDnsAliasResult
	err := ctx.Invoke("azure-native:sql:getServerDnsAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerDnsAliasArgs struct {
	DnsAliasName      string `pulumi:"dnsAliasName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerDnsAliasResult struct {
	AzureDnsRecord string `pulumi:"azureDnsRecord"`
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	Type           string `pulumi:"type"`
}
