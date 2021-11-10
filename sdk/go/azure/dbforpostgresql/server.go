


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
