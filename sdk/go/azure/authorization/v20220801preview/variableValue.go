


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VariableValue struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                               `pulumi:"name"`
	SystemData SystemDataResponseOutput                          `pulumi:"systemData"`
	Type       pulumi.StringOutput                               `pulumi:"type"`
	Values     PolicyVariableValueColumnValueResponseArrayOutput `pulumi:"values"`
}


func NewVariableValue(ctx *pulumi.Context,
	name string, args *VariableValueArgs, opts ...pulumi.ResourceOption) (*VariableValue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Values == nil {
		return nil, errors.New("invalid value for required argument 'Values'")
	}
	if args.VariableName == nil {
		return nil, errors.New("invalid value for required argument 'VariableName'")
	}
	var resource VariableValue
	err := ctx.RegisterResource("azure-native:authorization/v20220801preview:VariableValue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVariableValue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableValueState, opts ...pulumi.ResourceOption) (*VariableValue, error) {
	var resource VariableValue
	err := ctx.ReadResource("azure-native:authorization/v20220801preview:VariableValue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type variableValueState struct {
}

type VariableValueState struct {
}

func (VariableValueState) ElementType() reflect.Type {
	return reflect.TypeOf((*variableValueState)(nil)).Elem()
}

type variableValueArgs struct {
	Values            []PolicyVariableValueColumnValue `pulumi:"values"`
	VariableName      string                           `pulumi:"variableName"`
	VariableValueName *string                          `pulumi:"variableValueName"`
}


type VariableValueArgs struct {
	Values            PolicyVariableValueColumnValueArrayInput
	VariableName      pulumi.StringInput
	VariableValueName pulumi.StringPtrInput
}

func (VariableValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variableValueArgs)(nil)).Elem()
}

type VariableValueInput interface {
	pulumi.Input

	ToVariableValueOutput() VariableValueOutput
	ToVariableValueOutputWithContext(ctx context.Context) VariableValueOutput
}

func (*VariableValue) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableValue)(nil)).Elem()
}

func (i *VariableValue) ToVariableValueOutput() VariableValueOutput {
	return i.ToVariableValueOutputWithContext(context.Background())
}

func (i *VariableValue) ToVariableValueOutputWithContext(ctx context.Context) VariableValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableValueOutput)
}

type VariableValueOutput struct{ *pulumi.OutputState }

func (VariableValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableValue)(nil)).Elem()
}

func (o VariableValueOutput) ToVariableValueOutput() VariableValueOutput {
	return o
}

func (o VariableValueOutput) ToVariableValueOutputWithContext(ctx context.Context) VariableValueOutput {
	return o
}

func (o VariableValueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VariableValue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VariableValueOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VariableValue) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VariableValueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VariableValue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VariableValueOutput) Values() PolicyVariableValueColumnValueResponseArrayOutput {
	return o.ApplyT(func(v *VariableValue) PolicyVariableValueColumnValueResponseArrayOutput { return v.Values }).(PolicyVariableValueColumnValueResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VariableValueOutput{})
}
