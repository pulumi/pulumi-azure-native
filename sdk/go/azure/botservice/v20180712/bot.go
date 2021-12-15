


package v20180712

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Bot struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput      `pulumi:"etag"`
	Kind       pulumi.StringPtrOutput      `pulumi:"kind"`
	Location   pulumi.StringPtrOutput      `pulumi:"location"`
	Name       pulumi.StringOutput         `pulumi:"name"`
	Properties BotPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput        `pulumi:"sku"`
	Tags       pulumi.StringMapOutput      `pulumi:"tags"`
	Type       pulumi.StringOutput         `pulumi:"type"`
}


func NewBot(ctx *pulumi.Context,
	name string, args *BotArgs, opts ...pulumi.ResourceOption) (*Bot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:botservice:Bot"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20171201:Bot"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20200602:Bot"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20210301:Bot"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20210501preview:Bot"),
		},
	})
	opts = append(opts, aliases)
	var resource Bot
	err := ctx.RegisterResource("azure-native:botservice/v20180712:Bot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotState, opts ...pulumi.ResourceOption) (*Bot, error) {
	var resource Bot
	err := ctx.ReadResource("azure-native:botservice/v20180712:Bot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type botState struct {
}

type BotState struct {
}

func (BotState) ElementType() reflect.Type {
	return reflect.TypeOf((*botState)(nil)).Elem()
}

type botArgs struct {
	Etag              *string           `pulumi:"etag"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Properties        *BotProperties    `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type BotArgs struct {
	Etag              pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        BotPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
}

func (BotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*botArgs)(nil)).Elem()
}

type BotInput interface {
	pulumi.Input

	ToBotOutput() BotOutput
	ToBotOutputWithContext(ctx context.Context) BotOutput
}

func (*Bot) ElementType() reflect.Type {
	return reflect.TypeOf((**Bot)(nil)).Elem()
}

func (i *Bot) ToBotOutput() BotOutput {
	return i.ToBotOutputWithContext(context.Background())
}

func (i *Bot) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotOutput)
}

type BotOutput struct{ *pulumi.OutputState }

func (BotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bot)(nil)).Elem()
}

func (o BotOutput) ToBotOutput() BotOutput {
	return o
}

func (o BotOutput) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BotOutput{})
}
