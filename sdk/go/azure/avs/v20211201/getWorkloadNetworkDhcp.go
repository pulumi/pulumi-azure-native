


package v20211201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkDhcp(ctx *pulumi.Context, args *LookupWorkloadNetworkDhcpArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkDhcpResult, error) {
	var rv LookupWorkloadNetworkDhcpResult
	err := ctx.Invoke("azure-native:avs/v20211201:getWorkloadNetworkDhcp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkDhcpArgs struct {
	DhcpId            string `pulumi:"dhcpId"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWorkloadNetworkDhcpResult struct {
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}
