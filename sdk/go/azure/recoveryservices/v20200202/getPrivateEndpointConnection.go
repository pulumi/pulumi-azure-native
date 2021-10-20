


package v20200202

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionResult, error) {
	var rv LookupPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:recoveryservices/v20200202:getPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	VaultName                     string `pulumi:"vaultName"`
}


type LookupPrivateEndpointConnectionResult struct {
	ETag       *string                           `pulumi:"eTag"`
	Id         string                            `pulumi:"id"`
	Location   *string                           `pulumi:"location"`
	Name       string                            `pulumi:"name"`
	Properties PrivateEndpointConnectionResponse `pulumi:"properties"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       string                            `pulumi:"type"`
}
