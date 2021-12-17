


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppApplicationSettings struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput `pulumi:"kind"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppApplicationSettings(ctx *pulumi.Context,
	name string, args *WebAppApplicationSettingsArgs, opts ...pulumi.ResourceOption) (*WebAppApplicationSettings, error) {
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
			Type: pulumi.String("azure-native:web:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppApplicationSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppApplicationSettings"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppApplicationSettings
	err := ctx.RegisterResource("azure-native:web/v20200601:WebAppApplicationSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppApplicationSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppApplicationSettingsState, opts ...pulumi.ResourceOption) (*WebAppApplicationSettings, error) {
	var resource WebAppApplicationSettings
	err := ctx.ReadResource("azure-native:web/v20200601:WebAppApplicationSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppApplicationSettingsState struct {
}

type WebAppApplicationSettingsState struct {
}

func (WebAppApplicationSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppApplicationSettingsState)(nil)).Elem()
}

type webAppApplicationSettingsArgs struct {
	Kind              *string           `pulumi:"kind"`
	Name              string            `pulumi:"name"`
	Properties        map[string]string `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
}


type WebAppApplicationSettingsArgs struct {
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
}

func (WebAppApplicationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppApplicationSettingsArgs)(nil)).Elem()
}

type WebAppApplicationSettingsInput interface {
	pulumi.Input

	ToWebAppApplicationSettingsOutput() WebAppApplicationSettingsOutput
	ToWebAppApplicationSettingsOutputWithContext(ctx context.Context) WebAppApplicationSettingsOutput
}

func (*WebAppApplicationSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppApplicationSettings)(nil)).Elem()
}

func (i *WebAppApplicationSettings) ToWebAppApplicationSettingsOutput() WebAppApplicationSettingsOutput {
	return i.ToWebAppApplicationSettingsOutputWithContext(context.Background())
}

func (i *WebAppApplicationSettings) ToWebAppApplicationSettingsOutputWithContext(ctx context.Context) WebAppApplicationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppApplicationSettingsOutput)
}

type WebAppApplicationSettingsOutput struct{ *pulumi.OutputState }

func (WebAppApplicationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppApplicationSettings)(nil)).Elem()
}

func (o WebAppApplicationSettingsOutput) ToWebAppApplicationSettingsOutput() WebAppApplicationSettingsOutput {
	return o
}

func (o WebAppApplicationSettingsOutput) ToWebAppApplicationSettingsOutputWithContext(ctx context.Context) WebAppApplicationSettingsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppApplicationSettingsOutput{})
}
