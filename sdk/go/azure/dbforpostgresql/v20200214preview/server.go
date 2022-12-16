


package v20200214preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Server struct {
	pulumi.CustomResourceState

	AdministratorLogin       pulumi.StringPtrOutput                                    `pulumi:"administratorLogin"`
	AvailabilityZone         pulumi.StringPtrOutput                                    `pulumi:"availabilityZone"`
	ByokEnforcement          pulumi.StringOutput                                       `pulumi:"byokEnforcement"`
	DelegatedSubnetArguments ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput `pulumi:"delegatedSubnetArguments"`
	DisplayName              pulumi.StringPtrOutput                                    `pulumi:"displayName"`
	EarliestRestoreDate      pulumi.StringOutput                                       `pulumi:"earliestRestoreDate"`
	FullyQualifiedDomainName pulumi.StringOutput                                       `pulumi:"fullyQualifiedDomainName"`
	HaEnabled                pulumi.StringPtrOutput                                    `pulumi:"haEnabled"`
	HaState                  pulumi.StringOutput                                       `pulumi:"haState"`
	Identity                 IdentityResponsePtrOutput                                 `pulumi:"identity"`
	Location                 pulumi.StringOutput                                       `pulumi:"location"`
	LogBackupStorageSku      pulumi.StringPtrOutput                                    `pulumi:"logBackupStorageSku"`
	MaintenanceWindow        MaintenanceWindowResponsePtrOutput                        `pulumi:"maintenanceWindow"`
	MinorVersion             pulumi.StringOutput                                       `pulumi:"minorVersion"`
	Name                     pulumi.StringOutput                                       `pulumi:"name"`
	PointInTimeUTC           pulumi.StringPtrOutput                                    `pulumi:"pointInTimeUTC"`
	PrivateDnsZoneArguments  ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput  `pulumi:"privateDnsZoneArguments"`
	PublicNetworkAccess      pulumi.StringOutput                                       `pulumi:"publicNetworkAccess"`
	Sku                      SkuResponsePtrOutput                                      `pulumi:"sku"`
	SourceResourceGroupName  pulumi.StringPtrOutput                                    `pulumi:"sourceResourceGroupName"`
	SourceServerName         pulumi.StringPtrOutput                                    `pulumi:"sourceServerName"`
	SourceSubscriptionId     pulumi.StringPtrOutput                                    `pulumi:"sourceSubscriptionId"`
	StandbyAvailabilityZone  pulumi.StringOutput                                       `pulumi:"standbyAvailabilityZone"`
	StandbyCount             pulumi.IntPtrOutput                                       `pulumi:"standbyCount"`
	State                    pulumi.StringOutput                                       `pulumi:"state"`
	StorageProfile           StorageProfileResponsePtrOutput                           `pulumi:"storageProfile"`
	Tags                     pulumi.StringMapOutput                                    `pulumi:"tags"`
	Type                     pulumi.StringOutput                                       `pulumi:"type"`
	Version                  pulumi.StringPtrOutput                                    `pulumi:"version"`
}


func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20200214privatepreview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210410privatepreview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210601:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210601preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210615privatepreview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20220120preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20220308preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20221201:Server"),
		},
	})
	opts = append(opts, aliases)
	var resource Server
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20200214preview:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20200214preview:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverState struct {
}

