


package v20200301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionProxy(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionProxyArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionProxyResult, error) {
	var rv LookupPrivateEndpointConnectionProxyResult
	err := ctx.Invoke("azure-native:deviceupdate/v20200301preview:getPrivateEndpointConnectionProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionProxyArgs struct {
	AccountName                      string `pulumi:"accountName"`
	PrivateEndpointConnectionProxyId string `pulumi:"privateEndpointConnectionProxyId"`
	ResourceGroupName                string `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointConnectionProxyResult struct {
	ETag                  string                         `pulumi:"eTag"`
	Id                    string                         `pulumi:"id"`
	Name                  string                         `pulumi:"name"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
	RemotePrivateEndpoint *RemotePrivateEndpointResponse `pulumi:"remotePrivateEndpoint"`
	Status                *string                        `pulumi:"status"`
	SystemData            SystemDataResponse             `pulumi:"systemData"`
	Type                  string                         `pulumi:"type"`
}
