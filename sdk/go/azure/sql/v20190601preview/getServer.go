


package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:sql/v20190601preview:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerResult struct {
	AdministratorLogin         *string                                   `pulumi:"administratorLogin"`
	FullyQualifiedDomainName   string                                    `pulumi:"fullyQualifiedDomainName"`
	Id                         string                                    `pulumi:"id"`
	Identity                   *ResourceIdentityResponse                 `pulumi:"identity"`
	Kind                       string                                    `pulumi:"kind"`
	Location                   string                                    `pulumi:"location"`
	MinimalTlsVersion          *string                                   `pulumi:"minimalTlsVersion"`
	Name                       string                                    `pulumi:"name"`
	PrivateEndpointConnections []ServerPrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        *string                                   `pulumi:"publicNetworkAccess"`
	State                      string                                    `pulumi:"state"`
	Tags                       map[string]string                         `pulumi:"tags"`
	Type                       string                                    `pulumi:"type"`
	Version                    *string                                   `pulumi:"version"`
}
