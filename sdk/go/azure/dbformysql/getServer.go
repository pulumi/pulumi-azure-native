


package dbformysql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbformysql:getServer", args, &rv, opts...)
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
	ByokEnforcement            string                                    `pulumi:"byokEnforcement"`
	EarliestRestoreDate        *string                                   `pulumi:"earliestRestoreDate"`
	FullyQualifiedDomainName   *string                                   `pulumi:"fullyQualifiedDomainName"`
	Id                         string                                    `pulumi:"id"`
	Identity                   *ResourceIdentityResponse                 `pulumi:"identity"`
	InfrastructureEncryption   *string                                   `pulumi:"infrastructureEncryption"`
	Location                   string                                    `pulumi:"location"`
	MasterServerId             *string                                   `pulumi:"masterServerId"`
	MinimalTlsVersion          *string                                   `pulumi:"minimalTlsVersion"`
	Name                       string                                    `pulumi:"name"`
	PrivateEndpointConnections []ServerPrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        *string                                   `pulumi:"publicNetworkAccess"`
	ReplicaCapacity            *int                                      `pulumi:"replicaCapacity"`
	ReplicationRole            *string                                   `pulumi:"replicationRole"`
	Sku                        *SkuResponse                              `pulumi:"sku"`
	SslEnforcement             *string                                   `pulumi:"sslEnforcement"`
	StorageProfile             *StorageProfileResponse                   `pulumi:"storageProfile"`
	Tags                       map[string]string                         `pulumi:"tags"`
	Type                       string                                    `pulumi:"type"`
	UserVisibleState           *string                                   `pulumi:"userVisibleState"`
	Version                    *string                                   `pulumi:"version"`
}
