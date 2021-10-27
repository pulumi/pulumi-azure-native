


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionByWorkspace struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnectionByWorkspace(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionByWorkspaceArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByWorkspace, error) {
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
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210401preview:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210903preview:PrivateEndpointConnectionByWorkspace"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionByWorkspace
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20210401preview:PrivateEndpointConnectionByWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionByWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionByWorkspaceState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByWorkspace, error) {
	var resource PrivateEndpointConnectionByWorkspace
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20210401preview:PrivateEndpointConnectionByWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionByWorkspaceState struct {
}

type PrivateEndpointConnectionByWorkspaceState struct {
}

func (PrivateEndpointConnectionByWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByWorkspaceState)(nil)).Elem()
}

type privateEndpointConnectionByWorkspaceArgs struct {
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	WorkspaceName                     string                            `pulumi:"workspaceName"`
}


type PrivateEndpointConnectionByWorkspaceArgs struct {
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	WorkspaceName                     pulumi.StringInput
}

func (PrivateEndpointConnectionByWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByWorkspaceArgs)(nil)).Elem()
}

type PrivateEndpointConnectionByWorkspaceInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionByWorkspaceOutput() PrivateEndpointConnectionByWorkspaceOutput
	ToPrivateEndpointConnectionByWorkspaceOutputWithContext(ctx context.Context) PrivateEndpointConnectionByWorkspaceOutput
}

func (*PrivateEndpointConnectionByWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionByWorkspace)(nil))
}

func (i *PrivateEndpointConnectionByWorkspace) ToPrivateEndpointConnectionByWorkspaceOutput() PrivateEndpointConnectionByWorkspaceOutput {
	return i.ToPrivateEndpointConnectionByWorkspaceOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionByWorkspace) ToPrivateEndpointConnectionByWorkspaceOutputWithContext(ctx context.Context) PrivateEndpointConnectionByWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionByWorkspaceOutput)
}

type PrivateEndpointConnectionByWorkspaceOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionByWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionByWorkspace)(nil))
}

func (o PrivateEndpointConnectionByWorkspaceOutput) ToPrivateEndpointConnectionByWorkspaceOutput() PrivateEndpointConnectionByWorkspaceOutput {
	return o
}

func (o PrivateEndpointConnectionByWorkspaceOutput) ToPrivateEndpointConnectionByWorkspaceOutputWithContext(ctx context.Context) PrivateEndpointConnectionByWorkspaceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionByWorkspaceInput)(nil)).Elem(), &PrivateEndpointConnectionByWorkspace{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionByWorkspaceOutput{})
}
