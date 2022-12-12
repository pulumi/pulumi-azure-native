


package v20151031

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Variable struct {
	pulumi.CustomResourceState

	CreationTime     pulumi.StringPtrOutput `pulumi:"creationTime"`
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	IsEncrypted      pulumi.BoolPtrOutput   `pulumi:"isEncrypted"`
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Type             pulumi.StringOutput    `pulumi:"type"`
	Value            pulumi.StringPtrOutput `pulumi:"value"`
}


func NewVariable(ctx *pulumi.Context,
	name string, args *VariableArgs, opts ...pulumi.ResourceOption) (*Variable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Variable"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Variable"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:Variable"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:Variable"),
		},
	})
	opts = append(opts, aliases)
	var resource Variable
	err := ctx.RegisterResource("azure-native:automation/v20151031:Variable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableState, opts ...pulumi.ResourceOption) (*Variable, error) {
	var resource Variable
	err := ctx.ReadResource("azure-native:automation/v20151031:Variable", name, id, state, &resource, opts...)
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
	AutomationAccountName string  `pulumi:"automationAccountName"`
	Description           *string `pulumi:"description"`
	IsEncrypted           *bool   `pulumi:"isEncrypted"`
	Name                  string  `pulumi:"name"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	Value                 *string `pulumi:"value"`
	VariableName          *string `pulumi:"variableName"`
}


type VariableArgs struct {
	AutomationAccountName pulumi.StringInput
	Description           pulumi.StringPtrInput
	IsEncrypted           pulumi.BoolPtrInput
	Name                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Value                 pulumi.StringPtrInput
	VariableName          pulumi.StringPtrInput
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

func (o VariableOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o VariableOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VariableOutput) IsEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.BoolPtrOutput { return v.IsEncrypted }).(pulumi.BoolPtrOutput)
}

func (o VariableOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o VariableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VariableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VariableOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VariableOutput{})
}
