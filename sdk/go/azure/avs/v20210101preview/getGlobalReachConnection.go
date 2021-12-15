


package v20210101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGlobalReachConnection(ctx *pulumi.Context, args *LookupGlobalReachConnectionArgs, opts ...pulumi.InvokeOption) (*LookupGlobalReachConnectionResult, error) {
	var rv LookupGlobalReachConnectionResult
	err := ctx.Invoke("azure-native:avs/v20210101preview:getGlobalReachConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalReachConnectionArgs struct {
	GlobalReachConnectionName string `pulumi:"globalReachConnectionName"`
	PrivateCloudName          string `pulumi:"privateCloudName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupGlobalReachConnectionResult struct {
	AddressPrefix           string  `pulumi:"addressPrefix"`
	AuthorizationKey        *string `pulumi:"authorizationKey"`
	CircuitConnectionStatus string  `pulumi:"circuitConnectionStatus"`
	Id                      string  `pulumi:"id"`
	Name                    string  `pulumi:"name"`
	PeerExpressRouteCircuit *string `pulumi:"peerExpressRouteCircuit"`
	ProvisioningState       string  `pulumi:"provisioningState"`
	Type                    string  `pulumi:"type"`
}
