


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
			Type: pulumi.String("azure-nextgen:dbformysql/v20200701preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20200701privatepreview:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:dbformysql/v20200701privatepreview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20210501:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:dbformysql/v20210501:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformysql/v20210501preview:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:dbformysql/v20210501preview:Server"),
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
	return reflect.TypeOf((*Server)(nil))
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Server)(nil))
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
