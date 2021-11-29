


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppPremierAddOnSlot struct {
	pulumi.CustomResourceState

	Kind                 pulumi.StringPtrOutput `pulumi:"kind"`
	Location             pulumi.StringOutput    `pulumi:"location"`
	MarketplaceOffer     pulumi.StringPtrOutput `pulumi:"marketplaceOffer"`
	MarketplacePublisher pulumi.StringPtrOutput `pulumi:"marketplacePublisher"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	Product              pulumi.StringPtrOutput `pulumi:"product"`
	Sku                  pulumi.StringPtrOutput `pulumi:"sku"`
	Tags                 pulumi.StringMapOutput `pulumi:"tags"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
	Vendor               pulumi.StringPtrOutput `pulumi:"vendor"`
}


func NewWebAppPremierAddOnSlot(ctx *pulumi.Context,
	name string, args *WebAppPremierAddOnSlotArgs, opts ...pulumi.ResourceOption) (*WebAppPremierAddOnSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppPremierAddOnSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppPremierAddOnSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppPremierAddOnSlot
	err := ctx.RegisterResource("azure-native:web/v20210101:WebAppPremierAddOnSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppPremierAddOnSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPremierAddOnSlotState, opts ...pulumi.ResourceOption) (*WebAppPremierAddOnSlot, error) {
	var resource WebAppPremierAddOnSlot
	err := ctx.ReadResource("azure-native:web/v20210101:WebAppPremierAddOnSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppPremierAddOnSlotState struct {
}

type WebAppPremierAddOnSlotState struct {
}

func (WebAppPremierAddOnSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPremierAddOnSlotState)(nil)).Elem()
}

type webAppPremierAddOnSlotArgs struct {
	Kind                 *string           `pulumi:"kind"`
	Location             *string           `pulumi:"location"`
	MarketplaceOffer     *string           `pulumi:"marketplaceOffer"`
	MarketplacePublisher *string           `pulumi:"marketplacePublisher"`
	Name                 string            `pulumi:"name"`
	PremierAddOnName     *string           `pulumi:"premierAddOnName"`
	Product              *string           `pulumi:"product"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	Sku                  *string           `pulumi:"sku"`
	Slot                 string            `pulumi:"slot"`
	Tags                 map[string]string `pulumi:"tags"`
	Vendor               *string           `pulumi:"vendor"`
}


type WebAppPremierAddOnSlotArgs struct {
	Kind                 pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	MarketplaceOffer     pulumi.StringPtrInput
	MarketplacePublisher pulumi.StringPtrInput
	Name                 pulumi.StringInput
	PremierAddOnName     pulumi.StringPtrInput
	Product              pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Sku                  pulumi.StringPtrInput
	Slot                 pulumi.StringInput
	Tags                 pulumi.StringMapInput
	Vendor               pulumi.StringPtrInput
}

func (WebAppPremierAddOnSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPremierAddOnSlotArgs)(nil)).Elem()
}

type WebAppPremierAddOnSlotInput interface {
	pulumi.Input

	ToWebAppPremierAddOnSlotOutput() WebAppPremierAddOnSlotOutput
	ToWebAppPremierAddOnSlotOutputWithContext(ctx context.Context) WebAppPremierAddOnSlotOutput
}

func (*WebAppPremierAddOnSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppPremierAddOnSlot)(nil))
}

func (i *WebAppPremierAddOnSlot) ToWebAppPremierAddOnSlotOutput() WebAppPremierAddOnSlotOutput {
	return i.ToWebAppPremierAddOnSlotOutputWithContext(context.Background())
}

func (i *WebAppPremierAddOnSlot) ToWebAppPremierAddOnSlotOutputWithContext(ctx context.Context) WebAppPremierAddOnSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPremierAddOnSlotOutput)
}

type WebAppPremierAddOnSlotOutput struct{ *pulumi.OutputState }

func (WebAppPremierAddOnSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppPremierAddOnSlot)(nil))
}

func (o WebAppPremierAddOnSlotOutput) ToWebAppPremierAddOnSlotOutput() WebAppPremierAddOnSlotOutput {
	return o
}

func (o WebAppPremierAddOnSlotOutput) ToWebAppPremierAddOnSlotOutputWithContext(ctx context.Context) WebAppPremierAddOnSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppPremierAddOnSlotOutput{})
}
