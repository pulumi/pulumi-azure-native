


package operationsmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Solution struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput           `pulumi:"location"`
	Name       pulumi.StringOutput              `pulumi:"name"`
	Plan       SolutionPlanResponsePtrOutput    `pulumi:"plan"`
	Properties SolutionPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput           `pulumi:"tags"`
	Type       pulumi.StringOutput              `pulumi:"type"`
}


func NewSolution(ctx *pulumi.Context,
	name string, args *SolutionArgs, opts ...pulumi.ResourceOption) (*Solution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:operationsmanagement:Solution"),
		},
		{
			Type: pulumi.String("azure-native:operationsmanagement/v20151101preview:Solution"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationsmanagement/v20151101preview:Solution"),
		},
	})
	opts = append(opts, aliases)
	var resource Solution
	err := ctx.RegisterResource("azure-native:operationsmanagement:Solution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSolution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SolutionState, opts ...pulumi.ResourceOption) (*Solution, error) {
	var resource Solution
	err := ctx.ReadResource("azure-native:operationsmanagement:Solution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type solutionState struct {
}

type SolutionState struct {
}

func (SolutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*solutionState)(nil)).Elem()
}

type solutionArgs struct {
	Location          *string             `pulumi:"location"`
	Plan              *SolutionPlan       `pulumi:"plan"`
	Properties        *SolutionProperties `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	SolutionName      *string             `pulumi:"solutionName"`
	Tags              map[string]string   `pulumi:"tags"`
}


type SolutionArgs struct {
	Location          pulumi.StringPtrInput
	Plan              SolutionPlanPtrInput
	Properties        SolutionPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	SolutionName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (SolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*solutionArgs)(nil)).Elem()
}

type SolutionInput interface {
	pulumi.Input

	ToSolutionOutput() SolutionOutput
	ToSolutionOutputWithContext(ctx context.Context) SolutionOutput
}

func (*Solution) ElementType() reflect.Type {
	return reflect.TypeOf((*Solution)(nil))
}

func (i *Solution) ToSolutionOutput() SolutionOutput {
	return i.ToSolutionOutputWithContext(context.Background())
}

func (i *Solution) ToSolutionOutputWithContext(ctx context.Context) SolutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionOutput)
}

type SolutionOutput struct{ *pulumi.OutputState }

func (SolutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Solution)(nil))
}

func (o SolutionOutput) ToSolutionOutput() SolutionOutput {
	return o
}

func (o SolutionOutput) ToSolutionOutputWithContext(ctx context.Context) SolutionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionInput)(nil)).Elem(), &Solution{})
	pulumi.RegisterOutputType(SolutionOutput{})
}
