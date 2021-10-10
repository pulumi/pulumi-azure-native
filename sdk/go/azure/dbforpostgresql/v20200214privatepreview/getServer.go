


package v20200214privatepreview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20200214privatepreview:getServer", args, &rv, opts...)
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
	AdministratorLogin       *string                                           `pulumi:"administratorLogin"`
	AvailabilityZone         *string                                           `pulumi:"availabilityZone"`
	ByokEnforcement          string                                            `pulumi:"byokEnforcement"`
	DelegatedSubnetArguments *ServerPropertiesResponseDelegatedSubnetArguments `pulumi:"delegatedSubnetArguments"`
	DisplayName              *string                                           `pulumi:"displayName"`
	FullyQualifiedDomainName string                                            `pulumi:"fullyQualifiedDomainName"`
	HaEnabled                *string                                           `pulumi:"haEnabled"`
	HaState                  string                                            `pulumi:"haState"`
	Id                       string                                            `pulumi:"id"`
	Identity                 *IdentityResponse                                 `pulumi:"identity"`
	Location                 string                                            `pulumi:"location"`
	MaintenanceWindow        *MaintenanceWindowResponse                        `pulumi:"maintenanceWindow"`
	Name                     string                                            `pulumi:"name"`
	PointInTimeUTC           *string                                           `pulumi:"pointInTimeUTC"`
	PrivateDnsZoneArguments  *ServerPropertiesResponsePrivateDnsZoneArguments  `pulumi:"privateDnsZoneArguments"`
	PublicNetworkAccess      string                                            `pulumi:"publicNetworkAccess"`
	Sku                      *SkuResponse                                      `pulumi:"sku"`
	SourceResourceGroupName  *string                                           `pulumi:"sourceResourceGroupName"`
	SourceServerName         *string                                           `pulumi:"sourceServerName"`
	SourceSubscriptionId     *string                                           `pulumi:"sourceSubscriptionId"`
	StandbyAvailabilityZone  string                                            `pulumi:"standbyAvailabilityZone"`
	State                    string                                            `pulumi:"state"`
	StorageProfile           *StorageProfileResponse                           `pulumi:"storageProfile"`
	Tags                     map[string]string                                 `pulumi:"tags"`
	Type                     string                                            `pulumi:"type"`
	Version                  *string                                           `pulumi:"version"`
}
