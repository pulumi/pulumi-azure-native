


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppApplicationSettingsSlot struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput   `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	Properties pulumi.StringMapOutput   `pulumi:"properties"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewWebAppApplicationSettingsSlot(ctx *pulumi.Context,
	name string, args *WebAppApplicationSettingsSlotArgs, opts ...pulumi.ResourceOption) (*WebAppApplicationSettingsSlot, error) {
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
			Type: pulumi.String("azure-native:web:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppApplicationSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppApplicationSettingsSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppApplicationSettingsSlot
	err := ctx.RegisterResource("azure-native:web/v20200901:WebAppApplicationSettingsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppApplicationSettingsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppApplicationSettingsSlotState, opts ...pulumi.ResourceOption) (*WebAppApplicationSettingsSlot, error) {
	var resource WebAppApplicationSettingsSlot
	err := ctx.ReadResource("azure-native:web/v20200901:WebAppApplicationSettingsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppApplicationSettingsSlotState struct {
}

type WebAppApplicationSettingsSlotState struct {
}

func (WebAppApplicationSettingsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppApplicationSettingsSlotState)(nil)).Elem()
}

type webAppApplicationSettingsSlotArgs struct {
	Kind              *string           `pulumi:"kind"`
	Name              string            `pulumi:"name"`
	Properties        map[string]string `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Slot              string            `pulumi:"slot"`
}


type WebAppApplicationSettingsSlotArgs struct {
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
}

func (WebAppApplicationSettingsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppApplicationSettingsSlotArgs)(nil)).Elem()
}

type WebAppApplicationSettingsSlotInput interface {
	pulumi.Input

	ToWebAppApplicationSettingsSlotOutput() WebAppApplicationSettingsSlotOutput
	ToWebAppApplicationSettingsSlotOutputWithContext(ctx context.Context) WebAppApplicationSettingsSlotOutput
}

func (*WebAppApplicationSettingsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppApplicationSettingsSlot)(nil)).Elem()
}

func (i *WebAppApplicationSettingsSlot) ToWebAppApplicationSettingsSlotOutput() WebAppApplicationSettingsSlotOutput {
	return i.ToWebAppApplicationSettingsSlotOutputWithContext(context.Background())
}

func (i *WebAppApplicationSettingsSlot) ToWebAppApplicationSettingsSlotOutputWithContext(ctx context.Context) WebAppApplicationSettingsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppApplicationSettingsSlotOutput)
}

type WebAppApplicationSettingsSlotOutput struct{ *pulumi.OutputState }

func (WebAppApplicationSettingsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppApplicationSettingsSlot)(nil)).Elem()
}

func (o WebAppApplicationSettingsSlotOutput) ToWebAppApplicationSettingsSlotOutput() WebAppApplicationSettingsSlotOutput {
	return o
}

func (o WebAppApplicationSettingsSlotOutput) ToWebAppApplicationSettingsSlotOutputWithContext(ctx context.Context) WebAppApplicationSettingsSlotOutput {
	return o
}

func (o WebAppApplicationSettingsSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppApplicationSettingsSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppApplicationSettingsSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppApplicationSettingsSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppApplicationSettingsSlotOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppApplicationSettingsSlot) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

func (o WebAppApplicationSettingsSlotOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppApplicationSettingsSlot) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WebAppApplicationSettingsSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppApplicationSettingsSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppApplicationSettingsSlotOutput{})
}
