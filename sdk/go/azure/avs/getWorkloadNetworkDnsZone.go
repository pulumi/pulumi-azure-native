


package avs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkDnsZone(ctx *pulumi.Context, args *LookupWorkloadNetworkDnsZoneArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkDnsZoneResult, error) {
	var rv LookupWorkloadNetworkDnsZoneResult
	err := ctx.Invoke("azure-native:avs:getWorkloadNetworkDnsZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkDnsZoneArgs struct {
	DnsZoneId         string `pulumi:"dnsZoneId"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWorkloadNetworkDnsZoneResult struct {
	DisplayName       *string  `pulumi:"displayName"`
	DnsServerIps      []string `pulumi:"dnsServerIps"`
	DnsServices       *float64 `pulumi:"dnsServices"`
	Domain            []string `pulumi:"domain"`
	Id                string   `pulumi:"id"`
	Name              string   `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	SourceIp          *string  `pulumi:"sourceIp"`
	Type              string   `pulumi:"type"`
}
