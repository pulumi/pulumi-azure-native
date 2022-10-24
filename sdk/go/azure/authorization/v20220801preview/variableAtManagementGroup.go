


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VariableAtManagementGroup struct {
	pulumi.CustomResourceState

	Columns    PolicyVariableColumnResponseArrayOutput `pulumi:"columns"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewVariableAtManagementGroup(ctx *pulumi.Context,
	name string, args *VariableAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*VariableAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Columns == nil {
		return nil, errors.New("invalid value for required argument 'Columns'")
	}
	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	var resource VariableAtManagementGroup
	err := ctx.RegisterResource("azure-native:authorization/v20220801preview:VariableAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVariableAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableAtManagementGroupState, opts ...pulumi.ResourceOption) (*VariableAtManagementGroup, error) {
	var resource VariableAtManagementGroup
	err := ctx.ReadResource("azure-native:authorization/v20220801preview:VariableAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type variableAtManagementGroupState struct {
}

type VariableAtManagementGroupState struct {
}

func (VariableAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*variableAtManagementGroupState)(nil)).Elem()
}

type variableAtManagementGroupArgs struct {
	Columns           []PolicyVariableColumn `pulumi:"columns"`
	ManagementGroupId string                 `pulumi:"managementGroupId"`
	VariableName      *string                `pulumi:"variableName"`
}


type VariableAtManagementGroupArgs struct {
	Columns           PolicyVariableColumnArrayInput
	ManagementGroupId pulumi.StringInput
	VariableName      pulumi.StringPtrInput
}

func (VariableAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variableAtManagementGroupArgs)(nil)).Elem()
}

type VariableAtManagementGroupInput interface {
	pulumi.Input

	ToVariableAtManagementGroupOutput() VariableAtManagementGroupOutput
	ToVariableAtManagementGroupOutputWithContext(ctx context.Context) VariableAtManagementGroupOutput
}

func (*VariableAtManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableAtManagementGroup)(nil)).Elem()
}

func (i *VariableAtManagementGroup) ToVariableAtManagementGroupOutput() VariableAtManagementGroupOutput {
	return i.ToVariableAtManagementGroupOutputWithContext(context.Background())
}

func (i *VariableAtManagementGroup) ToVariableAtManagementGroupOutputWithContext(ctx context.Context) VariableAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableAtManagementGroupOutput)
}

type VariableAtManagementGroupOutput struct{ *pulumi.OutputState }

func (VariableAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableAtManagementGroup)(nil)).Elem()
}

func (o VariableAtManagementGroupOutput) ToVariableAtManagementGroupOutput() VariableAtManagementGroupOutput {
	return o
}

func (o VariableAtManagementGroupOutput) ToVariableAtManagementGroupOutputWithContext(ctx context.Context) VariableAtManagementGroupOutput {
	return o
}

func (o VariableAtManagementGroupOutput) Columns() PolicyVariableColumnResponseArrayOutput {
	return o.ApplyT(func(v *VariableAtManagementGroup) PolicyVariableColumnResponseArrayOutput { return v.Columns }).(PolicyVariableColumnResponseArrayOutput)
}

func (o VariableAtManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VariableAtManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VariableAtManagementGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VariableAtManagementGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VariableAtManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VariableAtManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VariableAtManagementGroupOutput{})
}
