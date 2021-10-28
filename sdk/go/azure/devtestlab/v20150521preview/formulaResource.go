


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FormulaResource struct {
	pulumi.CustomResourceState

	Author            pulumi.StringPtrOutput                   `pulumi:"author"`
	CreationDate      pulumi.StringPtrOutput                   `pulumi:"creationDate"`
	Description       pulumi.StringPtrOutput                   `pulumi:"description"`
	FormulaContent    LabVirtualMachineResponsePtrOutput       `pulumi:"formulaContent"`
	Location          pulumi.StringPtrOutput                   `pulumi:"location"`
	Name              pulumi.StringPtrOutput                   `pulumi:"name"`
	OsType            pulumi.StringPtrOutput                   `pulumi:"osType"`
	ProvisioningState pulumi.StringPtrOutput                   `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput                   `pulumi:"tags"`
	Type              pulumi.StringPtrOutput                   `pulumi:"type"`
	Vm                FormulaPropertiesFromVmResponsePtrOutput `pulumi:"vm"`
}


func NewFormulaResource(ctx *pulumi.Context,
	name string, args *FormulaResourceArgs, opts ...pulumi.ResourceOption) (*FormulaResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20150521preview:FormulaResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab:FormulaResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab:FormulaResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:FormulaResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20160515:FormulaResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:FormulaResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:FormulaResource"),
		},
	})
	opts = append(opts, aliases)
	var resource FormulaResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:FormulaResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFormulaResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FormulaResourceState, opts ...pulumi.ResourceOption) (*FormulaResource, error) {
	var resource FormulaResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:FormulaResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type formulaResourceState struct {
}

type FormulaResourceState struct {
}

func (FormulaResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*formulaResourceState)(nil)).Elem()
}

type formulaResourceArgs struct {
	Author            *string                  `pulumi:"author"`
	CreationDate      *string                  `pulumi:"creationDate"`
	Description       *string                  `pulumi:"description"`
	FormulaContent    *LabVirtualMachine       `pulumi:"formulaContent"`
	Id                *string                  `pulumi:"id"`
	LabName           string                   `pulumi:"labName"`
	Location          *string                  `pulumi:"location"`
	Name              *string                  `pulumi:"name"`
	OsType            *string                  `pulumi:"osType"`
	ProvisioningState *string                  `pulumi:"provisioningState"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	Tags              map[string]string        `pulumi:"tags"`
	Type              *string                  `pulumi:"type"`
	Vm                *FormulaPropertiesFromVm `pulumi:"vm"`
}


type FormulaResourceArgs struct {
	Author            pulumi.StringPtrInput
	CreationDate      pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	FormulaContent    LabVirtualMachinePtrInput
	Id                pulumi.StringPtrInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	OsType            pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
	Vm                FormulaPropertiesFromVmPtrInput
}

func (FormulaResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*formulaResourceArgs)(nil)).Elem()
}

type FormulaResourceInput interface {
	pulumi.Input

	ToFormulaResourceOutput() FormulaResourceOutput
	ToFormulaResourceOutputWithContext(ctx context.Context) FormulaResourceOutput
}

func (*FormulaResource) ElementType() reflect.Type {
	return reflect.TypeOf((*FormulaResource)(nil))
}

func (i *FormulaResource) ToFormulaResourceOutput() FormulaResourceOutput {
	return i.ToFormulaResourceOutputWithContext(context.Background())
}

func (i *FormulaResource) ToFormulaResourceOutputWithContext(ctx context.Context) FormulaResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaResourceOutput)
}

type FormulaResourceOutput struct{ *pulumi.OutputState }

func (FormulaResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FormulaResource)(nil))
}

func (o FormulaResourceOutput) ToFormulaResourceOutput() FormulaResourceOutput {
	return o
}

func (o FormulaResourceOutput) ToFormulaResourceOutputWithContext(ctx context.Context) FormulaResourceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FormulaResourceInput)(nil)).Elem(), &FormulaResource{})
	pulumi.RegisterOutputType(FormulaResourceOutput{})
}
