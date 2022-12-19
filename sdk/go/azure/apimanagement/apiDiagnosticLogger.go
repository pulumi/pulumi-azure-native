


package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiDiagnosticLogger struct {
	pulumi.CustomResourceState

	Credentials pulumi.StringMapOutput `pulumi:"credentials"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	IsBuffered  pulumi.BoolPtrOutput   `pulumi:"isBuffered"`
	LoggerType  pulumi.StringOutput    `pulumi:"loggerType"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewApiDiagnosticLogger(ctx *pulumi.Context,
	name string, args *ApiDiagnosticLoggerArgs, opts ...pulumi.ResourceOption) (*ApiDiagnosticLogger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.DiagnosticId == nil {
		return nil, errors.New("invalid value for required argument 'DiagnosticId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiDiagnosticLogger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiDiagnosticLogger"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiDiagnosticLogger
	err := ctx.RegisterResource("azure-native:apimanagement:ApiDiagnosticLogger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiDiagnosticLogger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiDiagnosticLoggerState, opts ...pulumi.ResourceOption) (*ApiDiagnosticLogger, error) {
	var resource ApiDiagnosticLogger
	err := ctx.ReadResource("azure-native:apimanagement:ApiDiagnosticLogger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiDiagnosticLoggerState struct {
}

type ApiDiagnosticLoggerState struct {
}

func (ApiDiagnosticLoggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticLoggerState)(nil)).Elem()
}

type apiDiagnosticLoggerArgs struct {
	ApiId             string  `pulumi:"apiId"`
	DiagnosticId      string  `pulumi:"diagnosticId"`
	Loggerid          *string `pulumi:"loggerid"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type ApiDiagnosticLoggerArgs struct {
	ApiId             pulumi.StringInput
	DiagnosticId      pulumi.StringInput
	Loggerid          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (ApiDiagnosticLoggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticLoggerArgs)(nil)).Elem()
}

type ApiDiagnosticLoggerInput interface {
	pulumi.Input

	ToApiDiagnosticLoggerOutput() ApiDiagnosticLoggerOutput
	ToApiDiagnosticLoggerOutputWithContext(ctx context.Context) ApiDiagnosticLoggerOutput
}

func (*ApiDiagnosticLogger) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDiagnosticLogger)(nil)).Elem()
}

func (i *ApiDiagnosticLogger) ToApiDiagnosticLoggerOutput() ApiDiagnosticLoggerOutput {
	return i.ToApiDiagnosticLoggerOutputWithContext(context.Background())
}

func (i *ApiDiagnosticLogger) ToApiDiagnosticLoggerOutputWithContext(ctx context.Context) ApiDiagnosticLoggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDiagnosticLoggerOutput)
}

type ApiDiagnosticLoggerOutput struct{ *pulumi.OutputState }

func (ApiDiagnosticLoggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDiagnosticLogger)(nil)).Elem()
}

func (o ApiDiagnosticLoggerOutput) ToApiDiagnosticLoggerOutput() ApiDiagnosticLoggerOutput {
	return o
}

func (o ApiDiagnosticLoggerOutput) ToApiDiagnosticLoggerOutputWithContext(ctx context.Context) ApiDiagnosticLoggerOutput {
	return o
}

func (o ApiDiagnosticLoggerOutput) Credentials() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiDiagnosticLogger) pulumi.StringMapOutput { return v.Credentials }).(pulumi.StringMapOutput)
}

func (o ApiDiagnosticLoggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDiagnosticLogger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiDiagnosticLoggerOutput) IsBuffered() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiDiagnosticLogger) pulumi.BoolPtrOutput { return v.IsBuffered }).(pulumi.BoolPtrOutput)
}

func (o ApiDiagnosticLoggerOutput) LoggerType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDiagnosticLogger) pulumi.StringOutput { return v.LoggerType }).(pulumi.StringOutput)
}

func (o ApiDiagnosticLoggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDiagnosticLogger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiDiagnosticLoggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDiagnosticLogger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiDiagnosticLoggerOutput{})
}
