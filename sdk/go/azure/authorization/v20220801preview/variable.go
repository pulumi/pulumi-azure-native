


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Variable struct {
	pulumi.CustomResourceState

	Columns    PolicyVariableColumnResponseArrayOutput `pulumi:"columns"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewVariable(ctx *pulumi.Context,
	name string, args *VariableArgs, opts ...pulumi.ResourceOption) (*Variable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Columns == nil {
		return nil, errors.New("invalid value for required argument 'Columns'")
	}
	var resource Variable
	err := ctx.RegisterResource("azure-native:authorization/v20220801preview:Variable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableState, opts ...pulumi.ResourceOption) (*Variable, error) {
	var resource Variable
	err := ctx.ReadResource("azure-native:authorization/v20220801preview:Variable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type variableState struct {
}

type VariableState struct {
}

func (VariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*variableState)(nil)).Elem()
}

type variableArgs struct {
	Columns      []PolicyVariableColumn `pulumi:"columns"`
	VariableName *string                `pulumi:"variableName"`
}


type VariableArgs struct {
	Columns      PolicyVariableColumnArrayInput
	VariableName pulumi.StringPtrInput
}

func (VariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variableArgs)(nil)).Elem()
}

type VariableInput interface {
	pulumi.Input

	ToVariableOutput() VariableOutput
	ToVariableOutputWithContext(ctx context.Context) VariableOutput
}

func (*Variable) ElementType() reflect.Type {
	return reflect.TypeOf((**Variable)(nil)).Elem()
}

func (i *Variable) ToVariableOutput() VariableOutput {
	return i.ToVariableOutputWithContext(context.Background())
}

func (i *Variable) ToVariableOutputWithContext(ctx context.Context) VariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableOutput)
}

type VariableOutput struct{ *pulumi.OutputState }

func (VariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Variable)(nil)).Elem()
}

func (o VariableOutput) ToVariableOutput() VariableOutput {
	return o
}

func (o VariableOutput) ToVariableOutputWithContext(ctx context.Context) VariableOutput {
	return o
}

func (o VariableOutput) Columns() PolicyVariableColumnResponseArrayOutput {
	return o.ApplyT(func(v *Variable) PolicyVariableColumnResponseArrayOutput { return v.Columns }).(PolicyVariableColumnResponseArrayOutput)
}

func (o VariableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VariableOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Variable) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VariableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VariableOutput{})
}
