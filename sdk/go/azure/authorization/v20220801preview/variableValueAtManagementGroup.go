


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VariableValueAtManagementGroup struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                               `pulumi:"name"`
	SystemData SystemDataResponseOutput                          `pulumi:"systemData"`
	Type       pulumi.StringOutput                               `pulumi:"type"`
	Values     PolicyVariableValueColumnValueResponseArrayOutput `pulumi:"values"`
}


func NewVariableValueAtManagementGroup(ctx *pulumi.Context,
	name string, args *VariableValueAtManagementGroupArgs, opts ...pulumi.ResourceOption) (*VariableValueAtManagementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	if args.Values == nil {
		return nil, errors.New("invalid value for required argument 'Values'")
	}
	if args.VariableName == nil {
		return nil, errors.New("invalid value for required argument 'VariableName'")
	}
	var resource VariableValueAtManagementGroup
	err := ctx.RegisterResource("azure-native:authorization/v20220801preview:VariableValueAtManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVariableValueAtManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableValueAtManagementGroupState, opts ...pulumi.ResourceOption) (*VariableValueAtManagementGroup, error) {
	var resource VariableValueAtManagementGroup
	err := ctx.ReadResource("azure-native:authorization/v20220801preview:VariableValueAtManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type variableValueAtManagementGroupState struct {
}

type VariableValueAtManagementGroupState struct {
}

func (VariableValueAtManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*variableValueAtManagementGroupState)(nil)).Elem()
}

type variableValueAtManagementGroupArgs struct {
	ManagementGroupId string                           `pulumi:"managementGroupId"`
	Values            []PolicyVariableValueColumnValue `pulumi:"values"`
	VariableName      string                           `pulumi:"variableName"`
	VariableValueName *string                          `pulumi:"variableValueName"`
}


type VariableValueAtManagementGroupArgs struct {
	ManagementGroupId pulumi.StringInput
	Values            PolicyVariableValueColumnValueArrayInput
	VariableName      pulumi.StringInput
	VariableValueName pulumi.StringPtrInput
}

func (VariableValueAtManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variableValueAtManagementGroupArgs)(nil)).Elem()
}

type VariableValueAtManagementGroupInput interface {
	pulumi.Input

	ToVariableValueAtManagementGroupOutput() VariableValueAtManagementGroupOutput
	ToVariableValueAtManagementGroupOutputWithContext(ctx context.Context) VariableValueAtManagementGroupOutput
}

func (*VariableValueAtManagementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableValueAtManagementGroup)(nil)).Elem()
}

func (i *VariableValueAtManagementGroup) ToVariableValueAtManagementGroupOutput() VariableValueAtManagementGroupOutput {
	return i.ToVariableValueAtManagementGroupOutputWithContext(context.Background())
}

func (i *VariableValueAtManagementGroup) ToVariableValueAtManagementGroupOutputWithContext(ctx context.Context) VariableValueAtManagementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableValueAtManagementGroupOutput)
}

type VariableValueAtManagementGroupOutput struct{ *pulumi.OutputState }

func (VariableValueAtManagementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableValueAtManagementGroup)(nil)).Elem()
}

func (o VariableValueAtManagementGroupOutput) ToVariableValueAtManagementGroupOutput() VariableValueAtManagementGroupOutput {
	return o
}

func (o VariableValueAtManagementGroupOutput) ToVariableValueAtManagementGroupOutputWithContext(ctx context.Context) VariableValueAtManagementGroupOutput {
	return o
}

func (o VariableValueAtManagementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VariableValueAtManagementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VariableValueAtManagementGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VariableValueAtManagementGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VariableValueAtManagementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VariableValueAtManagementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VariableValueAtManagementGroupOutput) Values() PolicyVariableValueColumnValueResponseArrayOutput {
	return o.ApplyT(func(v *VariableValueAtManagementGroup) PolicyVariableValueColumnValueResponseArrayOutput {
		return v.Values
	}).(PolicyVariableValueColumnValueResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VariableValueAtManagementGroupOutput{})
}
