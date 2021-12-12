


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationNetworkMapping(ctx *pulumi.Context, args *LookupReplicationNetworkMappingArgs, opts ...pulumi.InvokeOption) (*LookupReplicationNetworkMappingResult, error) {
	var rv LookupReplicationNetworkMappingResult
	err := ctx.Invoke("azure-native:recoveryservices/v20211001:getReplicationNetworkMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationNetworkMappingArgs struct {
	FabricName         string `pulumi:"fabricName"`
	NetworkMappingName string `pulumi:"networkMappingName"`
	NetworkName        string `pulumi:"networkName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	ResourceName       string `pulumi:"resourceName"`
}


type LookupReplicationNetworkMappingResult struct {
	Id         string                           `pulumi:"id"`
	Location   *string                          `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties NetworkMappingPropertiesResponse `pulumi:"properties"`
	Type       string                           `pulumi:"type"`
}
