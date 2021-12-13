


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Server struct {
	pulumi.CustomResourceState

	AdministratorLogin       pulumi.StringPtrOutput             `pulumi:"administratorLogin"`
	AvailabilityZone         pulumi.StringPtrOutput             `pulumi:"availabilityZone"`
	Backup                   BackupResponsePtrOutput            `pulumi:"backup"`
	FullyQualifiedDomainName pulumi.StringOutput                `pulumi:"fullyQualifiedDomainName"`
	HighAvailability         HighAvailabilityResponsePtrOutput  `pulumi:"highAvailability"`
	Location                 pulumi.StringOutput                `pulumi:"location"`
	MaintenanceWindow        MaintenanceWindowResponsePtrOutput `pulumi:"maintenanceWindow"`
	MinorVersion             pulumi.StringOutput                `pulumi:"minorVersion"`
	Name                     pulumi.StringOutput                `pulumi:"name"`
	Network                  NetworkResponsePtrOutput           `pulumi:"network"`
	Sku                      SkuResponsePtrOutput               `pulumi:"sku"`
	State                    pulumi.StringOutput                `pulumi:"state"`
	Storage                  StorageResponsePtrOutput           `pulumi:"storage"`
	SystemData               SystemDataResponseOutput           `pulumi:"systemData"`
	Tags                     pulumi.StringMapOutput             `pulumi:"tags"`
	Type                     pulumi.StringOutput                `pulumi:"type"`
	Version                  pulumi.StringPtrOutput             `pulumi:"version"`
}


func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.AvailabilityZone) {
		args.AvailabilityZone = pulumi.StringPtr("")
	}
	if args.Backup != nil {
		args.Backup = args.Backup.ToBackupPtrOutput().ApplyT(func(v *Backup) *Backup { return v.Defaults() }).(BackupPtrOutput)
	}
	if args.HighAvailability != nil {
		args.HighAvailability = args.HighAvailability.ToHighAvailabilityPtrOutput().ApplyT(func(v *HighAvailability) *HighAvailability { return v.Defaults() }).(HighAvailabilityPtrOutput)
	}
	if args.MaintenanceWindow != nil {
		args.MaintenanceWindow = args.MaintenanceWindow.ToMaintenanceWindowPtrOutput().ApplyT(func(v *MaintenanceWindow) *MaintenanceWindow { return v.Defaults() }).(MaintenanceWindowPtrOutput)
	}
	if args.Network != nil {
		args.Network = args.Network.ToNetworkPtrOutput().ApplyT(func(v *Network) *Network { return v.Defaults() }).(NetworkPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20200214preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20200214privatepreview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210410privatepreview:Server"),
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
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20210601:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20210601:Server", name, id, state, &resource, opts...)
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
	AdministratorLogin         *string            `pulumi:"administratorLogin"`
	AdministratorLoginPassword *string            `pulumi:"administratorLoginPassword"`
	AvailabilityZone           *string            `pulumi:"availabilityZone"`
	Backup                     *Backup            `pulumi:"backup"`
	CreateMode                 *string            `pulumi:"createMode"`
	HighAvailability           *HighAvailability  `pulumi:"highAvailability"`
	Location                   *string            `pulumi:"location"`
	MaintenanceWindow          *MaintenanceWindow `pulumi:"maintenanceWindow"`
	Network                    *Network           `pulumi:"network"`
	PointInTimeUTC             *string            `pulumi:"pointInTimeUTC"`
	ResourceGroupName          string             `pulumi:"resourceGroupName"`
	ServerName                 *string            `pulumi:"serverName"`
	Sku                        *Sku               `pulumi:"sku"`
	SourceServerResourceId     *string            `pulumi:"sourceServerResourceId"`
	Storage                    *Storage           `pulumi:"storage"`
	Tags                       map[string]string  `pulumi:"tags"`
	Version                    *string            `pulumi:"version"`
}


type ServerArgs struct {
	AdministratorLogin         pulumi.StringPtrInput
	AdministratorLoginPassword pulumi.StringPtrInput
	AvailabilityZone           pulumi.StringPtrInput
	Backup                     BackupPtrInput
	CreateMode                 pulumi.StringPtrInput
	HighAvailability           HighAvailabilityPtrInput
	Location                   pulumi.StringPtrInput
	MaintenanceWindow          MaintenanceWindowPtrInput
	Network                    NetworkPtrInput
	PointInTimeUTC             pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	ServerName                 pulumi.StringPtrInput
	Sku                        SkuPtrInput
	SourceServerResourceId     pulumi.StringPtrInput
	Storage                    StoragePtrInput
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

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
