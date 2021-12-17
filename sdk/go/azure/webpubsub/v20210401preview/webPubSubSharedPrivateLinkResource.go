


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebPubSubSharedPrivateLinkResource struct {
	pulumi.CustomResourceState

	GroupId               pulumi.StringOutput      `pulumi:"groupId"`
	Name                  pulumi.StringOutput      `pulumi:"name"`
	PrivateLinkResourceId pulumi.StringOutput      `pulumi:"privateLinkResourceId"`
	ProvisioningState     pulumi.StringOutput      `pulumi:"provisioningState"`
	RequestMessage        pulumi.StringPtrOutput   `pulumi:"requestMessage"`
	Status                pulumi.StringOutput      `pulumi:"status"`
	SystemData            SystemDataResponseOutput `pulumi:"systemData"`
	Type                  pulumi.StringOutput      `pulumi:"type"`
}


func NewWebPubSubSharedPrivateLinkResource(ctx *pulumi.Context,
	name string, args *WebPubSubSharedPrivateLinkResourceArgs, opts ...pulumi.ResourceOption) (*WebPubSubSharedPrivateLinkResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.PrivateLinkResourceId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:webpubsub:WebPubSubSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210601preview:WebPubSubSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210901preview:WebPubSubSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20211001:WebPubSubSharedPrivateLinkResource"),
		},
	})
	opts = append(opts, aliases)
	var resource WebPubSubSharedPrivateLinkResource
	err := ctx.RegisterResource("azure-native:webpubsub/v20210401preview:WebPubSubSharedPrivateLinkResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebPubSubSharedPrivateLinkResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubSharedPrivateLinkResourceState, opts ...pulumi.ResourceOption) (*WebPubSubSharedPrivateLinkResource, error) {
	var resource WebPubSubSharedPrivateLinkResource
	err := ctx.ReadResource("azure-native:webpubsub/v20210401preview:WebPubSubSharedPrivateLinkResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webPubSubSharedPrivateLinkResourceState struct {
}

type WebPubSubSharedPrivateLinkResourceState struct {
}

func (WebPubSubSharedPrivateLinkResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubSharedPrivateLinkResourceState)(nil)).Elem()
}

type webPubSubSharedPrivateLinkResourceArgs struct {
	GroupId                       string  `pulumi:"groupId"`
	PrivateLinkResourceId         string  `pulumi:"privateLinkResourceId"`
	RequestMessage                *string `pulumi:"requestMessage"`
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
	ResourceName                  string  `pulumi:"resourceName"`
	SharedPrivateLinkResourceName *string `pulumi:"sharedPrivateLinkResourceName"`
}


type WebPubSubSharedPrivateLinkResourceArgs struct {
	GroupId                       pulumi.StringInput
	PrivateLinkResourceId         pulumi.StringInput
	RequestMessage                pulumi.StringPtrInput
	ResourceGroupName             pulumi.StringInput
	ResourceName                  pulumi.StringInput
	SharedPrivateLinkResourceName pulumi.StringPtrInput
}

func (WebPubSubSharedPrivateLinkResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubSharedPrivateLinkResourceArgs)(nil)).Elem()
}

type WebPubSubSharedPrivateLinkResourceInput interface {
	pulumi.Input

	ToWebPubSubSharedPrivateLinkResourceOutput() WebPubSubSharedPrivateLinkResourceOutput
	ToWebPubSubSharedPrivateLinkResourceOutputWithContext(ctx context.Context) WebPubSubSharedPrivateLinkResourceOutput
}

func (*WebPubSubSharedPrivateLinkResource) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubSharedPrivateLinkResource)(nil)).Elem()
}

func (i *WebPubSubSharedPrivateLinkResource) ToWebPubSubSharedPrivateLinkResourceOutput() WebPubSubSharedPrivateLinkResourceOutput {
	return i.ToWebPubSubSharedPrivateLinkResourceOutputWithContext(context.Background())
}

func (i *WebPubSubSharedPrivateLinkResource) ToWebPubSubSharedPrivateLinkResourceOutputWithContext(ctx context.Context) WebPubSubSharedPrivateLinkResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubSharedPrivateLinkResourceOutput)
}

type WebPubSubSharedPrivateLinkResourceOutput struct{ *pulumi.OutputState }

func (WebPubSubSharedPrivateLinkResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubSharedPrivateLinkResource)(nil)).Elem()
}

func (o WebPubSubSharedPrivateLinkResourceOutput) ToWebPubSubSharedPrivateLinkResourceOutput() WebPubSubSharedPrivateLinkResourceOutput {
	return o
}

func (o WebPubSubSharedPrivateLinkResourceOutput) ToWebPubSubSharedPrivateLinkResourceOutputWithContext(ctx context.Context) WebPubSubSharedPrivateLinkResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebPubSubSharedPrivateLinkResourceOutput{})
}
