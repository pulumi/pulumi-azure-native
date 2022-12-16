


package v20210903preview

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
			Type: pulumi.String("azure-native:desktopvirtualization:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:PrivateEndpointConnectionByWorkspace"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionByWorkspace
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20210903preview:PrivateEndpointConnectionByWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionByWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionByWorkspaceState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByWorkspace, error) {
	var resource PrivateEndpointConnectionByWorkspace
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20210903preview:PrivateEndpointConnectionByWorkspace", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**PrivateEndpointConnectionByWorkspace)(nil)).Elem()
}

func (i *PrivateEndpointConnectionByWorkspace) ToPrivateEndpointConnectionByWorkspaceOutput() PrivateEndpointConnectionByWorkspaceOutput {
	return i.ToPrivateEndpointConnectionByWorkspaceOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionByWorkspace) ToPrivateEndpointConnectionByWorkspaceOutputWithContext(ctx context.Context) PrivateEndpointConnectionByWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionByWorkspaceOutput)
}

type PrivateEndpointConnectionByWorkspaceOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionByWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByWorkspace)(nil)).Elem()
}

func (o PrivateEndpointConnectionByWorkspaceOutput) ToPrivateEndpointConnectionByWorkspaceOutput() PrivateEndpointConnectionByWorkspaceOutput {
	return o
}

func (o PrivateEndpointConnectionByWorkspaceOutput) ToPrivateEndpointConnectionByWorkspaceOutputWithContext(ctx context.Context) PrivateEndpointConnectionByWorkspaceOutput {
	return o
}

func (o PrivateEndpointConnectionByWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionByWorkspaceOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) PrivateEndpointResponsePtrOutput {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionByWorkspaceOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionByWorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionByWorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionByWorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionByWorkspaceOutput{})
}
