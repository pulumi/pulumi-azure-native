


package v20201005privatepreview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerGroup(ctx *pulumi.Context, args *LookupServerGroupArgs, opts ...pulumi.InvokeOption) (*LookupServerGroupResult, error) {
	var rv LookupServerGroupResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20201005privatepreview:getServerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerGroupArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerGroupName   string `pulumi:"serverGroupName"`
}


type LookupServerGroupResult struct {
	AdministratorLogin        *string                                                `pulumi:"administratorLogin"`
	AvailabilityZone          *string                                                `pulumi:"availabilityZone"`
	BackupRetentionDays       *int                                                   `pulumi:"backupRetentionDays"`
	CitusVersion              *string                                                `pulumi:"citusVersion"`
	DelegatedSubnetArguments  *ServerGroupPropertiesResponseDelegatedSubnetArguments `pulumi:"delegatedSubnetArguments"`
	EarliestRestoreTime       string                                                 `pulumi:"earliestRestoreTime"`
	EnableMx                  *bool                                                  `pulumi:"enableMx"`
	EnableShardsOnCoordinator *bool                                                  `pulumi:"enableShardsOnCoordinator"`
	EnableZfs                 *bool                                                  `pulumi:"enableZfs"`
	Id                        string                                                 `pulumi:"id"`
	Location                  string                                                 `pulumi:"location"`
	MaintenanceWindow         *MaintenanceWindowResponse                             `pulumi:"maintenanceWindow"`
	Name                      string                                                 `pulumi:"name"`
	PostgresqlVersion         *string                                                `pulumi:"postgresqlVersion"`
	PrivateDnsZoneArguments   *ServerGroupPropertiesResponsePrivateDnsZoneArguments  `pulumi:"privateDnsZoneArguments"`
	ReadReplicas              []string                                               `pulumi:"readReplicas"`
	ResourceProviderType      string                                                 `pulumi:"resourceProviderType"`
	ServerRoleGroups          []ServerRoleGroupResponse                              `pulumi:"serverRoleGroups"`
	SourceServerGroup         string                                                 `pulumi:"sourceServerGroup"`
	StandbyAvailabilityZone   *string                                                `pulumi:"standbyAvailabilityZone"`
	State                     string                                                 `pulumi:"state"`
	SystemData                SystemDataResponse                                     `pulumi:"systemData"`
	Tags                      map[string]string                                      `pulumi:"tags"`
	Type                      string                                                 `pulumi:"type"`
}
