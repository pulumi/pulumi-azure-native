


package healthbot

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
	Sku        SkuResponseOutput                 `pulumi:"sku"`
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
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
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
			Type: pulumi.String("azure-native:healthbot/v20201208preview:Bot"),
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
	err := ctx.RegisterResource("azure-native:healthbot:Bot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotState, opts ...pulumi.ResourceOption) (*Bot, error) {
	var resource Bot
	err := ctx.ReadResource("azure-native:healthbot:Bot", name, id, state, &resource, opts...)
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
	Sku               Sku               `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type BotArgs struct {
	BotName           pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuInput
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

func (o BotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BotOutput) Properties() HealthBotPropertiesResponseOutput {
	return o.ApplyT(func(v *Bot) HealthBotPropertiesResponseOutput { return v.Properties }).(HealthBotPropertiesResponseOutput)
}

func (o BotOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Bot) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o BotOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Bot) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BotOutput{})
}
