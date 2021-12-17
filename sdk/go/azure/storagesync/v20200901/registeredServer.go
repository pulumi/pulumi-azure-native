


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegisteredServer struct {
	pulumi.CustomResourceState

	AgentVersion               pulumi.StringPtrOutput `pulumi:"agentVersion"`
	AgentVersionExpirationDate pulumi.StringOutput    `pulumi:"agentVersionExpirationDate"`
	AgentVersionStatus         pulumi.StringOutput    `pulumi:"agentVersionStatus"`
	ClusterId                  pulumi.StringPtrOutput `pulumi:"clusterId"`
	ClusterName                pulumi.StringPtrOutput `pulumi:"clusterName"`
	DiscoveryEndpointUri       pulumi.StringPtrOutput `pulumi:"discoveryEndpointUri"`
	FriendlyName               pulumi.StringPtrOutput `pulumi:"friendlyName"`
	LastHeartBeat              pulumi.StringPtrOutput `pulumi:"lastHeartBeat"`
	LastOperationName          pulumi.StringPtrOutput `pulumi:"lastOperationName"`
	LastWorkflowId             pulumi.StringPtrOutput `pulumi:"lastWorkflowId"`
	ManagementEndpointUri      pulumi.StringPtrOutput `pulumi:"managementEndpointUri"`
	MonitoringConfiguration    pulumi.StringPtrOutput `pulumi:"monitoringConfiguration"`
	MonitoringEndpointUri      pulumi.StringPtrOutput `pulumi:"monitoringEndpointUri"`
	Name                       pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState          pulumi.StringPtrOutput `pulumi:"provisioningState"`
	ResourceLocation           pulumi.StringPtrOutput `pulumi:"resourceLocation"`
	ServerCertificate          pulumi.StringPtrOutput `pulumi:"serverCertificate"`
	ServerId                   pulumi.StringPtrOutput `pulumi:"serverId"`
	ServerManagementErrorCode  pulumi.IntPtrOutput    `pulumi:"serverManagementErrorCode"`
	ServerName                 pulumi.StringOutput    `pulumi:"serverName"`
	ServerOSVersion            pulumi.StringPtrOutput `pulumi:"serverOSVersion"`
	ServerRole                 pulumi.StringPtrOutput `pulumi:"serverRole"`
	ServiceLocation            pulumi.StringPtrOutput `pulumi:"serviceLocation"`
	StorageSyncServiceUid      pulumi.StringPtrOutput `pulumi:"storageSyncServiceUid"`
	Type                       pulumi.StringOutput    `pulumi:"type"`
}


func NewRegisteredServer(ctx *pulumi.Context,
	name string, args *RegisteredServerArgs, opts ...pulumi.ResourceOption) (*RegisteredServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageSyncServiceName == nil {
		return nil, errors.New("invalid value for required argument 'StorageSyncServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagesync:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20170605preview:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180402:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180701:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20181001:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190201:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190301:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190601:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20191001:RegisteredServer"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200301:RegisteredServer"),
		},
	})
	opts = append(opts, aliases)
	var resource RegisteredServer
	err := ctx.RegisterResource("azure-native:storagesync/v20200901:RegisteredServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegisteredServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegisteredServerState, opts ...pulumi.ResourceOption) (*RegisteredServer, error) {
	var resource RegisteredServer
	err := ctx.ReadResource("azure-native:storagesync/v20200901:RegisteredServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registeredServerState struct {
}

type RegisteredServerState struct {
}

func (RegisteredServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredServerState)(nil)).Elem()
}

type registeredServerArgs struct {
	AgentVersion           *string `pulumi:"agentVersion"`
	ClusterId              *string `pulumi:"clusterId"`
	ClusterName            *string `pulumi:"clusterName"`
	FriendlyName           *string `pulumi:"friendlyName"`
	LastHeartBeat          *string `pulumi:"lastHeartBeat"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	ServerCertificate      *string `pulumi:"serverCertificate"`
	ServerId               *string `pulumi:"serverId"`
	ServerOSVersion        *string `pulumi:"serverOSVersion"`
	ServerRole             *string `pulumi:"serverRole"`
	StorageSyncServiceName string  `pulumi:"storageSyncServiceName"`
}


type RegisteredServerArgs struct {
	AgentVersion           pulumi.StringPtrInput
	ClusterId              pulumi.StringPtrInput
	ClusterName            pulumi.StringPtrInput
	FriendlyName           pulumi.StringPtrInput
	LastHeartBeat          pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	ServerCertificate      pulumi.StringPtrInput
	ServerId               pulumi.StringPtrInput
	ServerOSVersion        pulumi.StringPtrInput
	ServerRole             pulumi.StringPtrInput
	StorageSyncServiceName pulumi.StringInput
}

func (RegisteredServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredServerArgs)(nil)).Elem()
}

type RegisteredServerInput interface {
	pulumi.Input

	ToRegisteredServerOutput() RegisteredServerOutput
	ToRegisteredServerOutputWithContext(ctx context.Context) RegisteredServerOutput
}

func (*RegisteredServer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegisteredServer)(nil)).Elem()
}

func (i *RegisteredServer) ToRegisteredServerOutput() RegisteredServerOutput {
	return i.ToRegisteredServerOutputWithContext(context.Background())
}

func (i *RegisteredServer) ToRegisteredServerOutputWithContext(ctx context.Context) RegisteredServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegisteredServerOutput)
}

type RegisteredServerOutput struct{ *pulumi.OutputState }

func (RegisteredServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegisteredServer)(nil)).Elem()
}

func (o RegisteredServerOutput) ToRegisteredServerOutput() RegisteredServerOutput {
	return o
}

func (o RegisteredServerOutput) ToRegisteredServerOutputWithContext(ctx context.Context) RegisteredServerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegisteredServerOutput{})
}
