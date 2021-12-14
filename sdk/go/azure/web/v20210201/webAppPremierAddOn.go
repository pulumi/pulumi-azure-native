


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppPremierAddOn struct {
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


func NewWebAppPremierAddOn(ctx *pulumi.Context,
	name string, args *WebAppPremierAddOnArgs, opts ...pulumi.ResourceOption) (*WebAppPremierAddOn, error) {
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
			Type: pulumi.String("azure-native:web:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppPremierAddOn"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppPremierAddOn"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppPremierAddOn
	err := ctx.RegisterResource("azure-native:web/v20210201:WebAppPremierAddOn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppPremierAddOn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPremierAddOnState, opts ...pulumi.ResourceOption) (*WebAppPremierAddOn, error) {
	var resource WebAppPremierAddOn
	err := ctx.ReadResource("azure-native:web/v20210201:WebAppPremierAddOn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppPremierAddOnState struct {
}

type WebAppPremierAddOnState struct {
}

func (WebAppPremierAddOnState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPremierAddOnState)(nil)).Elem()
}

type webAppPremierAddOnArgs struct {
	Kind                 *string           `pulumi:"kind"`
	Location             *string           `pulumi:"location"`
	MarketplaceOffer     *string           `pulumi:"marketplaceOffer"`
	MarketplacePublisher *string           `pulumi:"marketplacePublisher"`
	Name                 string            `pulumi:"name"`
	PremierAddOnName     *string           `pulumi:"premierAddOnName"`
	Product              *string           `pulumi:"product"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	Sku                  *string           `pulumi:"sku"`
	Tags                 map[string]string `pulumi:"tags"`
	Vendor               *string           `pulumi:"vendor"`
}


type WebAppPremierAddOnArgs struct {
	Kind                 pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	MarketplaceOffer     pulumi.StringPtrInput
	MarketplacePublisher pulumi.StringPtrInput
	Name                 pulumi.StringInput
	PremierAddOnName     pulumi.StringPtrInput
	Product              pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Sku                  pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
	Vendor               pulumi.StringPtrInput
}

func (WebAppPremierAddOnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPremierAddOnArgs)(nil)).Elem()
}

type WebAppPremierAddOnInput interface {
	pulumi.Input

	ToWebAppPremierAddOnOutput() WebAppPremierAddOnOutput
	ToWebAppPremierAddOnOutputWithContext(ctx context.Context) WebAppPremierAddOnOutput
}

func (*WebAppPremierAddOn) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPremierAddOn)(nil)).Elem()
}

func (i *WebAppPremierAddOn) ToWebAppPremierAddOnOutput() WebAppPremierAddOnOutput {
	return i.ToWebAppPremierAddOnOutputWithContext(context.Background())
}

func (i *WebAppPremierAddOn) ToWebAppPremierAddOnOutputWithContext(ctx context.Context) WebAppPremierAddOnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPremierAddOnOutput)
}

type WebAppPremierAddOnOutput struct{ *pulumi.OutputState }

func (WebAppPremierAddOnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPremierAddOn)(nil)).Elem()
}

func (o WebAppPremierAddOnOutput) ToWebAppPremierAddOnOutput() WebAppPremierAddOnOutput {
	return o
}

func (o WebAppPremierAddOnOutput) ToWebAppPremierAddOnOutputWithContext(ctx context.Context) WebAppPremierAddOnOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppPremierAddOnOutput{})
}
