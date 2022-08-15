


package v20201005privatepreview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerGroup struct {
	pulumi.CustomResourceState

	AdministratorLogin        pulumi.StringPtrOutput                                         `pulumi:"administratorLogin"`
	AvailabilityZone          pulumi.StringPtrOutput                                         `pulumi:"availabilityZone"`
	BackupRetentionDays       pulumi.IntPtrOutput                                            `pulumi:"backupRetentionDays"`
	CitusVersion              pulumi.StringPtrOutput                                         `pulumi:"citusVersion"`
	DelegatedSubnetArguments  ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput `pulumi:"delegatedSubnetArguments"`
	EarliestRestoreTime       pulumi.StringOutput                                            `pulumi:"earliestRestoreTime"`
	EnableMx                  pulumi.BoolPtrOutput                                           `pulumi:"enableMx"`
	EnableShardsOnCoordinator pulumi.BoolPtrOutput                                           `pulumi:"enableShardsOnCoordinator"`
	EnableZfs                 pulumi.BoolPtrOutput                                           `pulumi:"enableZfs"`
	Location                  pulumi.StringOutput                                            `pulumi:"location"`
	MaintenanceWindow         MaintenanceWindowResponsePtrOutput                             `pulumi:"maintenanceWindow"`
	Name                      pulumi.StringOutput                                            `pulumi:"name"`
	PostgresqlVersion         pulumi.StringPtrOutput                                         `pulumi:"postgresqlVersion"`
	PrivateDnsZoneArguments   ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput  `pulumi:"privateDnsZoneArguments"`
	ReadReplicas              pulumi.StringArrayOutput                                       `pulumi:"readReplicas"`
	ResourceProviderType      pulumi.StringOutput                                            `pulumi:"resourceProviderType"`
	ServerRoleGroups          ServerRoleGroupResponseArrayOutput                             `pulumi:"serverRoleGroups"`
	SourceServerGroup         pulumi.StringOutput                                            `pulumi:"sourceServerGroup"`
	StandbyAvailabilityZone   pulumi.StringPtrOutput                                         `pulumi:"standbyAvailabilityZone"`
	State                     pulumi.StringOutput                                            `pulumi:"state"`
	SystemData                SystemDataResponseOutput                                       `pulumi:"systemData"`
	Tags                      pulumi.StringMapOutput                                         `pulumi:"tags"`
	Type                      pulumi.StringOutput                                            `pulumi:"type"`
}


func NewServerGroup(ctx *pulumi.Context,
	name string, args *ServerGroupArgs, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ServerGroup
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20201005privatepreview:ServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerGroupState, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	var resource ServerGroup
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20201005privatepreview:ServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverGroupState struct {
}

type ServerGroupState struct {
}

func (ServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupState)(nil)).Elem()
}

type serverGroupArgs struct {
	AdministratorLogin         *string                                        `pulumi:"administratorLogin"`
	AdministratorLoginPassword *string                                        `pulumi:"administratorLoginPassword"`
	AvailabilityZone           *string                                        `pulumi:"availabilityZone"`
	BackupRetentionDays        *int                                           `pulumi:"backupRetentionDays"`
	CitusVersion               *string                                        `pulumi:"citusVersion"`
	CreateMode                 *string                                        `pulumi:"createMode"`
	DelegatedSubnetArguments   *ServerGroupPropertiesDelegatedSubnetArguments `pulumi:"delegatedSubnetArguments"`
	EnableMx                   *bool                                          `pulumi:"enableMx"`
	EnableShardsOnCoordinator  *bool                                          `pulumi:"enableShardsOnCoordinator"`
	EnableZfs                  *bool                                          `pulumi:"enableZfs"`
	Location                   *string                                        `pulumi:"location"`
	MaintenanceWindow          *MaintenanceWindow                             `pulumi:"maintenanceWindow"`
	PointInTimeUTC             *string                                        `pulumi:"pointInTimeUTC"`
	PostgresqlVersion          *string                                        `pulumi:"postgresqlVersion"`
	PrivateDnsZoneArguments    *ServerGroupPropertiesPrivateDnsZoneArguments  `pulumi:"privateDnsZoneArguments"`
	ResourceGroupName          string                                         `pulumi:"resourceGroupName"`
	ServerGroupName            *string                                        `pulumi:"serverGroupName"`
	ServerRoleGroups           []ServerRoleGroup                              `pulumi:"serverRoleGroups"`
	SourceLocation             *string                                        `pulumi:"sourceLocation"`
	SourceResourceGroupName    *string                                        `pulumi:"sourceResourceGroupName"`
	SourceServerGroupName      *string                                        `pulumi:"sourceServerGroupName"`
	SourceSubscriptionId       *string                                        `pulumi:"sourceSubscriptionId"`
	StandbyAvailabilityZone    *string                                        `pulumi:"standbyAvailabilityZone"`
	Tags                       map[string]string                              `pulumi:"tags"`
}


