


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppSitePushSettingsSlot struct {
	pulumi.CustomResourceState

	DynamicTagsJson   pulumi.StringPtrOutput `pulumi:"dynamicTagsJson"`
	IsPushEnabled     pulumi.BoolOutput      `pulumi:"isPushEnabled"`
	Kind              pulumi.StringPtrOutput `pulumi:"kind"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	TagWhitelistJson  pulumi.StringPtrOutput `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth pulumi.StringPtrOutput `pulumi:"tagsRequiringAuth"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppSitePushSettingsSlot(ctx *pulumi.Context,
	name string, args *WebAppSitePushSettingsSlotArgs, opts ...pulumi.ResourceOption) (*WebAppSitePushSettingsSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IsPushEnabled == nil {
		return nil, errors.New("invalid value for required argument 'IsPushEnabled'")
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
			Type: pulumi.String("azure-native:web:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSitePushSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSitePushSettingsSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSitePushSettingsSlot
	err := ctx.RegisterResource("azure-native:web/v20200601:WebAppSitePushSettingsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSitePushSettingsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSitePushSettingsSlotState, opts ...pulumi.ResourceOption) (*WebAppSitePushSettingsSlot, error) {
	var resource WebAppSitePushSettingsSlot
	err := ctx.ReadResource("azure-native:web/v20200601:WebAppSitePushSettingsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppSitePushSettingsSlotState struct {
}

type WebAppSitePushSettingsSlotState struct {
}

func (WebAppSitePushSettingsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSitePushSettingsSlotState)(nil)).Elem()
}

type webAppSitePushSettingsSlotArgs struct {
	DynamicTagsJson   *string `pulumi:"dynamicTagsJson"`
	IsPushEnabled     bool    `pulumi:"isPushEnabled"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Slot              string  `pulumi:"slot"`
	TagWhitelistJson  *string `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
}


type WebAppSitePushSettingsSlotArgs struct {
	DynamicTagsJson   pulumi.StringPtrInput
	IsPushEnabled     pulumi.BoolInput
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
	TagWhitelistJson  pulumi.StringPtrInput
	TagsRequiringAuth pulumi.StringPtrInput
}

func (WebAppSitePushSettingsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSitePushSettingsSlotArgs)(nil)).Elem()
}

type WebAppSitePushSettingsSlotInput interface {
	pulumi.Input

	ToWebAppSitePushSettingsSlotOutput() WebAppSitePushSettingsSlotOutput
	ToWebAppSitePushSettingsSlotOutputWithContext(ctx context.Context) WebAppSitePushSettingsSlotOutput
}

func (*WebAppSitePushSettingsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSitePushSettingsSlot)(nil))
}

func (i *WebAppSitePushSettingsSlot) ToWebAppSitePushSettingsSlotOutput() WebAppSitePushSettingsSlotOutput {
	return i.ToWebAppSitePushSettingsSlotOutputWithContext(context.Background())
}

func (i *WebAppSitePushSettingsSlot) ToWebAppSitePushSettingsSlotOutputWithContext(ctx context.Context) WebAppSitePushSettingsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSitePushSettingsSlotOutput)
}

type WebAppSitePushSettingsSlotOutput struct{ *pulumi.OutputState }

func (WebAppSitePushSettingsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSitePushSettingsSlot)(nil))
}

func (o WebAppSitePushSettingsSlotOutput) ToWebAppSitePushSettingsSlotOutput() WebAppSitePushSettingsSlotOutput {
	return o
}

func (o WebAppSitePushSettingsSlotOutput) ToWebAppSitePushSettingsSlotOutputWithContext(ctx context.Context) WebAppSitePushSettingsSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppSitePushSettingsSlotOutput{})
}
