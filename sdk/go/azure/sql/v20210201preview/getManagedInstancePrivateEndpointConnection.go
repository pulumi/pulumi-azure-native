


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedInstancePrivateEndpointConnection(ctx *pulumi.Context, args *LookupManagedInstancePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstancePrivateEndpointConnectionResult, error) {
	var rv LookupManagedInstancePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getManagedInstancePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstancePrivateEndpointConnectionArgs struct {
	ManagedInstanceName           string `pulumi:"managedInstanceName"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupManagedInstancePrivateEndpointConnectionResult struct {
	Id                                string                                                            `pulumi:"id"`
	Name                              string                                                            `pulumi:"name"`
	PrivateEndpoint                   *ManagedInstancePrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                                            `pulumi:"provisioningState"`
	Type                              string                                                            `pulumi:"type"`
}
