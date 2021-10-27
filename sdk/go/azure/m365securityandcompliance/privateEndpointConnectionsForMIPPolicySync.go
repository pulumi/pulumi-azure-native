


package m365securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionsForMIPPolicySync struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnectionsForMIPPolicySync(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionsForMIPPolicySyncArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsForMIPPolicySync, error) {
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
			Type: pulumi.String("azure-nextgen:m365securityandcompliance:PrivateEndpointConnectionsForMIPPolicySync"),
		},
		{
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsForMIPPolicySync"),
		},
		{
			Type: pulumi.String("azure-nextgen:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsForMIPPolicySync"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionsForMIPPolicySync
	err := ctx.RegisterResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsForMIPPolicySync", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionsForMIPPolicySync(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionsForMIPPolicySyncState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsForMIPPolicySync, error) {
	var resource PrivateEndpointConnectionsForMIPPolicySync
	err := ctx.ReadResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsForMIPPolicySync", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionsForMIPPolicySyncState struct {
}

type PrivateEndpointConnectionsForMIPPolicySyncState struct {
}

func (PrivateEndpointConnectionsForMIPPolicySyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsForMIPPolicySyncState)(nil)).Elem()
}

type privateEndpointConnectionsForMIPPolicySyncArgs struct {
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	ResourceName                      string                            `pulumi:"resourceName"`
}


type PrivateEndpointConnectionsForMIPPolicySyncArgs struct {
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	ResourceName                      pulumi.StringInput
}

func (PrivateEndpointConnectionsForMIPPolicySyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsForMIPPolicySyncArgs)(nil)).Elem()
}

type PrivateEndpointConnectionsForMIPPolicySyncInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionsForMIPPolicySyncOutput() PrivateEndpointConnectionsForMIPPolicySyncOutput
	ToPrivateEndpointConnectionsForMIPPolicySyncOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForMIPPolicySyncOutput
}

func (*PrivateEndpointConnectionsForMIPPolicySync) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionsForMIPPolicySync)(nil))
}

func (i *PrivateEndpointConnectionsForMIPPolicySync) ToPrivateEndpointConnectionsForMIPPolicySyncOutput() PrivateEndpointConnectionsForMIPPolicySyncOutput {
	return i.ToPrivateEndpointConnectionsForMIPPolicySyncOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionsForMIPPolicySync) ToPrivateEndpointConnectionsForMIPPolicySyncOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForMIPPolicySyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionsForMIPPolicySyncOutput)
}

type PrivateEndpointConnectionsForMIPPolicySyncOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionsForMIPPolicySyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionsForMIPPolicySync)(nil))
}

func (o PrivateEndpointConnectionsForMIPPolicySyncOutput) ToPrivateEndpointConnectionsForMIPPolicySyncOutput() PrivateEndpointConnectionsForMIPPolicySyncOutput {
	return o
}

func (o PrivateEndpointConnectionsForMIPPolicySyncOutput) ToPrivateEndpointConnectionsForMIPPolicySyncOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForMIPPolicySyncOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionsForMIPPolicySyncInput)(nil)).Elem(), &PrivateEndpointConnectionsForMIPPolicySync{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionsForMIPPolicySyncOutput{})
}
