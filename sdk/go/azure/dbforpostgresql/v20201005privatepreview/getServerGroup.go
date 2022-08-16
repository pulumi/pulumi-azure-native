


package v20201005privatepreview

import (
	"context"
	"reflect"

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

func LookupServerGroupOutput(ctx *pulumi.Context, args LookupServerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupServerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerGroupResult, error) {
			args := v.(LookupServerGroupArgs)
			r, err := LookupServerGroup(ctx, &args, opts...)
			var s LookupServerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerGroupResultOutput)
}

type LookupServerGroupOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerGroupName   pulumi.StringInput `pulumi:"serverGroupName"`
}

func (LookupServerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerGroupArgs)(nil)).Elem()
}


type LookupServerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupServerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerGroupResult)(nil)).Elem()
}

func (o LookupServerGroupResultOutput) ToLookupServerGroupResultOutput() LookupServerGroupResultOutput {
	return o
}

func (o LookupServerGroupResultOutput) ToLookupServerGroupResultOutputWithContext(ctx context.Context) LookupServerGroupResultOutput {
	return o
}

func (o LookupServerGroupResultOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *string { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

func (o LookupServerGroupResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o LookupServerGroupResultOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o LookupServerGroupResultOutput) CitusVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *string { return v.CitusVersion }).(pulumi.StringPtrOutput)
}

func (o LookupServerGroupResultOutput) DelegatedSubnetArguments() ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *ServerGroupPropertiesResponseDelegatedSubnetArguments {
		return v.DelegatedSubnetArguments
	}).(ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput)
}

func (o LookupServerGroupResultOutput) EarliestRestoreTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.EarliestRestoreTime }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) EnableMx() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *bool { return v.EnableMx }).(pulumi.BoolPtrOutput)
}

func (o LookupServerGroupResultOutput) EnableShardsOnCoordinator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *bool { return v.EnableShardsOnCoordinator }).(pulumi.BoolPtrOutput)
}

func (o LookupServerGroupResultOutput) EnableZfs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *bool { return v.EnableZfs }).(pulumi.BoolPtrOutput)
}

func (o LookupServerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *MaintenanceWindowResponse { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

func (o LookupServerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) PostgresqlVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *string { return v.PostgresqlVersion }).(pulumi.StringPtrOutput)
}

func (o LookupServerGroupResultOutput) PrivateDnsZoneArguments() ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *ServerGroupPropertiesResponsePrivateDnsZoneArguments {
		return v.PrivateDnsZoneArguments
	}).(ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput)
}

func (o LookupServerGroupResultOutput) ReadReplicas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerGroupResult) []string { return v.ReadReplicas }).(pulumi.StringArrayOutput)
}

func (o LookupServerGroupResultOutput) ResourceProviderType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.ResourceProviderType }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) ServerRoleGroups() ServerRoleGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupServerGroupResult) []ServerRoleGroupResponse { return v.ServerRoleGroups }).(ServerRoleGroupResponseArrayOutput)
}

func (o LookupServerGroupResultOutput) SourceServerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.SourceServerGroup }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) StandbyAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerGroupResult) *string { return v.StandbyAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o LookupServerGroupResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupServerGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServerGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupServerGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerGroupResultOutput{})
}
