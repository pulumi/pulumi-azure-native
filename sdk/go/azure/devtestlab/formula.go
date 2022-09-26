


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Formula struct {
	pulumi.CustomResourceState

	Author            pulumi.StringOutput                                 `pulumi:"author"`
	CreationDate      pulumi.StringOutput                                 `pulumi:"creationDate"`
	Description       pulumi.StringPtrOutput                              `pulumi:"description"`
	FormulaContent    LabVirtualMachineCreationParameterResponsePtrOutput `pulumi:"formulaContent"`
	Location          pulumi.StringPtrOutput                              `pulumi:"location"`
	Name              pulumi.StringOutput                                 `pulumi:"name"`
	OsType            pulumi.StringPtrOutput                              `pulumi:"osType"`
	ProvisioningState pulumi.StringOutput                                 `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput                              `pulumi:"tags"`
	Type              pulumi.StringOutput                                 `pulumi:"type"`
	UniqueIdentifier  pulumi.StringOutput                                 `pulumi:"uniqueIdentifier"`
	Vm                FormulaPropertiesFromVmResponsePtrOutput            `pulumi:"vm"`
}


func NewFormula(ctx *pulumi.Context,
	name string, args *FormulaArgs, opts ...pulumi.ResourceOption) (*Formula, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.FormulaContent != nil {
		args.FormulaContent = args.FormulaContent.ToLabVirtualMachineCreationParameterPtrOutput().ApplyT(func(v *LabVirtualMachineCreationParameter) *LabVirtualMachineCreationParameter { return v.Defaults() }).(LabVirtualMachineCreationParameterPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:Formula"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:Formula"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:Formula"),
		},
	})
	opts = append(opts, aliases)
	var resource Formula
	err := ctx.RegisterResource("azure-native:devtestlab:Formula", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFormula(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FormulaState, opts ...pulumi.ResourceOption) (*Formula, error) {
	var resource Formula
	err := ctx.ReadResource("azure-native:devtestlab:Formula", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type formulaState struct {
}

type FormulaState struct {
}

func (FormulaState) ElementType() reflect.Type {
	return reflect.TypeOf((*formulaState)(nil)).Elem()
}

type formulaArgs struct {
	Description       *string                             `pulumi:"description"`
	FormulaContent    *LabVirtualMachineCreationParameter `pulumi:"formulaContent"`
	LabName           string                              `pulumi:"labName"`
	Location          *string                             `pulumi:"location"`
	Name              *string                             `pulumi:"name"`
	OsType            *string                             `pulumi:"osType"`
	ResourceGroupName string                              `pulumi:"resourceGroupName"`
	Tags              map[string]string                   `pulumi:"tags"`
	Vm                *FormulaPropertiesFromVm            `pulumi:"vm"`
}


type FormulaArgs struct {
	Description       pulumi.StringPtrInput
	FormulaContent    LabVirtualMachineCreationParameterPtrInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	OsType            pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Vm                FormulaPropertiesFromVmPtrInput
}

func (FormulaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*formulaArgs)(nil)).Elem()
}

type FormulaInput interface {
	pulumi.Input

	ToFormulaOutput() FormulaOutput
	ToFormulaOutputWithContext(ctx context.Context) FormulaOutput
}

func (*Formula) ElementType() reflect.Type {
	return reflect.TypeOf((**Formula)(nil)).Elem()
}

func (i *Formula) ToFormulaOutput() FormulaOutput {
	return i.ToFormulaOutputWithContext(context.Background())
}

func (i *Formula) ToFormulaOutputWithContext(ctx context.Context) FormulaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaOutput)
}

type FormulaOutput struct{ *pulumi.OutputState }

func (FormulaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Formula)(nil)).Elem()
}

func (o FormulaOutput) ToFormulaOutput() FormulaOutput {
	return o
}

func (o FormulaOutput) ToFormulaOutputWithContext(ctx context.Context) FormulaOutput {
	return o
}

func (o FormulaOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.Author }).(pulumi.StringOutput)
}

func (o FormulaOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o FormulaOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o FormulaOutput) FormulaContent() LabVirtualMachineCreationParameterResponsePtrOutput {
	return o.ApplyT(func(v *Formula) LabVirtualMachineCreationParameterResponsePtrOutput { return v.FormulaContent }).(LabVirtualMachineCreationParameterResponsePtrOutput)
}

func (o FormulaOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o FormulaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FormulaOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o FormulaOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FormulaOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FormulaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o FormulaOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o FormulaOutput) Vm() FormulaPropertiesFromVmResponsePtrOutput {
	return o.ApplyT(func(v *Formula) FormulaPropertiesFromVmResponsePtrOutput { return v.Vm }).(FormulaPropertiesFromVmResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FormulaOutput{})
}