type ServerGroupArgs struct {
	AdministratorLogin         pulumi.StringPtrInput
	AdministratorLoginPassword pulumi.StringPtrInput
	AvailabilityZone           pulumi.StringPtrInput
	BackupRetentionDays        pulumi.IntPtrInput
	CitusVersion               pulumi.StringPtrInput
	CreateMode                 pulumi.StringPtrInput
	DelegatedSubnetArguments   ServerGroupPropertiesDelegatedSubnetArgumentsPtrInput
	EnableMx                   pulumi.BoolPtrInput
	EnableShardsOnCoordinator  pulumi.BoolPtrInput
	EnableZfs                  pulumi.BoolPtrInput
	Location                   pulumi.StringPtrInput
	MaintenanceWindow          MaintenanceWindowPtrInput
	PointInTimeUTC             pulumi.StringPtrInput
	PostgresqlVersion          pulumi.StringPtrInput
	PrivateDnsZoneArguments    ServerGroupPropertiesPrivateDnsZoneArgumentsPtrInput
	ResourceGroupName          pulumi.StringInput
	ServerGroupName            pulumi.StringPtrInput
	ServerRoleGroups           ServerRoleGroupArrayInput
	SourceLocation             pulumi.StringPtrInput
	SourceResourceGroupName    pulumi.StringPtrInput
	SourceServerGroupName      pulumi.StringPtrInput
	SourceSubscriptionId       pulumi.StringPtrInput
	StandbyAvailabilityZone    pulumi.StringPtrInput
	Tags                       pulumi.StringMapInput
}

func (ServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupArgs)(nil)).Elem()
}

type ServerGroupInput interface {
	pulumi.Input

	ToServerGroupOutput() ServerGroupOutput
	ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput
}

func (*ServerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (i *ServerGroup) ToServerGroupOutput() ServerGroupOutput {
	return i.ToServerGroupOutputWithContext(context.Background())
}

func (i *ServerGroup) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupOutput)
}

type ServerGroupOutput struct{ *pulumi.OutputState }

func (ServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (o ServerGroupOutput) ToServerGroupOutput() ServerGroupOutput {
	return o
}

func (o ServerGroupOutput) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return o
}

func (o ServerGroupOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringPtrOutput { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

func (o ServerGroupOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o ServerGroupOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.IntPtrOutput { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o ServerGroupOutput) CitusVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringPtrOutput { return v.CitusVersion }).(pulumi.StringPtrOutput)
}

func (o ServerGroupOutput) DelegatedSubnetArguments() ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
	return o.ApplyT(func(v *ServerGroup) ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
		return v.DelegatedSubnetArguments
	}).(ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput)
}

func (o ServerGroupOutput) EarliestRestoreTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.EarliestRestoreTime }).(pulumi.StringOutput)
}

func (o ServerGroupOutput) EnableMx() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.BoolPtrOutput { return v.EnableMx }).(pulumi.BoolPtrOutput)
}

func (o ServerGroupOutput) EnableShardsOnCoordinator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.BoolPtrOutput { return v.EnableShardsOnCoordinator }).(pulumi.BoolPtrOutput)
}

func (o ServerGroupOutput) EnableZfs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.BoolPtrOutput { return v.EnableZfs }).(pulumi.BoolPtrOutput)
}

func (o ServerGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServerGroupOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v *ServerGroup) MaintenanceWindowResponsePtrOutput { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

func (o ServerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerGroupOutput) PostgresqlVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringPtrOutput { return v.PostgresqlVersion }).(pulumi.StringPtrOutput)
}

func (o ServerGroupOutput) PrivateDnsZoneArguments() ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
	return o.ApplyT(func(v *ServerGroup) ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
		return v.PrivateDnsZoneArguments
	}).(ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput)
}

func (o ServerGroupOutput) ReadReplicas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringArrayOutput { return v.ReadReplicas }).(pulumi.StringArrayOutput)
}

func (o ServerGroupOutput) ResourceProviderType() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.ResourceProviderType }).(pulumi.StringOutput)
}

func (o ServerGroupOutput) ServerRoleGroups() ServerRoleGroupResponseArrayOutput {
	return o.ApplyT(func(v *ServerGroup) ServerRoleGroupResponseArrayOutput { return v.ServerRoleGroups }).(ServerRoleGroupResponseArrayOutput)
}

func (o ServerGroupOutput) SourceServerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.SourceServerGroup }).(pulumi.StringOutput)
}

func (o ServerGroupOutput) StandbyAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringPtrOutput { return v.StandbyAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o ServerGroupOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ServerGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ServerGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ServerGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServerGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerGroupOutput{})
}
