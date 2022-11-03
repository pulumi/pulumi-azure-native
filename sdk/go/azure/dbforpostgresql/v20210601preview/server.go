


package v20210601preview

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
			Type: pulumi.String("azure-native:dbforpostgresql/v20210601:Server"),
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
	})
	opts = append(opts, aliases)
	var resource Server
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20210601preview:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20210601preview:Server", name, id, state, &resource, opts...)
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

func (o ServerOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) Backup() BackupResponsePtrOutput {
	return o.ApplyT(func(v *Server) BackupResponsePtrOutput { return v.Backup }).(BackupResponsePtrOutput)
}

func (o ServerOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

func (o ServerOutput) HighAvailability() HighAvailabilityResponsePtrOutput {
	return o.ApplyT(func(v *Server) HighAvailabilityResponsePtrOutput { return v.HighAvailability }).(HighAvailabilityResponsePtrOutput)
}

func (o ServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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

func (o ServerOutput) Network() NetworkResponsePtrOutput {
	return o.ApplyT(func(v *Server) NetworkResponsePtrOutput { return v.Network }).(NetworkResponsePtrOutput)
}

func (o ServerOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Server) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ServerOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ServerOutput) Storage() StorageResponsePtrOutput {
	return o.ApplyT(func(v *Server) StorageResponsePtrOutput { return v.Storage }).(StorageResponsePtrOutput)
}

func (o ServerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Server) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
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
