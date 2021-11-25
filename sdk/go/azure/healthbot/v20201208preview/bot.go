


package v20201208preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Bot struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput               `pulumi:"location"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties HealthBotPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput              `pulumi:"sku"`
	SystemData SystemDataResponseOutput          `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput            `pulumi:"tags"`
	Type       pulumi.StringOutput               `pulumi:"type"`
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
			Type: pulumi.String("azure-native:healthbot:Bot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20201020:Bot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20201020preview:Bot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20201208:Bot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20210610:Bot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20210824:Bot"),
		},
	})
	opts = append(opts, aliases)
	var resource Bot
	err := ctx.RegisterResource("azure-native:healthbot/v20201208preview:Bot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotState, opts ...pulumi.ResourceOption) (*Bot, error) {
	var resource Bot
	err := ctx.ReadResource("azure-native:healthbot/v20201208preview:Bot", name, id, state, &resource, opts...)
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
	BotName           *string           `pulumi:"botName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type BotArgs struct {
	BotName           pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
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
	return reflect.TypeOf((*Bot)(nil))
}

func (i *Bot) ToBotOutput() BotOutput {
	return i.ToBotOutputWithContext(context.Background())
}

func (i *Bot) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotOutput)
}

type BotOutput struct{ *pulumi.OutputState }

func (BotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Bot)(nil))
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
