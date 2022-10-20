


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Server struct {
	pulumi.CustomResourceState

	AdministratorLogin       pulumi.StringPtrOutput                    `pulumi:"administratorLogin"`
	AvailabilityZone         pulumi.StringPtrOutput                    `pulumi:"availabilityZone"`
	ByokEnforcement          pulumi.StringOutput                       `pulumi:"byokEnforcement"`
	DelegatedSubnetArguments DelegatedSubnetArgumentsResponsePtrOutput `pulumi:"delegatedSubnetArguments"`
	EarliestRestoreDate      pulumi.StringOutput                       `pulumi:"earliestRestoreDate"`
	FullyQualifiedDomainName pulumi.StringOutput                       `pulumi:"fullyQualifiedDomainName"`
	HaEnabled                pulumi.StringPtrOutput                    `pulumi:"haEnabled"`
	HaState                  pulumi.StringOutput                       `pulumi:"haState"`
	Identity                 IdentityResponsePtrOutput                 `pulumi:"identity"`
	Location                 pulumi.StringOutput                       `pulumi:"location"`
	MaintenanceWindow        MaintenanceWindowResponsePtrOutput        `pulumi:"maintenanceWindow"`
	Name                     pulumi.StringOutput                       `pulumi:"name"`
	PrivateDnsZoneArguments  PrivateDnsZoneArgumentsResponsePtrOutput  `pulumi:"privateDnsZoneArguments"`
	PublicNetworkAccess      pulumi.StringOutput                       `pulumi:"publicNetworkAccess"`
	ReplicaCapacity          pulumi.IntOutput                          `pulumi:"replicaCapacity"`
	ReplicationRole          pulumi.StringPtrOutput                    `pulumi:"replicationRole"`
	Sku                      SkuResponsePtrOutput                      `pulumi:"sku"`
	SourceServerId           pulumi.StringPtrOutput                    `pulumi:"sourceServerId"`
	SslEnforcement           pulumi.StringPtrOutput                    `pulumi:"sslEnforcement"`
	StandbyAvailabilityZone  pulumi.StringOutput                       `pulumi:"standbyAvailabilityZone"`
	State                    pulumi.StringOutput                       `pulumi:"state"`
	StorageProfile           StorageProfileResponsePtrOutput           `pulumi:"storageProfile"`
	Tags                     pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                     pulumi.StringOutput                       `pulumi:"type"`
	Version                  pulumi.StringPtrOutput                    `pulumi:"version"`
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
			Type: pulumi.String("azure-native:dbformysql/v20200701privatepreview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20210501:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20210501preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20211201preview:Server"),
		},
	})
	opts = append(opts, aliases)
	var resource Server
	err := ctx.RegisterResource("azure-native:dbformysql/v20200701preview:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure-native:dbformysql/v20200701preview:Server", name, id, state, &resource, opts...)
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
	AdministratorLogin         *string                   `pulumi:"administratorLogin"`
	AdministratorLoginPassword *string                   `pulumi:"administratorLoginPassword"`
	AvailabilityZone           *string                   `pulumi:"availabilityZone"`
	CreateMode                 *string                   `pulumi:"createMode"`
	DelegatedSubnetArguments   *DelegatedSubnetArguments `pulumi:"delegatedSubnetArguments"`
	HaEnabled                  *string                   `pulumi:"haEnabled"`
	Identity                   *Identity                 `pulumi:"identity"`
	InfrastructureEncryption   *string                   `pulumi:"infrastructureEncryption"`
	Location                   *string                   `pulumi:"location"`
	MaintenanceWindow          *MaintenanceWindow        `pulumi:"maintenanceWindow"`
	PrivateDnsZoneArguments    *PrivateDnsZoneArguments  `pulumi:"privateDnsZoneArguments"`
	ReplicationRole            *string                   `pulumi:"replicationRole"`
	ResourceGroupName          string                    `pulumi:"resourceGroupName"`
	RestorePointInTime         *string                   `pulumi:"restorePointInTime"`
	ServerName                 *string                   `pulumi:"serverName"`
	Sku                        *Sku                      `pulumi:"sku"`
	SourceServerId             *string                   `pulumi:"sourceServerId"`
	SslEnforcement             *string                   `pulumi:"sslEnforcement"`
	StorageProfile             *StorageProfile           `pulumi:"storageProfile"`
	Tags                       map[string]string         `pulumi:"tags"`
	Version                    *string                   `pulumi:"version"`
}


type ServerArgs struct {
	AdministratorLogin         pulumi.StringPtrInput
	AdministratorLoginPassword pulumi.StringPtrInput
	AvailabilityZone           pulumi.StringPtrInput
	CreateMode                 pulumi.StringPtrInput
	DelegatedSubnetArguments   DelegatedSubnetArgumentsPtrInput
	HaEnabled                  pulumi.StringPtrInput
	Identity                   IdentityPtrInput
	InfrastructureEncryption   pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	MaintenanceWindow          MaintenanceWindowPtrInput
	PrivateDnsZoneArguments    PrivateDnsZoneArgumentsPtrInput
	ReplicationRole            pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	RestorePointInTime         pulumi.StringPtrInput
	ServerName                 pulumi.StringPtrInput
	Sku                        SkuPtrInput
	SourceServerId             pulumi.StringPtrInput
	SslEnforcement             pulumi.StringPtrInput
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

func (o ServerOutput) DelegatedSubnetArguments() DelegatedSubnetArgumentsResponsePtrOutput {
	return o.ApplyT(func(v *Server) DelegatedSubnetArgumentsResponsePtrOutput { return v.DelegatedSubnetArguments }).(DelegatedSubnetArgumentsResponsePtrOutput)
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

func (o ServerOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v *Server) MaintenanceWindowResponsePtrOutput { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

func (o ServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerOutput) PrivateDnsZoneArguments() PrivateDnsZoneArgumentsResponsePtrOutput {
	return o.ApplyT(func(v *Server) PrivateDnsZoneArgumentsResponsePtrOutput { return v.PrivateDnsZoneArguments }).(PrivateDnsZoneArgumentsResponsePtrOutput)
}

func (o ServerOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

func (o ServerOutput) ReplicaCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *Server) pulumi.IntOutput { return v.ReplicaCapacity }).(pulumi.IntOutput)
}

func (o ServerOutput) ReplicationRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.ReplicationRole }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Server) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ServerOutput) SourceServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SourceServerId }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) SslEnforcement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SslEnforcement }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) StandbyAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.StandbyAvailabilityZone }).(pulumi.StringOutput)
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
