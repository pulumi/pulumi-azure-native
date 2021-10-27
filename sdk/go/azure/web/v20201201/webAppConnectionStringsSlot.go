


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppConnectionStringsSlot struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput                   `pulumi:"kind"`
	Name       pulumi.StringOutput                      `pulumi:"name"`
	Properties ConnStringValueTypePairResponseMapOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                      `pulumi:"type"`
}


func NewWebAppConnectionStringsSlot(ctx *pulumi.Context,
	name string, args *WebAppConnectionStringsSlotArgs, opts ...pulumi.ResourceOption) (*WebAppConnectionStringsSlot, error) {
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
			Type: pulumi.String("azure-nextgen:web/v20201201:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:WebAppConnectionStringsSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppConnectionStringsSlot
	err := ctx.RegisterResource("azure-native:web/v20201201:WebAppConnectionStringsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppConnectionStringsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppConnectionStringsSlotState, opts ...pulumi.ResourceOption) (*WebAppConnectionStringsSlot, error) {
	var resource WebAppConnectionStringsSlot
	err := ctx.ReadResource("azure-native:web/v20201201:WebAppConnectionStringsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppConnectionStringsSlotState struct {
}

type WebAppConnectionStringsSlotState struct {
}

func (WebAppConnectionStringsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppConnectionStringsSlotState)(nil)).Elem()
}

type webAppConnectionStringsSlotArgs struct {
	Kind              *string                            `pulumi:"kind"`
	Name              string                             `pulumi:"name"`
	Properties        map[string]ConnStringValueTypePair `pulumi:"properties"`
	ResourceGroupName string                             `pulumi:"resourceGroupName"`
	Slot              string                             `pulumi:"slot"`
}


type WebAppConnectionStringsSlotArgs struct {
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        ConnStringValueTypePairMapInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
}

func (WebAppConnectionStringsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppConnectionStringsSlotArgs)(nil)).Elem()
}

type WebAppConnectionStringsSlotInput interface {
	pulumi.Input

	ToWebAppConnectionStringsSlotOutput() WebAppConnectionStringsSlotOutput
	ToWebAppConnectionStringsSlotOutputWithContext(ctx context.Context) WebAppConnectionStringsSlotOutput
}

func (*WebAppConnectionStringsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppConnectionStringsSlot)(nil))
}

func (i *WebAppConnectionStringsSlot) ToWebAppConnectionStringsSlotOutput() WebAppConnectionStringsSlotOutput {
	return i.ToWebAppConnectionStringsSlotOutputWithContext(context.Background())
}

func (i *WebAppConnectionStringsSlot) ToWebAppConnectionStringsSlotOutputWithContext(ctx context.Context) WebAppConnectionStringsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppConnectionStringsSlotOutput)
}

type WebAppConnectionStringsSlotOutput struct{ *pulumi.OutputState }

func (WebAppConnectionStringsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppConnectionStringsSlot)(nil))
}

func (o WebAppConnectionStringsSlotOutput) ToWebAppConnectionStringsSlotOutput() WebAppConnectionStringsSlotOutput {
	return o
}

func (o WebAppConnectionStringsSlotOutput) ToWebAppConnectionStringsSlotOutputWithContext(ctx context.Context) WebAppConnectionStringsSlotOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebAppConnectionStringsSlotInput)(nil)).Elem(), &WebAppConnectionStringsSlot{})
	pulumi.RegisterOutputType(WebAppConnectionStringsSlotOutput{})
}
