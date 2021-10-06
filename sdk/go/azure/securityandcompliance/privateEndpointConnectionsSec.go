


package securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionsSec struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnectionsSec(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionsSecArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsSec, error) {
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
			Type: pulumi.String("azure-nextgen:securityandcompliance:PrivateEndpointConnectionsSec"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210111:PrivateEndpointConnectionsSec"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance/v20210111:PrivateEndpointConnectionsSec"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsSec"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance/v20210308:PrivateEndpointConnectionsSec"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionsSec
	err := ctx.RegisterResource("azure-native:securityandcompliance:PrivateEndpointConnectionsSec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionsSec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionsSecState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsSec, error) {
	var resource PrivateEndpointConnectionsSec
	err := ctx.ReadResource("azure-native:securityandcompliance:PrivateEndpointConnectionsSec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionsSecState struct {
}

type PrivateEndpointConnectionsSecState struct {
}

func (PrivateEndpointConnectionsSecState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsSecState)(nil)).Elem()
}

type privateEndpointConnectionsSecArgs struct {
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	ResourceName                      string                            `pulumi:"resourceName"`
}


type PrivateEndpointConnectionsSecArgs struct {
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	ResourceName                      pulumi.StringInput
}

func (PrivateEndpointConnectionsSecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsSecArgs)(nil)).Elem()
}

type PrivateEndpointConnectionsSecInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionsSecOutput() PrivateEndpointConnectionsSecOutput
	ToPrivateEndpointConnectionsSecOutputWithContext(ctx context.Context) PrivateEndpointConnectionsSecOutput
}

func (*PrivateEndpointConnectionsSec) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionsSec)(nil))
}

func (i *PrivateEndpointConnectionsSec) ToPrivateEndpointConnectionsSecOutput() PrivateEndpointConnectionsSecOutput {
	return i.ToPrivateEndpointConnectionsSecOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionsSec) ToPrivateEndpointConnectionsSecOutputWithContext(ctx context.Context) PrivateEndpointConnectionsSecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionsSecOutput)
}

type PrivateEndpointConnectionsSecOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionsSecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionsSec)(nil))
}

func (o PrivateEndpointConnectionsSecOutput) ToPrivateEndpointConnectionsSecOutput() PrivateEndpointConnectionsSecOutput {
	return o
}

func (o PrivateEndpointConnectionsSecOutput) ToPrivateEndpointConnectionsSecOutputWithContext(ctx context.Context) PrivateEndpointConnectionsSecOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionsSecOutput{})
}
