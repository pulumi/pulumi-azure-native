


package v20180101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Diagnostic struct {
	pulumi.CustomResourceState

	Enabled pulumi.BoolOutput   `pulumi:"enabled"`
	Name    pulumi.StringOutput `pulumi:"name"`
	Type    pulumi.StringOutput `pulumi:"type"`
}


func NewDiagnostic(ctx *pulumi.Context,
	name string, args *DiagnosticArgs, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Diagnostic"),
		},
	})
	opts = append(opts, aliases)
	var resource Diagnostic
	err := ctx.RegisterResource("azure-native:apimanagement/v20180101:Diagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticState, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	var resource Diagnostic
	err := ctx.ReadResource("azure-native:apimanagement/v20180101:Diagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diagnosticState struct {
}

type DiagnosticState struct {
}

func (DiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticState)(nil)).Elem()
}

type diagnosticArgs struct {
	DiagnosticId      *string `pulumi:"diagnosticId"`
	Enabled           bool    `pulumi:"enabled"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type DiagnosticArgs struct {
	DiagnosticId      pulumi.StringPtrInput
	Enabled           pulumi.BoolInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (DiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticArgs)(nil)).Elem()
}

type DiagnosticInput interface {
	pulumi.Input

	ToDiagnosticOutput() DiagnosticOutput
	ToDiagnosticOutputWithContext(ctx context.Context) DiagnosticOutput
}

func (*Diagnostic) ElementType() reflect.Type {
	return reflect.TypeOf((**Diagnostic)(nil)).Elem()
}

func (i *Diagnostic) ToDiagnosticOutput() DiagnosticOutput {
	return i.ToDiagnosticOutputWithContext(context.Background())
}

func (i *Diagnostic) ToDiagnosticOutputWithContext(ctx context.Context) DiagnosticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticOutput)
}

type DiagnosticOutput struct{ *pulumi.OutputState }

func (DiagnosticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Diagnostic)(nil)).Elem()
}

func (o DiagnosticOutput) ToDiagnosticOutput() DiagnosticOutput {
	return o
}

func (o DiagnosticOutput) ToDiagnosticOutputWithContext(ctx context.Context) DiagnosticOutput {
	return o
}

func (o DiagnosticOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Diagnostic) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o DiagnosticOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Diagnostic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiagnosticOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Diagnostic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DiagnosticOutput{})
}
