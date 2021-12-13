


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebPubSub struct {
	pulumi.CustomResourceState

	EventHandler               EventHandlerSettingsResponsePtrOutput        `pulumi:"eventHandler"`
	ExternalIP                 pulumi.StringOutput                          `pulumi:"externalIP"`
	Features                   WebPubSubFeatureResponseArrayOutput          `pulumi:"features"`
	HostName                   pulumi.StringOutput                          `pulumi:"hostName"`
	Identity                   ManagedIdentityResponsePtrOutput             `pulumi:"identity"`
	Location                   pulumi.StringPtrOutput                       `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	NetworkACLs                WebPubSubNetworkACLsResponsePtrOutput        `pulumi:"networkACLs"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccess"`
	PublicPort                 pulumi.IntOutput                             `pulumi:"publicPort"`
	ServerPort                 pulumi.IntOutput                             `pulumi:"serverPort"`
	SharedPrivateLinkResources SharedPrivateLinkResourceResponseArrayOutput `pulumi:"sharedPrivateLinkResources"`
	Sku                        ResourceSkuResponsePtrOutput                 `pulumi:"sku"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Tls                        WebPubSubTlsSettingsResponsePtrOutput        `pulumi:"tls"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
	Version                    pulumi.StringOutput                          `pulumi:"version"`
}


func NewWebPubSub(ctx *pulumi.Context,
	name string, args *WebPubSubArgs, opts ...pulumi.ResourceOption) (*WebPubSub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.NetworkACLs != nil {
		args.NetworkACLs = args.NetworkACLs.ToWebPubSubNetworkACLsPtrOutput().ApplyT(func(v *WebPubSubNetworkACLs) *WebPubSubNetworkACLs { return v.Defaults() }).(WebPubSubNetworkACLsPtrOutput)
	}
	if isZero(args.PublicNetworkAccess) {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	if args.Tls != nil {
		args.Tls = args.Tls.ToWebPubSubTlsSettingsPtrOutput().ApplyT(func(v *WebPubSubTlsSettings) *WebPubSubTlsSettings { return v.Defaults() }).(WebPubSubTlsSettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:webpubsub:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210601preview:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210901preview:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20211001:WebPubSub"),
		},
	})
	opts = append(opts, aliases)
	var resource WebPubSub
	err := ctx.RegisterResource("azure-native:webpubsub/v20210401preview:WebPubSub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebPubSub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubState, opts ...pulumi.ResourceOption) (*WebPubSub, error) {
	var resource WebPubSub
	err := ctx.ReadResource("azure-native:webpubsub/v20210401preview:WebPubSub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webPubSubState struct {
}

type WebPubSubState struct {
}

func (WebPubSubState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubState)(nil)).Elem()
}

type webPubSubArgs struct {
	EventHandler        *EventHandlerSettings `pulumi:"eventHandler"`
	Features            []WebPubSubFeature    `pulumi:"features"`
	Identity            *ManagedIdentity      `pulumi:"identity"`
	Location            *string               `pulumi:"location"`
	NetworkACLs         *WebPubSubNetworkACLs `pulumi:"networkACLs"`
	PublicNetworkAccess *string               `pulumi:"publicNetworkAccess"`
	ResourceGroupName   string                `pulumi:"resourceGroupName"`
	ResourceName        *string               `pulumi:"resourceName"`
	Sku                 *ResourceSku          `pulumi:"sku"`
	Tags                map[string]string     `pulumi:"tags"`
	Tls                 *WebPubSubTlsSettings `pulumi:"tls"`
}


type WebPubSubArgs struct {
	EventHandler        EventHandlerSettingsPtrInput
	Features            WebPubSubFeatureArrayInput
	Identity            ManagedIdentityPtrInput
	Location            pulumi.StringPtrInput
	NetworkACLs         WebPubSubNetworkACLsPtrInput
	PublicNetworkAccess pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringPtrInput
	Sku                 ResourceSkuPtrInput
	Tags                pulumi.StringMapInput
	Tls                 WebPubSubTlsSettingsPtrInput
}

func (WebPubSubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubArgs)(nil)).Elem()
}

type WebPubSubInput interface {
	pulumi.Input

	ToWebPubSubOutput() WebPubSubOutput
	ToWebPubSubOutputWithContext(ctx context.Context) WebPubSubOutput
}

func (*WebPubSub) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSub)(nil)).Elem()
}

func (i *WebPubSub) ToWebPubSubOutput() WebPubSubOutput {
	return i.ToWebPubSubOutputWithContext(context.Background())
}

func (i *WebPubSub) ToWebPubSubOutputWithContext(ctx context.Context) WebPubSubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubOutput)
}

type WebPubSubOutput struct{ *pulumi.OutputState }

func (WebPubSubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSub)(nil)).Elem()
}

func (o WebPubSubOutput) ToWebPubSubOutput() WebPubSubOutput {
	return o
}

func (o WebPubSubOutput) ToWebPubSubOutputWithContext(ctx context.Context) WebPubSubOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebPubSubOutput{})
}
