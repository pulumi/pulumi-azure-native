


package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppFunction struct {
	pulumi.CustomResourceState

	Config             pulumi.AnyOutput         `pulumi:"config"`
	ConfigHref         pulumi.StringPtrOutput   `pulumi:"configHref"`
	Files              pulumi.StringMapOutput   `pulumi:"files"`
	FunctionAppId      pulumi.StringPtrOutput   `pulumi:"functionAppId"`
	Href               pulumi.StringPtrOutput   `pulumi:"href"`
	InvokeUrlTemplate  pulumi.StringPtrOutput   `pulumi:"invokeUrlTemplate"`
	IsDisabled         pulumi.BoolPtrOutput     `pulumi:"isDisabled"`
	Kind               pulumi.StringPtrOutput   `pulumi:"kind"`
	Language           pulumi.StringPtrOutput   `pulumi:"language"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ScriptHref         pulumi.StringPtrOutput   `pulumi:"scriptHref"`
	ScriptRootPathHref pulumi.StringPtrOutput   `pulumi:"scriptRootPathHref"`
	SecretsFileHref    pulumi.StringPtrOutput   `pulumi:"secretsFileHref"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	TestData           pulumi.StringPtrOutput   `pulumi:"testData"`
	TestDataHref       pulumi.StringPtrOutput   `pulumi:"testDataHref"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewWebAppFunction(ctx *pulumi.Context,
	name string, args *WebAppFunctionArgs, opts ...pulumi.ResourceOption) (*WebAppFunction, error) {
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
			Type: pulumi.String("azure-native:web:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppFunction"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppFunction
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppFunctionState, opts ...pulumi.ResourceOption) (*WebAppFunction, error) {
	var resource WebAppFunction
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppFunctionState struct {
}

type WebAppFunctionState struct {
}

func (WebAppFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppFunctionState)(nil)).Elem()
}

type webAppFunctionArgs struct {
	Config             interface{}       `pulumi:"config"`
	ConfigHref         *string           `pulumi:"configHref"`
	Files              map[string]string `pulumi:"files"`
	FunctionAppId      *string           `pulumi:"functionAppId"`
	FunctionName       *string           `pulumi:"functionName"`
	Href               *string           `pulumi:"href"`
	InvokeUrlTemplate  *string           `pulumi:"invokeUrlTemplate"`
	IsDisabled         *bool             `pulumi:"isDisabled"`
	Kind               *string           `pulumi:"kind"`
	Language           *string           `pulumi:"language"`
	Name               string            `pulumi:"name"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	ScriptHref         *string           `pulumi:"scriptHref"`
	ScriptRootPathHref *string           `pulumi:"scriptRootPathHref"`
	SecretsFileHref    *string           `pulumi:"secretsFileHref"`
	TestData           *string           `pulumi:"testData"`
	TestDataHref       *string           `pulumi:"testDataHref"`
}


type WebAppFunctionArgs struct {
	Config             pulumi.Input
	ConfigHref         pulumi.StringPtrInput
	Files              pulumi.StringMapInput
	FunctionAppId      pulumi.StringPtrInput
	FunctionName       pulumi.StringPtrInput
	Href               pulumi.StringPtrInput
	InvokeUrlTemplate  pulumi.StringPtrInput
	IsDisabled         pulumi.BoolPtrInput
	Kind               pulumi.StringPtrInput
	Language           pulumi.StringPtrInput
	Name               pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ScriptHref         pulumi.StringPtrInput
	ScriptRootPathHref pulumi.StringPtrInput
	SecretsFileHref    pulumi.StringPtrInput
	TestData           pulumi.StringPtrInput
	TestDataHref       pulumi.StringPtrInput
}

func (WebAppFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppFunctionArgs)(nil)).Elem()
}

type WebAppFunctionInput interface {
	pulumi.Input

	ToWebAppFunctionOutput() WebAppFunctionOutput
	ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput
}

func (*WebAppFunction) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppFunction)(nil)).Elem()
}

func (i *WebAppFunction) ToWebAppFunctionOutput() WebAppFunctionOutput {
	return i.ToWebAppFunctionOutputWithContext(context.Background())
}

func (i *WebAppFunction) ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppFunctionOutput)
}

type WebAppFunctionOutput struct{ *pulumi.OutputState }

func (WebAppFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppFunction)(nil)).Elem()
}

func (o WebAppFunctionOutput) ToWebAppFunctionOutput() WebAppFunctionOutput {
	return o
}

func (o WebAppFunctionOutput) ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput {
	return o
}

func (o WebAppFunctionOutput) Config() pulumi.AnyOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.AnyOutput { return v.Config }).(pulumi.AnyOutput)
}

func (o WebAppFunctionOutput) ConfigHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.ConfigHref }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) Files() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringMapOutput { return v.Files }).(pulumi.StringMapOutput)
}

func (o WebAppFunctionOutput) FunctionAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.FunctionAppId }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.Href }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) InvokeUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.InvokeUrlTemplate }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.BoolPtrOutput { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppFunctionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.Language }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppFunctionOutput) ScriptHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.ScriptHref }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) ScriptRootPathHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.ScriptRootPathHref }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) SecretsFileHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.SecretsFileHref }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppFunction) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WebAppFunctionOutput) TestData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.TestData }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) TestDataHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.TestDataHref }).(pulumi.StringPtrOutput)
}

func (o WebAppFunctionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppFunctionOutput{})
}
