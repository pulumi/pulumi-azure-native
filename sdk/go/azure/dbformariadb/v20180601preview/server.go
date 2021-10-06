


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Server struct {
	pulumi.CustomResourceState

	AdministratorLogin       pulumi.StringPtrOutput            `pulumi:"administratorLogin"`
	EarliestRestoreDate      pulumi.StringPtrOutput            `pulumi:"earliestRestoreDate"`
	FullyQualifiedDomainName pulumi.StringPtrOutput            `pulumi:"fullyQualifiedDomainName"`
	Identity                 ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	Location                 pulumi.StringOutput               `pulumi:"location"`
	MasterServerId           pulumi.StringPtrOutput            `pulumi:"masterServerId"`
	MinimalTlsVersion        pulumi.StringPtrOutput            `pulumi:"minimalTlsVersion"`
	Name                     pulumi.StringOutput               `pulumi:"name"`
	ReplicaCapacity          pulumi.IntPtrOutput               `pulumi:"replicaCapacity"`
	ReplicationRole          pulumi.StringPtrOutput            `pulumi:"replicationRole"`
	Sku                      SkuResponsePtrOutput              `pulumi:"sku"`
	SslEnforcement           pulumi.StringPtrOutput            `pulumi:"sslEnforcement"`
	StorageProfile           StorageProfileResponsePtrOutput   `pulumi:"storageProfile"`
	Tags                     pulumi.StringMapOutput            `pulumi:"tags"`
	Type                     pulumi.StringOutput               `pulumi:"type"`
	UserVisibleState         pulumi.StringPtrOutput            `pulumi:"userVisibleState"`
	Version                  pulumi.StringPtrOutput            `pulumi:"version"`
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
			Type: pulumi.String("azure-nextgen:dbformariadb/v20180601preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformariadb:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:dbformariadb:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbformariadb/v20180601:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:dbformariadb/v20180601:Server"),
		},
	})
	opts = append(opts, aliases)
	var resource Server
	err := ctx.RegisterResource("azure-native:dbformariadb/v20180601preview:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure-native:dbformariadb/v20180601preview:Server", name, id, state, &resource, opts...)
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
	Location          *string           `pulumi:"location"`
	Properties        interface{}       `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ServerName        *string           `pulumi:"serverName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type ServerArgs struct {
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
