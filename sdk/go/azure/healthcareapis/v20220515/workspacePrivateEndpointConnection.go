


package v20220515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkspacePrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewWorkspacePrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *WorkspacePrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*WorkspacePrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthcareapis:WorkspacePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20211101:WorkspacePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:WorkspacePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220601:WorkspacePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221001preview:WorkspacePrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkspacePrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:healthcareapis/v20220515:WorkspacePrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspacePrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspacePrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*WorkspacePrivateEndpointConnection, error) {
	var resource WorkspacePrivateEndpointConnection
	err := ctx.ReadResource("azure-native:healthcareapis/v20220515:WorkspacePrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workspacePrivateEndpointConnectionState struct {
}

type WorkspacePrivateEndpointConnectionState struct {
}

func (WorkspacePrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspacePrivateEndpointConnectionState)(nil)).Elem()
}

type workspacePrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	WorkspaceName                     string                            `pulumi:"workspaceName"`
}


type WorkspacePrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	WorkspaceName                     pulumi.StringInput
}

func (WorkspacePrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspacePrivateEndpointConnectionArgs)(nil)).Elem()
}

type WorkspacePrivateEndpointConnectionInput interface {
	pulumi.Input

	ToWorkspacePrivateEndpointConnectionOutput() WorkspacePrivateEndpointConnectionOutput
	ToWorkspacePrivateEndpointConnectionOutputWithContext(ctx context.Context) WorkspacePrivateEndpointConnectionOutput
}

func (*WorkspacePrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePrivateEndpointConnection)(nil)).Elem()
}

func (i *WorkspacePrivateEndpointConnection) ToWorkspacePrivateEndpointConnectionOutput() WorkspacePrivateEndpointConnectionOutput {
	return i.ToWorkspacePrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *WorkspacePrivateEndpointConnection) ToWorkspacePrivateEndpointConnectionOutputWithContext(ctx context.Context) WorkspacePrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePrivateEndpointConnectionOutput)
}

type WorkspacePrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (WorkspacePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePrivateEndpointConnection)(nil)).Elem()
}

func (o WorkspacePrivateEndpointConnectionOutput) ToWorkspacePrivateEndpointConnectionOutput() WorkspacePrivateEndpointConnectionOutput {
	return o
}

func (o WorkspacePrivateEndpointConnectionOutput) ToWorkspacePrivateEndpointConnectionOutputWithContext(ctx context.Context) WorkspacePrivateEndpointConnectionOutput {
	return o
}

func (o WorkspacePrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspacePrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o WorkspacePrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o WorkspacePrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkspacePrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WorkspacePrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspacePrivateEndpointConnectionOutput{})
}
