


package v20210801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:sql/v20210801preview:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
}


type LookupServerResult struct {
	AdministratorLogin            *string                                   `pulumi:"administratorLogin"`
	Administrators                *ServerExternalAdministratorResponse      `pulumi:"administrators"`
	FederatedClientId             *string                                   `pulumi:"federatedClientId"`
	FullyQualifiedDomainName      string                                    `pulumi:"fullyQualifiedDomainName"`
	Id                            string                                    `pulumi:"id"`
	Identity                      *ResourceIdentityResponse                 `pulumi:"identity"`
	KeyId                         *string                                   `pulumi:"keyId"`
	Kind                          string                                    `pulumi:"kind"`
	Location                      string                                    `pulumi:"location"`
	MinimalTlsVersion             *string                                   `pulumi:"minimalTlsVersion"`
	Name                          string                                    `pulumi:"name"`
	PrimaryUserAssignedIdentityId *string                                   `pulumi:"primaryUserAssignedIdentityId"`
	PrivateEndpointConnections    []ServerPrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess           *string                                   `pulumi:"publicNetworkAccess"`
	RestrictOutboundNetworkAccess *string                                   `pulumi:"restrictOutboundNetworkAccess"`
	State                         string                                    `pulumi:"state"`
	Tags                          map[string]string                         `pulumi:"tags"`
	Type                          string                                    `pulumi:"type"`
	Version                       *string                                   `pulumi:"version"`
	WorkspaceFeature              string                                    `pulumi:"workspaceFeature"`
}
