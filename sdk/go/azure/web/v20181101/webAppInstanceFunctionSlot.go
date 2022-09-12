


package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppInstanceFunctionSlot struct {
	pulumi.CustomResourceState

	Config             pulumi.AnyOutput       `pulumi:"config"`
	ConfigHref         pulumi.StringPtrOutput `pulumi:"configHref"`
	Files              pulumi.StringMapOutput `pulumi:"files"`
	FunctionAppId      pulumi.StringPtrOutput `pulumi:"functionAppId"`
	Href               pulumi.StringPtrOutput `pulumi:"href"`
	Kind               pulumi.StringPtrOutput `pulumi:"kind"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	ScriptHref         pulumi.StringPtrOutput `pulumi:"scriptHref"`
	ScriptRootPathHref pulumi.StringPtrOutput `pulumi:"scriptRootPathHref"`
	SecretsFileHref    pulumi.StringPtrOutput `pulumi:"secretsFileHref"`
	TestData           pulumi.StringPtrOutput `pulumi:"testData"`
	Type               pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppInstanceFunctionSlot(ctx *pulumi.Context,
	name string, args *WebAppInstanceFunctionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppInstanceFunctionSlot, error) {
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
			Type: pulumi.String("azure-native:web:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppInstanceFunctionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppInstanceFunctionSlot
	err := ctx.RegisterResource("azure-native:web/v20181101:WebAppInstanceFunctionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppInstanceFunctionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppInstanceFunctionSlotState, opts ...pulumi.ResourceOption) (*WebAppInstanceFunctionSlot, error) {
	var resource WebAppInstanceFunctionSlot
	err := ctx.ReadResource("azure-native:web/v20181101:WebAppInstanceFunctionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppInstanceFunctionSlotState struct {
}

type WebAppInstanceFunctionSlotState struct {
}

func (WebAppInstanceFunctionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppInstanceFunctionSlotState)(nil)).Elem()
}

type webAppInstanceFunctionSlotArgs struct {
	Config             interface{}       `pulumi:"config"`
	ConfigHref         *string           `pulumi:"configHref"`
	Files              map[string]string `pulumi:"files"`
	FunctionAppId      *string           `pulumi:"functionAppId"`
	FunctionName       *string           `pulumi:"functionName"`
	Href               *string           `pulumi:"href"`
	Kind               *string           `pulumi:"kind"`
	Name               string            `pulumi:"name"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	ScriptHref         *string           `pulumi:"scriptHref"`
	ScriptRootPathHref *string           `pulumi:"scriptRootPathHref"`
	SecretsFileHref    *string           `pulumi:"secretsFileHref"`
	Slot               string            `pulumi:"slot"`
	TestData           *string           `pulumi:"testData"`
}


type WebAppInstanceFunctionSlotArgs struct {
	Config             pulumi.Input
	ConfigHref         pulumi.StringPtrInput
	Files              pulumi.StringMapInput
	FunctionAppId      pulumi.StringPtrInput
	FunctionName       pulumi.StringPtrInput
	Href               pulumi.StringPtrInput
	Kind               pulumi.StringPtrInput
	Name               pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ScriptHref         pulumi.StringPtrInput
	ScriptRootPathHref pulumi.StringPtrInput
	SecretsFileHref    pulumi.StringPtrInput
	Slot               pulumi.StringInput
	TestData           pulumi.StringPtrInput
}

func (WebAppInstanceFunctionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppInstanceFunctionSlotArgs)(nil)).Elem()
}

type WebAppInstanceFunctionSlotInput interface {
	pulumi.Input

	ToWebAppInstanceFunctionSlotOutput() WebAppInstanceFunctionSlotOutput
	ToWebAppInstanceFunctionSlotOutputWithContext(ctx context.Context) WebAppInstanceFunctionSlotOutput
}

func (*WebAppInstanceFunctionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppInstanceFunctionSlot)(nil)).Elem()
}

func (i *WebAppInstanceFunctionSlot) ToWebAppInstanceFunctionSlotOutput() WebAppInstanceFunctionSlotOutput {
	return i.ToWebAppInstanceFunctionSlotOutputWithContext(context.Background())
}

func (i *WebAppInstanceFunctionSlot) ToWebAppInstanceFunctionSlotOutputWithContext(ctx context.Context) WebAppInstanceFunctionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppInstanceFunctionSlotOutput)
}

type WebAppInstanceFunctionSlotOutput struct{ *pulumi.OutputState }

func (WebAppInstanceFunctionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppInstanceFunctionSlot)(nil)).Elem()
}

func (o WebAppInstanceFunctionSlotOutput) ToWebAppInstanceFunctionSlotOutput() WebAppInstanceFunctionSlotOutput {
	return o
}

func (o WebAppInstanceFunctionSlotOutput) ToWebAppInstanceFunctionSlotOutputWithContext(ctx context.Context) WebAppInstanceFunctionSlotOutput {
	return o
}

func (o WebAppInstanceFunctionSlotOutput) Config() pulumi.AnyOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.AnyOutput { return v.Config }).(pulumi.AnyOutput)
}

func (o WebAppInstanceFunctionSlotOutput) ConfigHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringPtrOutput { return v.ConfigHref }).(pulumi.StringPtrOutput)
}

func (o WebAppInstanceFunctionSlotOutput) Files() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringMapOutput { return v.Files }).(pulumi.StringMapOutput)
}

func (o WebAppInstanceFunctionSlotOutput) FunctionAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringPtrOutput { return v.FunctionAppId }).(pulumi.StringPtrOutput)
}

func (o WebAppInstanceFunctionSlotOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringPtrOutput { return v.Href }).(pulumi.StringPtrOutput)
}

func (o WebAppInstanceFunctionSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppInstanceFunctionSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppInstanceFunctionSlotOutput) ScriptHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringPtrOutput { return v.ScriptHref }).(pulumi.StringPtrOutput)
}

func (o WebAppInstanceFunctionSlotOutput) ScriptRootPathHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringPtrOutput { return v.ScriptRootPathHref }).(pulumi.StringPtrOutput)
}

func (o WebAppInstanceFunctionSlotOutput) SecretsFileHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringPtrOutput { return v.SecretsFileHref }).(pulumi.StringPtrOutput)
}

func (o WebAppInstanceFunctionSlotOutput) TestData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringPtrOutput { return v.TestData }).(pulumi.StringPtrOutput)
}

func (o WebAppInstanceFunctionSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppInstanceFunctionSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppInstanceFunctionSlotOutput{})
}
