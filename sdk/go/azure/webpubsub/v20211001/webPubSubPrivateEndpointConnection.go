


package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebPubSubPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	GroupIds                          pulumi.StringArrayOutput                           `pulumi:"groupIds"`
	Name                              pulumi.StringOutput                                `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                                `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                           `pulumi:"systemData"`
	Type                              pulumi.StringOutput                                `pulumi:"type"`
}


func NewWebPubSubPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *WebPubSubPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*WebPubSubPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:webpubsub:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210401preview:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210601preview:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210901preview:WebPubSubPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WebPubSubPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:webpubsub/v20211001:WebPubSubPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebPubSubPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*WebPubSubPrivateEndpointConnection, error) {
	var resource WebPubSubPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:webpubsub/v20211001:WebPubSubPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webPubSubPrivateEndpointConnectionState struct {
}

type WebPubSubPrivateEndpointConnectionState struct {
}

func (WebPubSubPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubPrivateEndpointConnectionState)(nil)).Elem()
}

type webPubSubPrivateEndpointConnectionArgs struct {
	PrivateEndpoint                   *PrivateEndpoint                   `pulumi:"privateEndpoint"`
	PrivateEndpointConnectionName     *string                            `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                             `pulumi:"resourceGroupName"`
	ResourceName                      string                             `pulumi:"resourceName"`
}


type WebPubSubPrivateEndpointConnectionArgs struct {
	PrivateEndpoint                   PrivateEndpointPtrInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput
	ResourceGroupName                 pulumi.StringInput
	ResourceName                      pulumi.StringInput
}

func (WebPubSubPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubPrivateEndpointConnectionArgs)(nil)).Elem()
}

type WebPubSubPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToWebPubSubPrivateEndpointConnectionOutput() WebPubSubPrivateEndpointConnectionOutput
	ToWebPubSubPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebPubSubPrivateEndpointConnectionOutput
}

func (*WebPubSubPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubPrivateEndpointConnection)(nil)).Elem()
}

func (i *WebPubSubPrivateEndpointConnection) ToWebPubSubPrivateEndpointConnectionOutput() WebPubSubPrivateEndpointConnectionOutput {
	return i.ToWebPubSubPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *WebPubSubPrivateEndpointConnection) ToWebPubSubPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebPubSubPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubPrivateEndpointConnectionOutput)
}

type WebPubSubPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (WebPubSubPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubPrivateEndpointConnection)(nil)).Elem()
}

func (o WebPubSubPrivateEndpointConnectionOutput) ToWebPubSubPrivateEndpointConnectionOutput() WebPubSubPrivateEndpointConnectionOutput {
	return o
}

func (o WebPubSubPrivateEndpointConnectionOutput) ToWebPubSubPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebPubSubPrivateEndpointConnectionOutput {
	return o
}

func (o WebPubSubPrivateEndpointConnectionOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o WebPubSubPrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebPubSubPrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o WebPubSubPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) PrivateLinkServiceConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o WebPubSubPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WebPubSubPrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WebPubSubPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebPubSubPrivateEndpointConnectionOutput{})
}
