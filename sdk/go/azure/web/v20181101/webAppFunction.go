


package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppFunction struct {
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
			Type: pulumi.String("azure-native:web/v20190801:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppFunction"),
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
	})
	opts = append(opts, aliases)
	var resource WebAppFunction
	err := ctx.RegisterResource("azure-native:web/v20181101:WebAppFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppFunctionState, opts ...pulumi.ResourceOption) (*WebAppFunction, error) {
	var resource WebAppFunction
	err := ctx.ReadResource("azure-native:web/v20181101:WebAppFunction", name, id, state, &resource, opts...)
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
	Kind               *string           `pulumi:"kind"`
	Name               string            `pulumi:"name"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	ScriptHref         *string           `pulumi:"scriptHref"`
	ScriptRootPathHref *string           `pulumi:"scriptRootPathHref"`
	SecretsFileHref    *string           `pulumi:"secretsFileHref"`
	TestData           *string           `pulumi:"testData"`
}


type WebAppFunctionArgs struct {
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
	TestData           pulumi.StringPtrInput
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
	return reflect.TypeOf((*WebAppFunction)(nil))
}

func (i *WebAppFunction) ToWebAppFunctionOutput() WebAppFunctionOutput {
	return i.ToWebAppFunctionOutputWithContext(context.Background())
}

func (i *WebAppFunction) ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppFunctionOutput)
}

type WebAppFunctionOutput struct{ *pulumi.OutputState }

func (WebAppFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppFunction)(nil))
}

func (o WebAppFunctionOutput) ToWebAppFunctionOutput() WebAppFunctionOutput {
	return o
}

func (o WebAppFunctionOutput) ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppFunctionOutput{})
}
