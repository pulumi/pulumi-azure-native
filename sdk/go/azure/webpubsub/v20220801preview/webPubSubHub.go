


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebPubSubHub struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties WebPubSubHubPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput             `pulumi:"systemData"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewWebPubSubHub(ctx *pulumi.Context,
	name string, args *WebPubSubHubArgs, opts ...pulumi.ResourceOption) (*WebPubSubHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	args.Properties = args.Properties.ToWebPubSubHubPropertiesOutput().ApplyT(func(v WebPubSubHubProperties) WebPubSubHubProperties { return *v.Defaults() }).(WebPubSubHubPropertiesOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:webpubsub:WebPubSubHub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20211001:WebPubSubHub"),
		},
	})
	opts = append(opts, aliases)
	var resource WebPubSubHub
	err := ctx.RegisterResource("azure-native:webpubsub/v20220801preview:WebPubSubHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebPubSubHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubHubState, opts ...pulumi.ResourceOption) (*WebPubSubHub, error) {
	var resource WebPubSubHub
	err := ctx.ReadResource("azure-native:webpubsub/v20220801preview:WebPubSubHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webPubSubHubState struct {
}

type WebPubSubHubState struct {
}

func (WebPubSubHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubHubState)(nil)).Elem()
}

type webPubSubHubArgs struct {
	HubName           *string                `pulumi:"hubName"`
	Properties        WebPubSubHubProperties `pulumi:"properties"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	ResourceName      string                 `pulumi:"resourceName"`
}


type WebPubSubHubArgs struct {
	HubName           pulumi.StringPtrInput
	Properties        WebPubSubHubPropertiesInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
}

func (WebPubSubHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubHubArgs)(nil)).Elem()
}

type WebPubSubHubInput interface {
	pulumi.Input

	ToWebPubSubHubOutput() WebPubSubHubOutput
	ToWebPubSubHubOutputWithContext(ctx context.Context) WebPubSubHubOutput
}

func (*WebPubSubHub) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubHub)(nil)).Elem()
}

func (i *WebPubSubHub) ToWebPubSubHubOutput() WebPubSubHubOutput {
	return i.ToWebPubSubHubOutputWithContext(context.Background())
}

func (i *WebPubSubHub) ToWebPubSubHubOutputWithContext(ctx context.Context) WebPubSubHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubHubOutput)
}

type WebPubSubHubOutput struct{ *pulumi.OutputState }

func (WebPubSubHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubHub)(nil)).Elem()
}

func (o WebPubSubHubOutput) ToWebPubSubHubOutput() WebPubSubHubOutput {
	return o
}

func (o WebPubSubHubOutput) ToWebPubSubHubOutputWithContext(ctx context.Context) WebPubSubHubOutput {
	return o
}

func (o WebPubSubHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebPubSubHubOutput) Properties() WebPubSubHubPropertiesResponseOutput {
	return o.ApplyT(func(v *WebPubSubHub) WebPubSubHubPropertiesResponseOutput { return v.Properties }).(WebPubSubHubPropertiesResponseOutput)
}

func (o WebPubSubHubOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebPubSubHub) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WebPubSubHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebPubSubHubOutput{})
}