type ServerState struct {
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	AdministratorLogin         *string                                   `pulumi:"administratorLogin"`
	AdministratorLoginPassword *string                                   `pulumi:"administratorLoginPassword"`
	AvailabilityZone           *string                                   `pulumi:"availabilityZone"`
	CreateMode                 *string                                   `pulumi:"createMode"`
	DelegatedSubnetArguments   *ServerPropertiesDelegatedSubnetArguments `pulumi:"delegatedSubnetArguments"`
	DisplayName                *string                                   `pulumi:"displayName"`
	HaEnabled                  *HAEnabledEnum                            `pulumi:"haEnabled"`
	Identity                   *Identity                                 `pulumi:"identity"`
	Location                   *string                                   `pulumi:"location"`
	LogBackupStorageSku        *string                                   `pulumi:"logBackupStorageSku"`
	MaintenanceWindow          *MaintenanceWindow                        `pulumi:"maintenanceWindow"`
	PointInTimeUTC             *string                                   `pulumi:"pointInTimeUTC"`
	PrivateDnsZoneArguments    *ServerPropertiesPrivateDnsZoneArguments  `pulumi:"privateDnsZoneArguments"`
	ResourceGroupName          string                                    `pulumi:"resourceGroupName"`
	ServerName                 *string                                   `pulumi:"serverName"`
	Sku                        *Sku                                      `pulumi:"sku"`
	SourceResourceGroupName    *string                                   `pulumi:"sourceResourceGroupName"`
	SourceServerName           *string                                   `pulumi:"sourceServerName"`
	SourceSubscriptionId       *string                                   `pulumi:"sourceSubscriptionId"`
	StandbyCount               *int                                      `pulumi:"standbyCount"`
	StorageProfile             *StorageProfile                           `pulumi:"storageProfile"`
	Tags                       map[string]string                         `pulumi:"tags"`
	Version                    *string                                   `pulumi:"version"`
}


type ServerArgs struct {
	AdministratorLogin         pulumi.StringPtrInput
	AdministratorLoginPassword pulumi.StringPtrInput
	AvailabilityZone           pulumi.StringPtrInput
	CreateMode                 pulumi.StringPtrInput
	DelegatedSubnetArguments   ServerPropertiesDelegatedSubnetArgumentsPtrInput
	DisplayName                pulumi.StringPtrInput
	HaEnabled                  HAEnabledEnumPtrInput
	Identity                   IdentityPtrInput
	Location                   pulumi.StringPtrInput
	LogBackupStorageSku        pulumi.StringPtrInput
	MaintenanceWindow          MaintenanceWindowPtrInput
	PointInTimeUTC             pulumi.StringPtrInput
	PrivateDnsZoneArguments    ServerPropertiesPrivateDnsZoneArgumentsPtrInput
	ResourceGroupName          pulumi.StringInput
	ServerName                 pulumi.StringPtrInput
	Sku                        SkuPtrInput
	SourceResourceGroupName    pulumi.StringPtrInput
	SourceServerName           pulumi.StringPtrInput
	SourceSubscriptionId       pulumi.StringPtrInput
	StandbyCount               pulumi.IntPtrInput
	StorageProfile             StorageProfilePtrInput
	Tags                       pulumi.StringMapInput
	Version                    pulumi.StringPtrInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

func (o ServerOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) ByokEnforcement() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ByokEnforcement }).(pulumi.StringOutput)
}

func (o ServerOutput) DelegatedSubnetArguments() ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
	return o.ApplyT(func(v *Server) ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
		return v.DelegatedSubnetArguments
	}).(ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput)
}

func (o ServerOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) EarliestRestoreDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.EarliestRestoreDate }).(pulumi.StringOutput)
}

func (o ServerOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

func (o ServerOutput) HaEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.HaEnabled }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) HaState() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.HaState }).(pulumi.StringOutput)
}

func (o ServerOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Server) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o ServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServerOutput) LogBackupStorageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.LogBackupStorageSku }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v *Server) MaintenanceWindowResponsePtrOutput { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

func (o ServerOutput) MinorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.MinorVersion }).(pulumi.StringOutput)
}

func (o ServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerOutput) PointInTimeUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.PointInTimeUTC }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) PrivateDnsZoneArguments() ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
	return o.ApplyT(func(v *Server) ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
		return v.PrivateDnsZoneArguments
	}).(ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput)
}

func (o ServerOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

func (o ServerOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Server) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ServerOutput) SourceResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SourceResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) SourceServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SourceServerName }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) SourceSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SourceSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) StandbyAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.StandbyAvailabilityZone }).(pulumi.StringOutput)
}

func (o ServerOutput) StandbyCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.IntPtrOutput { return v.StandbyCount }).(pulumi.IntPtrOutput)
}

func (o ServerOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ServerOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *Server) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o ServerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Server) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ServerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
