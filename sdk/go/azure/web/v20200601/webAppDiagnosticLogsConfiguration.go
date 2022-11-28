


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppDiagnosticLogsConfiguration struct {
	pulumi.CustomResourceState

	ApplicationLogs       ApplicationLogsConfigResponsePtrOutput `pulumi:"applicationLogs"`
	DetailedErrorMessages EnabledConfigResponsePtrOutput         `pulumi:"detailedErrorMessages"`
	FailedRequestsTracing EnabledConfigResponsePtrOutput         `pulumi:"failedRequestsTracing"`
	HttpLogs              HttpLogsConfigResponsePtrOutput        `pulumi:"httpLogs"`
	Kind                  pulumi.StringPtrOutput                 `pulumi:"kind"`
	Name                  pulumi.StringOutput                    `pulumi:"name"`
	Type                  pulumi.StringOutput                    `pulumi:"type"`
}


func NewWebAppDiagnosticLogsConfiguration(ctx *pulumi.Context,
	name string, args *WebAppDiagnosticLogsConfigurationArgs, opts ...pulumi.ResourceOption) (*WebAppDiagnosticLogsConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ApplicationLogs != nil {
		args.ApplicationLogs = args.ApplicationLogs.ToApplicationLogsConfigPtrOutput().ApplyT(func(v *ApplicationLogsConfig) *ApplicationLogsConfig { return v.Defaults() }).(ApplicationLogsConfigPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppDiagnosticLogsConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppDiagnosticLogsConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppDiagnosticLogsConfiguration
	err := ctx.RegisterResource("azure-native:web/v20200601:WebAppDiagnosticLogsConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppDiagnosticLogsConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppDiagnosticLogsConfigurationState, opts ...pulumi.ResourceOption) (*WebAppDiagnosticLogsConfiguration, error) {
	var resource WebAppDiagnosticLogsConfiguration
	err := ctx.ReadResource("azure-native:web/v20200601:WebAppDiagnosticLogsConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppDiagnosticLogsConfigurationState struct {
}

type WebAppDiagnosticLogsConfigurationState struct {
}

func (WebAppDiagnosticLogsConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDiagnosticLogsConfigurationState)(nil)).Elem()
}

type webAppDiagnosticLogsConfigurationArgs struct {
	ApplicationLogs       *ApplicationLogsConfig `pulumi:"applicationLogs"`
	DetailedErrorMessages *EnabledConfig         `pulumi:"detailedErrorMessages"`
	FailedRequestsTracing *EnabledConfig         `pulumi:"failedRequestsTracing"`
	HttpLogs              *HttpLogsConfig        `pulumi:"httpLogs"`
	Kind                  *string                `pulumi:"kind"`
	Name                  string                 `pulumi:"name"`
	ResourceGroupName     string                 `pulumi:"resourceGroupName"`
}


type WebAppDiagnosticLogsConfigurationArgs struct {
	ApplicationLogs       ApplicationLogsConfigPtrInput
	DetailedErrorMessages EnabledConfigPtrInput
	FailedRequestsTracing EnabledConfigPtrInput
	HttpLogs              HttpLogsConfigPtrInput
	Kind                  pulumi.StringPtrInput
	Name                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
}

func (WebAppDiagnosticLogsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDiagnosticLogsConfigurationArgs)(nil)).Elem()
}

type WebAppDiagnosticLogsConfigurationInput interface {
	pulumi.Input

	ToWebAppDiagnosticLogsConfigurationOutput() WebAppDiagnosticLogsConfigurationOutput
	ToWebAppDiagnosticLogsConfigurationOutputWithContext(ctx context.Context) WebAppDiagnosticLogsConfigurationOutput
}

func (*WebAppDiagnosticLogsConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppDiagnosticLogsConfiguration)(nil)).Elem()
}

func (i *WebAppDiagnosticLogsConfiguration) ToWebAppDiagnosticLogsConfigurationOutput() WebAppDiagnosticLogsConfigurationOutput {
	return i.ToWebAppDiagnosticLogsConfigurationOutputWithContext(context.Background())
}

func (i *WebAppDiagnosticLogsConfiguration) ToWebAppDiagnosticLogsConfigurationOutputWithContext(ctx context.Context) WebAppDiagnosticLogsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppDiagnosticLogsConfigurationOutput)
}

type WebAppDiagnosticLogsConfigurationOutput struct{ *pulumi.OutputState }

func (WebAppDiagnosticLogsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppDiagnosticLogsConfiguration)(nil)).Elem()
}

func (o WebAppDiagnosticLogsConfigurationOutput) ToWebAppDiagnosticLogsConfigurationOutput() WebAppDiagnosticLogsConfigurationOutput {
	return o
}

func (o WebAppDiagnosticLogsConfigurationOutput) ToWebAppDiagnosticLogsConfigurationOutputWithContext(ctx context.Context) WebAppDiagnosticLogsConfigurationOutput {
	return o
}

func (o WebAppDiagnosticLogsConfigurationOutput) ApplicationLogs() ApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *WebAppDiagnosticLogsConfiguration) ApplicationLogsConfigResponsePtrOutput {
		return v.ApplicationLogs
	}).(ApplicationLogsConfigResponsePtrOutput)
}

func (o WebAppDiagnosticLogsConfigurationOutput) DetailedErrorMessages() EnabledConfigResponsePtrOutput {
	return o.ApplyT(func(v *WebAppDiagnosticLogsConfiguration) EnabledConfigResponsePtrOutput {
		return v.DetailedErrorMessages
	}).(EnabledConfigResponsePtrOutput)
}

func (o WebAppDiagnosticLogsConfigurationOutput) FailedRequestsTracing() EnabledConfigResponsePtrOutput {
	return o.ApplyT(func(v *WebAppDiagnosticLogsConfiguration) EnabledConfigResponsePtrOutput {
		return v.FailedRequestsTracing
	}).(EnabledConfigResponsePtrOutput)
}

func (o WebAppDiagnosticLogsConfigurationOutput) HttpLogs() HttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *WebAppDiagnosticLogsConfiguration) HttpLogsConfigResponsePtrOutput { return v.HttpLogs }).(HttpLogsConfigResponsePtrOutput)
}

func (o WebAppDiagnosticLogsConfigurationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppDiagnosticLogsConfiguration) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppDiagnosticLogsConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppDiagnosticLogsConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppDiagnosticLogsConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppDiagnosticLogsConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppDiagnosticLogsConfigurationOutput{})
}
