


package m365securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionsComp struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnectionsComp(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionsCompArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsComp, error) {
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
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsComp"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionsComp
	err := ctx.RegisterResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsComp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionsComp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionsCompState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsComp, error) {
	var resource PrivateEndpointConnectionsComp
	err := ctx.ReadResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsComp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionsCompState struct {
}

type PrivateEndpointConnectionsCompState struct {
}

func (PrivateEndpointConnectionsCompState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsCompState)(nil)).Elem()
}

type privateEndpointConnectionsCompArgs struct {
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	ResourceName                      string                            `pulumi:"resourceName"`
}


type PrivateEndpointConnectionsCompArgs struct {
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	ResourceName                      pulumi.StringInput
}

func (PrivateEndpointConnectionsCompArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsCompArgs)(nil)).Elem()
}

type PrivateEndpointConnectionsCompInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionsCompOutput() PrivateEndpointConnectionsCompOutput
	ToPrivateEndpointConnectionsCompOutputWithContext(ctx context.Context) PrivateEndpointConnectionsCompOutput
}

func (*PrivateEndpointConnectionsComp) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsComp)(nil)).Elem()
}

func (i *PrivateEndpointConnectionsComp) ToPrivateEndpointConnectionsCompOutput() PrivateEndpointConnectionsCompOutput {
	return i.ToPrivateEndpointConnectionsCompOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionsComp) ToPrivateEndpointConnectionsCompOutputWithContext(ctx context.Context) PrivateEndpointConnectionsCompOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionsCompOutput)
}

type PrivateEndpointConnectionsCompOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionsCompOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsComp)(nil)).Elem()
}

func (o PrivateEndpointConnectionsCompOutput) ToPrivateEndpointConnectionsCompOutput() PrivateEndpointConnectionsCompOutput {
	return o
}

func (o PrivateEndpointConnectionsCompOutput) ToPrivateEndpointConnectionsCompOutputWithContext(ctx context.Context) PrivateEndpointConnectionsCompOutput {
	return o
}

func (o PrivateEndpointConnectionsCompOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsComp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionsCompOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsComp) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionsCompOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsComp) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionsCompOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsComp) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionsCompOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsComp) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionsCompOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsComp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionsCompOutput{})
}
