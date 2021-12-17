


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Server struct {
	pulumi.CustomResourceState

	AdministratorLogin            pulumi.StringPtrOutput                             `pulumi:"administratorLogin"`
	Administrators                ServerExternalAdministratorResponsePtrOutput       `pulumi:"administrators"`
	FederatedClientId             pulumi.StringPtrOutput                             `pulumi:"federatedClientId"`
	FullyQualifiedDomainName      pulumi.StringOutput                                `pulumi:"fullyQualifiedDomainName"`
	Identity                      ResourceIdentityResponsePtrOutput                  `pulumi:"identity"`
	KeyId                         pulumi.StringPtrOutput                             `pulumi:"keyId"`
	Kind                          pulumi.StringOutput                                `pulumi:"kind"`
	Location                      pulumi.StringOutput                                `pulumi:"location"`
	MinimalTlsVersion             pulumi.StringPtrOutput                             `pulumi:"minimalTlsVersion"`
	Name                          pulumi.StringOutput                                `pulumi:"name"`
	PrimaryUserAssignedIdentityId pulumi.StringPtrOutput                             `pulumi:"primaryUserAssignedIdentityId"`
	PrivateEndpointConnections    ServerPrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess           pulumi.StringPtrOutput                             `pulumi:"publicNetworkAccess"`
	RestrictOutboundNetworkAccess pulumi.StringPtrOutput                             `pulumi:"restrictOutboundNetworkAccess"`
	State                         pulumi.StringOutput                                `pulumi:"state"`
	Tags                          pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                          pulumi.StringOutput                                `pulumi:"type"`
	Version                       pulumi.StringPtrOutput                             `pulumi:"version"`
	WorkspaceFeature              pulumi.StringOutput                                `pulumi:"workspaceFeature"`
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
			Type: pulumi.String("azure-native:sql:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:Server"),
		},
	})
	opts = append(opts, aliases)
	var resource Server
	err := ctx.RegisterResource("azure-native:sql/v20210501preview:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure-native:sql/v20210501preview:Server", name, id, state, &resource, opts...)
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
	AdministratorLogin            *string                      `pulumi:"administratorLogin"`
	AdministratorLoginPassword    *string                      `pulumi:"administratorLoginPassword"`
	Administrators                *ServerExternalAdministrator `pulumi:"administrators"`
	FederatedClientId             *string                      `pulumi:"federatedClientId"`
	Identity                      *ResourceIdentity            `pulumi:"identity"`
	KeyId                         *string                      `pulumi:"keyId"`
	Location                      *string                      `pulumi:"location"`
	MinimalTlsVersion             *string                      `pulumi:"minimalTlsVersion"`
	PrimaryUserAssignedIdentityId *string                      `pulumi:"primaryUserAssignedIdentityId"`
	PublicNetworkAccess           *string                      `pulumi:"publicNetworkAccess"`
	ResourceGroupName             string                       `pulumi:"resourceGroupName"`
	RestrictOutboundNetworkAccess *string                      `pulumi:"restrictOutboundNetworkAccess"`
	ServerName                    *string                      `pulumi:"serverName"`
	Tags                          map[string]string            `pulumi:"tags"`
	Version                       *string                      `pulumi:"version"`
}


type ServerArgs struct {
	AdministratorLogin            pulumi.StringPtrInput
	AdministratorLoginPassword    pulumi.StringPtrInput
	Administrators                ServerExternalAdministratorPtrInput
	FederatedClientId             pulumi.StringPtrInput
	Identity                      ResourceIdentityPtrInput
	KeyId                         pulumi.StringPtrInput
	Location                      pulumi.StringPtrInput
	MinimalTlsVersion             pulumi.StringPtrInput
	PrimaryUserAssignedIdentityId pulumi.StringPtrInput
	PublicNetworkAccess           pulumi.StringPtrInput
	ResourceGroupName             pulumi.StringInput
	RestrictOutboundNetworkAccess pulumi.StringPtrInput
	ServerName                    pulumi.StringPtrInput
	Tags                          pulumi.StringMapInput
	Version                       pulumi.StringPtrInput
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
