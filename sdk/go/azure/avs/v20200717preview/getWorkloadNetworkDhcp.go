


package v20200717preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkDhcp(ctx *pulumi.Context, args *LookupWorkloadNetworkDhcpArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkDhcpResult, error) {
	var rv LookupWorkloadNetworkDhcpResult
	err := ctx.Invoke("azure-native:avs/v20200717preview:getWorkloadNetworkDhcp", args, &rv, opts...)
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
	DhcpType          string   `pulumi:"dhcpType"`
	DisplayName       *string  `pulumi:"displayName"`
	Id                string   `pulumi:"id"`
	Name              string   `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	Segments          []string `pulumi:"segments"`
	Type              string   `pulumi:"type"`
}
