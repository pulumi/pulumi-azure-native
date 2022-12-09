


package v20210115

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppConnectionStrings struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput                   `pulumi:"kind"`
	Name       pulumi.StringOutput                      `pulumi:"name"`
	Properties ConnStringValueTypePairResponseMapOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                      `pulumi:"type"`
}


func NewWebAppConnectionStrings(ctx *pulumi.Context,
	name string, args *WebAppConnectionStringsArgs, opts ...pulumi.ResourceOption) (*WebAppConnectionStrings, error) {
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
			Type: pulumi.String("azure-native:web:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppConnectionStrings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppConnectionStrings"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppConnectionStrings
	err := ctx.RegisterResource("azure-native:web/v20210115:WebAppConnectionStrings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppConnectionStrings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppConnectionStringsState, opts ...pulumi.ResourceOption) (*WebAppConnectionStrings, error) {
	var resource WebAppConnectionStrings
	err := ctx.ReadResource("azure-native:web/v20210115:WebAppConnectionStrings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppConnectionStringsState struct {
}

type WebAppConnectionStringsState struct {
}

func (WebAppConnectionStringsState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppConnectionStringsState)(nil)).Elem()
}

type webAppConnectionStringsArgs struct {
	Kind              *string                            `pulumi:"kind"`
	Name              string                             `pulumi:"name"`
	Properties        map[string]ConnStringValueTypePair `pulumi:"properties"`
	ResourceGroupName string                             `pulumi:"resourceGroupName"`
}


type WebAppConnectionStringsArgs struct {
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        ConnStringValueTypePairMapInput
	ResourceGroupName pulumi.StringInput
}

func (WebAppConnectionStringsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppConnectionStringsArgs)(nil)).Elem()
}

type WebAppConnectionStringsInput interface {
	pulumi.Input

	ToWebAppConnectionStringsOutput() WebAppConnectionStringsOutput
	ToWebAppConnectionStringsOutputWithContext(ctx context.Context) WebAppConnectionStringsOutput
}

func (*WebAppConnectionStrings) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppConnectionStrings)(nil)).Elem()
}

func (i *WebAppConnectionStrings) ToWebAppConnectionStringsOutput() WebAppConnectionStringsOutput {
	return i.ToWebAppConnectionStringsOutputWithContext(context.Background())
}

func (i *WebAppConnectionStrings) ToWebAppConnectionStringsOutputWithContext(ctx context.Context) WebAppConnectionStringsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppConnectionStringsOutput)
}

type WebAppConnectionStringsOutput struct{ *pulumi.OutputState }

func (WebAppConnectionStringsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppConnectionStrings)(nil)).Elem()
}

func (o WebAppConnectionStringsOutput) ToWebAppConnectionStringsOutput() WebAppConnectionStringsOutput {
	return o
}

func (o WebAppConnectionStringsOutput) ToWebAppConnectionStringsOutputWithContext(ctx context.Context) WebAppConnectionStringsOutput {
	return o
}

func (o WebAppConnectionStringsOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppConnectionStrings) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppConnectionStringsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppConnectionStrings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppConnectionStringsOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v *WebAppConnectionStrings) ConnStringValueTypePairResponseMapOutput { return v.Properties }).(ConnStringValueTypePairResponseMapOutput)
}

func (o WebAppConnectionStringsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppConnectionStrings) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppConnectionStringsOutput{})
}
