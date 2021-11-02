


package migrate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Solution struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput           `pulumi:"etag"`
	Name       pulumi.StringOutput              `pulumi:"name"`
	Properties SolutionPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput              `pulumi:"type"`
}


func NewSolution(ctx *pulumi.Context,
	name string, args *SolutionArgs, opts ...pulumi.ResourceOption) (*Solution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MigrateProjectName == nil {
		return nil, errors.New("invalid value for required argument 'MigrateProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:migrate:Solution"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20180901preview:Solution"),
		},
		{
			Type: pulumi.String("azure-nextgen:migrate/v20180901preview:Solution"),
		},
	})
	opts = append(opts, aliases)
	var resource Solution
	err := ctx.RegisterResource("azure-native:migrate:Solution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSolution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SolutionState, opts ...pulumi.ResourceOption) (*Solution, error) {
	var resource Solution
	err := ctx.ReadResource("azure-native:migrate:Solution", name, id, state, &resource, opts...)
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
	Etag               *string             `pulumi:"etag"`
	MigrateProjectName string              `pulumi:"migrateProjectName"`
	Properties         *SolutionProperties `pulumi:"properties"`
	ResourceGroupName  string              `pulumi:"resourceGroupName"`
	SolutionName       *string             `pulumi:"solutionName"`
}


type SolutionArgs struct {
	Etag               pulumi.StringPtrInput
	MigrateProjectName pulumi.StringInput
	Properties         SolutionPropertiesPtrInput
	ResourceGroupName  pulumi.StringInput
	SolutionName       pulumi.StringPtrInput
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
	pulumi.RegisterOutputType(SolutionOutput{})
}
