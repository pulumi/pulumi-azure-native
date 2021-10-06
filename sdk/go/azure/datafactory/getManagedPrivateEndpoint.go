


package datafactory

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedPrivateEndpoint(ctx *pulumi.Context, args *LookupManagedPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupManagedPrivateEndpointResult, error) {
	var rv LookupManagedPrivateEndpointResult
	err := ctx.Invoke("azure-native:datafactory:getManagedPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedPrivateEndpointArgs struct {
	FactoryName                string `pulumi:"factoryName"`
	ManagedPrivateEndpointName string `pulumi:"managedPrivateEndpointName"`
	ManagedVirtualNetworkName  string `pulumi:"managedVirtualNetworkName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}


type LookupManagedPrivateEndpointResult struct {
	Etag       string                         `pulumi:"etag"`
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties ManagedPrivateEndpointResponse `pulumi:"properties"`
	Type       string                         `pulumi:"type"`
}
