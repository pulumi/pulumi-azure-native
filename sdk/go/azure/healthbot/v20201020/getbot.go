


package v20201020

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Getbot struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput               `pulumi:"location"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties HealthBotPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput              `pulumi:"sku"`
	SystemData SystemDataResponseOutput          `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput            `pulumi:"tags"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}


func NewGetbot(ctx *pulumi.Context,
	name string, args *GetbotArgs, opts ...pulumi.ResourceOption) (*Getbot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthbot:getbot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20201020preview:getbot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20201208:getbot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20201208preview:getbot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20210610:getbot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20210824:getbot"),
		},
		{
			Type: pulumi.String("azure-native:healthbot/v20220808:getbot"),
		},
	})
	opts = append(opts, aliases)
	var resource Getbot
	err := ctx.RegisterResource("azure-native:healthbot/v20201020:getbot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGetbot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GetbotState, opts ...pulumi.ResourceOption) (*Getbot, error) {
	var resource Getbot
	err := ctx.ReadResource("azure-native:healthbot/v20201020:getbot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type getbotState struct {
}

type GetbotState struct {
}

func (GetbotState) ElementType() reflect.Type {
	return reflect.TypeOf((*getbotState)(nil)).Elem()
}

type getbotArgs struct {
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type GetbotArgs struct {
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
}

func (GetbotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*getbotArgs)(nil)).Elem()
}

type GetbotInput interface {
	pulumi.Input

	ToGetbotOutput() GetbotOutput
	ToGetbotOutputWithContext(ctx context.Context) GetbotOutput
}

func (*Getbot) ElementType() reflect.Type {
	return reflect.TypeOf((**Getbot)(nil)).Elem()
}

func (i *Getbot) ToGetbotOutput() GetbotOutput {
	return i.ToGetbotOutputWithContext(context.Background())
}

func (i *Getbot) ToGetbotOutputWithContext(ctx context.Context) GetbotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetbotOutput)
}

type GetbotOutput struct{ *pulumi.OutputState }

func (GetbotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Getbot)(nil)).Elem()
}

func (o GetbotOutput) ToGetbotOutput() GetbotOutput {
	return o
}

func (o GetbotOutput) ToGetbotOutputWithContext(ctx context.Context) GetbotOutput {
	return o
}

func (o GetbotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Getbot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GetbotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Getbot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GetbotOutput) Properties() HealthBotPropertiesResponseOutput {
	return o.ApplyT(func(v *Getbot) HealthBotPropertiesResponseOutput { return v.Properties }).(HealthBotPropertiesResponseOutput)
}

func (o GetbotOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Getbot) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o GetbotOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Getbot) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetbotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Getbot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetbotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Getbot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetbotOutput{})
}
