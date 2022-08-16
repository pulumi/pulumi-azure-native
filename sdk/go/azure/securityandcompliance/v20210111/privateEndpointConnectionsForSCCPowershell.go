


package v20210111

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionsForSCCPowershell struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnectionsForSCCPowershell(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionsForSCCPowershellArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsForSCCPowershell, error) {
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
			Type: pulumi.String("azure-native:securityandcompliance:PrivateEndpointConnectionsForSCCPowershell"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsForSCCPowershell"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionsForSCCPowershell
	err := ctx.RegisterResource("azure-native:securityandcompliance/v20210111:PrivateEndpointConnectionsForSCCPowershell", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionsForSCCPowershell(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionsForSCCPowershellState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsForSCCPowershell, error) {
	var resource PrivateEndpointConnectionsForSCCPowershell
	err := ctx.ReadResource("azure-native:securityandcompliance/v20210111:PrivateEndpointConnectionsForSCCPowershell", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionsForSCCPowershellState struct {
}

type PrivateEndpointConnectionsForSCCPowershellState struct {
}

func (PrivateEndpointConnectionsForSCCPowershellState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsForSCCPowershellState)(nil)).Elem()
}

type privateEndpointConnectionsForSCCPowershellArgs struct {
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	ResourceName                      string                            `pulumi:"resourceName"`
}


type PrivateEndpointConnectionsForSCCPowershellArgs struct {
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	ResourceName                      pulumi.StringInput
}

func (PrivateEndpointConnectionsForSCCPowershellArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsForSCCPowershellArgs)(nil)).Elem()
}

type PrivateEndpointConnectionsForSCCPowershellInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionsForSCCPowershellOutput() PrivateEndpointConnectionsForSCCPowershellOutput
	ToPrivateEndpointConnectionsForSCCPowershellOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForSCCPowershellOutput
}

func (*PrivateEndpointConnectionsForSCCPowershell) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsForSCCPowershell)(nil)).Elem()
}

func (i *PrivateEndpointConnectionsForSCCPowershell) ToPrivateEndpointConnectionsForSCCPowershellOutput() PrivateEndpointConnectionsForSCCPowershellOutput {
	return i.ToPrivateEndpointConnectionsForSCCPowershellOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionsForSCCPowershell) ToPrivateEndpointConnectionsForSCCPowershellOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForSCCPowershellOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionsForSCCPowershellOutput)
}

type PrivateEndpointConnectionsForSCCPowershellOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionsForSCCPowershellOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsForSCCPowershell)(nil)).Elem()
}

func (o PrivateEndpointConnectionsForSCCPowershellOutput) ToPrivateEndpointConnectionsForSCCPowershellOutput() PrivateEndpointConnectionsForSCCPowershellOutput {
	return o
}

func (o PrivateEndpointConnectionsForSCCPowershellOutput) ToPrivateEndpointConnectionsForSCCPowershellOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForSCCPowershellOutput {
	return o
}

func (o PrivateEndpointConnectionsForSCCPowershellOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionsForSCCPowershellOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) PrivateEndpointResponsePtrOutput {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionsForSCCPowershellOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionsForSCCPowershellOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionsForSCCPowershellOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionsForSCCPowershellOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionsForSCCPowershellOutput{})
}
