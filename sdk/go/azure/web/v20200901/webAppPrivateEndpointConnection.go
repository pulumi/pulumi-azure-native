


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Kind                              pulumi.StringPtrOutput                      `pulumi:"kind"`
	Name                              pulumi.StringOutput                         `pulumi:"name"`
	PrivateEndpoint                   ArmIdWrapperResponsePtrOutput               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                         `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                    `pulumi:"systemData"`
	Type                              pulumi.StringOutput                         `pulumi:"type"`
}


func NewWebAppPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *WebAppPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*WebAppPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:WebAppPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:web/v20200901:WebAppPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*WebAppPrivateEndpointConnection, error) {
	var resource WebAppPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:web/v20200901:WebAppPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppPrivateEndpointConnectionState struct {
}

type WebAppPrivateEndpointConnectionState struct {
}

func (WebAppPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPrivateEndpointConnectionState)(nil)).Elem()
}

type webAppPrivateEndpointConnectionArgs struct {
	Kind                              *string                     `pulumi:"kind"`
	Name                              string                      `pulumi:"name"`
	PrivateEndpointConnectionName     *string                     `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                      `pulumi:"resourceGroupName"`
}


type WebAppPrivateEndpointConnectionArgs struct {
	Kind                              pulumi.StringPtrInput
	Name                              pulumi.StringInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkConnectionStatePtrInput
	ResourceGroupName                 pulumi.StringInput
}

func (WebAppPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPrivateEndpointConnectionArgs)(nil)).Elem()
}

type WebAppPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToWebAppPrivateEndpointConnectionOutput() WebAppPrivateEndpointConnectionOutput
	ToWebAppPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebAppPrivateEndpointConnectionOutput
}

func (*WebAppPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppPrivateEndpointConnection)(nil))
}

func (i *WebAppPrivateEndpointConnection) ToWebAppPrivateEndpointConnectionOutput() WebAppPrivateEndpointConnectionOutput {
	return i.ToWebAppPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *WebAppPrivateEndpointConnection) ToWebAppPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebAppPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPrivateEndpointConnectionOutput)
}

type WebAppPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (WebAppPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppPrivateEndpointConnection)(nil))
}

func (o WebAppPrivateEndpointConnectionOutput) ToWebAppPrivateEndpointConnectionOutput() WebAppPrivateEndpointConnectionOutput {
	return o
}

func (o WebAppPrivateEndpointConnectionOutput) ToWebAppPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebAppPrivateEndpointConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppPrivateEndpointConnectionOutput{})
}
