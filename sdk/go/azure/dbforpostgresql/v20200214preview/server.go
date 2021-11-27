


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
