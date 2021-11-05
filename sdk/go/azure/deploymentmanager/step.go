


package deploymentmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Step struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput    `pulumi:"location"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Properties pulumi.AnyOutput       `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewStep(ctx *pulumi.Context,
	name string, args *StepArgs, opts ...pulumi.ResourceOption) (*Step, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager:Step"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20180901preview:Step"),
		},
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager/v20180901preview:Step"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20191101preview:Step"),
		},
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager/v20191101preview:Step"),
		},
	})
	opts = append(opts, aliases)
	var resource Step
	err := ctx.RegisterResource("azure-native:deploymentmanager:Step", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStep(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StepState, opts ...pulumi.ResourceOption) (*Step, error) {
	var resource Step
	err := ctx.ReadResource("azure-native:deploymentmanager:Step", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type stepState struct {
}

type StepState struct {
}

func (StepState) ElementType() reflect.Type {
	return reflect.TypeOf((*stepState)(nil)).Elem()
}

type stepArgs struct {
	Location          *string           `pulumi:"location"`
	Properties        interface{}       `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	StepName          *string           `pulumi:"stepName"`
	Tags              map[string]string `pulumi:"tags"`
}


type StepArgs struct {
	Location          pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
	StepName          pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (StepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stepArgs)(nil)).Elem()
}

type StepInput interface {
	pulumi.Input

	ToStepOutput() StepOutput
	ToStepOutputWithContext(ctx context.Context) StepOutput
}

func (*Step) ElementType() reflect.Type {
	return reflect.TypeOf((*Step)(nil))
}

func (i *Step) ToStepOutput() StepOutput {
	return i.ToStepOutputWithContext(context.Background())
}

func (i *Step) ToStepOutputWithContext(ctx context.Context) StepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepOutput)
}

type StepOutput struct{ *pulumi.OutputState }

func (StepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Step)(nil))
}

func (o StepOutput) ToStepOutput() StepOutput {
	return o
}

func (o StepOutput) ToStepOutputWithContext(ctx context.Context) StepOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StepOutput{})
}
