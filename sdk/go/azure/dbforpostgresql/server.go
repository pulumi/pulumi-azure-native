


package dbforpostgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Server struct {
	pulumi.CustomResourceState

	AdministratorLogin         pulumi.StringPtrOutput                             `pulumi:"administratorLogin"`
	ByokEnforcement            pulumi.StringOutput                                `pulumi:"byokEnforcement"`
	EarliestRestoreDate        pulumi.StringPtrOutput                             `pulumi:"earliestRestoreDate"`
	FullyQualifiedDomainName   pulumi.StringPtrOutput                             `pulumi:"fullyQualifiedDomainName"`
	Identity                   ResourceIdentityResponsePtrOutput                  `pulumi:"identity"`
	InfrastructureEncryption   pulumi.StringPtrOutput                             `pulumi:"infrastructureEncryption"`
	Location                   pulumi.StringOutput                                `pulumi:"location"`
	MasterServerId             pulumi.StringPtrOutput                             `pulumi:"masterServerId"`
	MinimalTlsVersion          pulumi.StringPtrOutput                             `pulumi:"minimalTlsVersion"`
	Name                       pulumi.StringOutput                                `pulumi:"name"`
	PrivateEndpointConnections ServerPrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess        pulumi.StringPtrOutput                             `pulumi:"publicNetworkAccess"`
	ReplicaCapacity            pulumi.IntPtrOutput                                `pulumi:"replicaCapacity"`
	ReplicationRole            pulumi.StringPtrOutput                             `pulumi:"replicationRole"`
	Sku                        SkuResponsePtrOutput                               `pulumi:"sku"`
	SslEnforcement             pulumi.StringPtrOutput                             `pulumi:"sslEnforcement"`
	StorageProfile             StorageProfileResponsePtrOutput                    `pulumi:"storageProfile"`
	Tags                       pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                       pulumi.StringOutput                                `pulumi:"type"`
	UserVisibleState           pulumi.StringPtrOutput                             `pulumi:"userVisibleState"`
	Version                    pulumi.StringPtrOutput                             `pulumi:"version"`
}


func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20171201:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20171201preview:Server"),
		},
	})
	opts = append(opts, aliases)
	var resource Server
	err := ctx.RegisterResource("azure-native:dbforpostgresql:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure-native:dbforpostgresql:Server", name, id, state, &resource, opts...)
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
	Identity          *ResourceIdentity `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	Properties        interface{}       `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ServerName        *string           `pulumi:"serverName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type ServerArgs struct {
	Identity          ResourceIdentityPtrInput
	Location          pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringPtrInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
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

func (o ServerOutput) ByokEnforcement() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ByokEnforcement }).(pulumi.StringOutput)
}

func (o ServerOutput) EarliestRestoreDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.EarliestRestoreDate }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) FullyQualifiedDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.FullyQualifiedDomainName }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Server) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o ServerOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServerOutput) MasterServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.MasterServerId }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerOutput) PrivateEndpointConnections() ServerPrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Server) ServerPrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(ServerPrivateEndpointConnectionResponseArrayOutput)
}

func (o ServerOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) ReplicaCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.IntPtrOutput { return v.ReplicaCapacity }).(pulumi.IntPtrOutput)
}

func (o ServerOutput) ReplicationRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.ReplicationRole }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Server) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ServerOutput) SslEnforcement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SslEnforcement }).(pulumi.StringPtrOutput)
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

func (o ServerOutput) UserVisibleState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.UserVisibleState }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
