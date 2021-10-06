


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type App struct {
	pulumi.CustomResourceState

	ApplicationId pulumi.StringOutput      `pulumi:"applicationId"`
	DisplayName   pulumi.StringPtrOutput   `pulumi:"displayName"`
	Location      pulumi.StringOutput      `pulumi:"location"`
	Name          pulumi.StringOutput      `pulumi:"name"`
	Sku           AppSkuInfoResponseOutput `pulumi:"sku"`
	State         pulumi.StringOutput      `pulumi:"state"`
	Subdomain     pulumi.StringPtrOutput   `pulumi:"subdomain"`
	Tags          pulumi.StringMapOutput   `pulumi:"tags"`
	Template      pulumi.StringPtrOutput   `pulumi:"template"`
	Type          pulumi.StringOutput      `pulumi:"type"`
}


func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
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
			Type: pulumi.String("azure-nextgen:iotcentral/v20210601:App"),
		},
		{
			Type: pulumi.String("azure-native:iotcentral:App"),
		},
		{
			Type: pulumi.String("azure-nextgen:iotcentral:App"),
		},
		{
			Type: pulumi.String("azure-native:iotcentral/v20170701privatepreview:App"),
		},
		{
			Type: pulumi.String("azure-nextgen:iotcentral/v20170701privatepreview:App"),
		},
		{
			Type: pulumi.String("azure-native:iotcentral/v20180901:App"),
		},
		{
			Type: pulumi.String("azure-nextgen:iotcentral/v20180901:App"),
		},
	})
	opts = append(opts, aliases)
	var resource App
	err := ctx.RegisterResource("azure-native:iotcentral/v20210601:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("azure-native:iotcentral/v20210601:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appState struct {
}

type AppState struct {
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	DisplayName       *string           `pulumi:"displayName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Sku               AppSkuInfo        `pulumi:"sku"`
	Subdomain         *string           `pulumi:"subdomain"`
	Tags              map[string]string `pulumi:"tags"`
	Template          *string           `pulumi:"template"`
}


type AppArgs struct {
	DisplayName       pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Sku               AppSkuInfoInput
	Subdomain         pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Template          pulumi.StringPtrInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((*App)(nil))
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

type AppOutput struct{ *pulumi.OutputState }

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*App)(nil))
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AppOutput{})
}
