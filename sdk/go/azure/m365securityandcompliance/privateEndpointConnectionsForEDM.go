


package m365securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionsForEDM struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnectionsForEDM(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionsForEDMArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsForEDM, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:m365securityandcompliance:PrivateEndpointConnectionsForEDM"),
		},
		{
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsForEDM"),
		},
		{
			Type: pulumi.String("azure-nextgen:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsForEDM"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionsForEDM
	err := ctx.RegisterResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsForEDM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionsForEDM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionsForEDMState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsForEDM, error) {
	var resource PrivateEndpointConnectionsForEDM
	err := ctx.ReadResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsForEDM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionsForEDMState struct {
}

type PrivateEndpointConnectionsForEDMState struct {
}

func (PrivateEndpointConnectionsForEDMState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsForEDMState)(nil)).Elem()
}

type privateEndpointConnectionsForEDMArgs struct {
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	ResourceName                      string                            `pulumi:"resourceName"`
}


type PrivateEndpointConnectionsForEDMArgs struct {
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	ResourceName                      pulumi.StringInput
}

func (PrivateEndpointConnectionsForEDMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsForEDMArgs)(nil)).Elem()
}

type PrivateEndpointConnectionsForEDMInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionsForEDMOutput() PrivateEndpointConnectionsForEDMOutput
	ToPrivateEndpointConnectionsForEDMOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForEDMOutput
}

func (*PrivateEndpointConnectionsForEDM) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionsForEDM)(nil))
}

func (i *PrivateEndpointConnectionsForEDM) ToPrivateEndpointConnectionsForEDMOutput() PrivateEndpointConnectionsForEDMOutput {
	return i.ToPrivateEndpointConnectionsForEDMOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionsForEDM) ToPrivateEndpointConnectionsForEDMOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForEDMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionsForEDMOutput)
}

type PrivateEndpointConnectionsForEDMOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionsForEDMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionsForEDM)(nil))
}

func (o PrivateEndpointConnectionsForEDMOutput) ToPrivateEndpointConnectionsForEDMOutput() PrivateEndpointConnectionsForEDMOutput {
	return o
}

func (o PrivateEndpointConnectionsForEDMOutput) ToPrivateEndpointConnectionsForEDMOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForEDMOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionsForEDMInput)(nil)).Elem(), &PrivateEndpointConnectionsForEDM{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionsForEDMOutput{})
}
