


package v20200701privatepreview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbformysql/v20200701privatepreview:getServer", args, &rv, opts...)
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
	AdministratorLogin       *string                           `pulumi:"administratorLogin"`
	AvailabilityZone         *string                           `pulumi:"availabilityZone"`
	ByokEnforcement          string                            `pulumi:"byokEnforcement"`
	DelegatedSubnetArguments *DelegatedSubnetArgumentsResponse `pulumi:"delegatedSubnetArguments"`
	EarliestRestoreDate      string                            `pulumi:"earliestRestoreDate"`
	FullyQualifiedDomainName string                            `pulumi:"fullyQualifiedDomainName"`
	HaEnabled                *string                           `pulumi:"haEnabled"`
	HaState                  string                            `pulumi:"haState"`
	Id                       string                            `pulumi:"id"`
	Identity                 *IdentityResponse                 `pulumi:"identity"`
	Location                 string                            `pulumi:"location"`
	MaintenanceWindow        *MaintenanceWindowResponse        `pulumi:"maintenanceWindow"`
	Name                     string                            `pulumi:"name"`
	PublicNetworkAccess      string                            `pulumi:"publicNetworkAccess"`
	ReplicaCapacity          int                               `pulumi:"replicaCapacity"`
	ReplicationRole          *string                           `pulumi:"replicationRole"`
	Sku                      *SkuResponse                      `pulumi:"sku"`
	SourceServerId           *string                           `pulumi:"sourceServerId"`
	SslEnforcement           *string                           `pulumi:"sslEnforcement"`
	StandbyAvailabilityZone  string                            `pulumi:"standbyAvailabilityZone"`
	State                    string                            `pulumi:"state"`
	StorageProfile           *StorageProfileResponse           `pulumi:"storageProfile"`
	Tags                     map[string]string                 `pulumi:"tags"`
	Type                     string                            `pulumi:"type"`
	Version                  *string                           `pulumi:"version"`
}
