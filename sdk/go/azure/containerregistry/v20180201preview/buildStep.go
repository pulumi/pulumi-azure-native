


package v20180201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BuildStep struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput           `pulumi:"name"`
	Properties DockerBuildStepResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput           `pulumi:"type"`
}


func NewBuildStep(ctx *pulumi.Context,
	name string, args *BuildStepArgs, opts ...pulumi.ResourceOption) (*BuildStep, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BuildTaskName == nil {
		return nil, errors.New("invalid value for required argument 'BuildTaskName'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20180201preview:BuildStep"),
		},
	})
	opts = append(opts, aliases)
	var resource BuildStep
	err := ctx.RegisterResource("azure-native:containerregistry/v20180201preview:BuildStep", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBuildStep(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildStepState, opts ...pulumi.ResourceOption) (*BuildStep, error) {
	var resource BuildStep
	err := ctx.ReadResource("azure-native:containerregistry/v20180201preview:BuildStep", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type buildStepState struct {
}

type BuildStepState struct {
}

func (BuildStepState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildStepState)(nil)).Elem()
}

type buildStepArgs struct {
	BuildTaskName     string  `pulumi:"buildTaskName"`
	RegistryName      string  `pulumi:"registryName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	StepName          *string `pulumi:"stepName"`
}


type BuildStepArgs struct {
	BuildTaskName     pulumi.StringInput
	RegistryName      pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	StepName          pulumi.StringPtrInput
}

func (BuildStepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildStepArgs)(nil)).Elem()
}

type BuildStepInput interface {
	pulumi.Input

	ToBuildStepOutput() BuildStepOutput
	ToBuildStepOutputWithContext(ctx context.Context) BuildStepOutput
}

func (*BuildStep) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildStep)(nil))
}

func (i *BuildStep) ToBuildStepOutput() BuildStepOutput {
	return i.ToBuildStepOutputWithContext(context.Background())
}

func (i *BuildStep) ToBuildStepOutputWithContext(ctx context.Context) BuildStepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildStepOutput)
}

type BuildStepOutput struct{ *pulumi.OutputState }

func (BuildStepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildStep)(nil))
}

func (o BuildStepOutput) ToBuildStepOutput() BuildStepOutput {
	return o
}

func (o BuildStepOutput) ToBuildStepOutputWithContext(ctx context.Context) BuildStepOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BuildStepInput)(nil)).Elem(), &BuildStep{})
	pulumi.RegisterOutputType(BuildStepOutput{})
}
