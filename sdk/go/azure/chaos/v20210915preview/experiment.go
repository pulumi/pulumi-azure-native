


package v20210915preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Experiment struct {
	pulumi.CustomResourceState

	Identity   ResourceIdentityResponsePtrOutput  `pulumi:"identity"`
	Location   pulumi.StringOutput                `pulumi:"location"`
	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties ExperimentPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput           `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput             `pulumi:"tags"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewExperiment(ctx *pulumi.Context,
	name string, args *ExperimentArgs, opts ...pulumi.ResourceOption) (*Experiment, error) {
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
			Type: pulumi.String("azure-native:chaos:Experiment"),
		},
	})
	opts = append(opts, aliases)
	var resource Experiment
	err := ctx.RegisterResource("azure-native:chaos/v20210915preview:Experiment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExperiment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExperimentState, opts ...pulumi.ResourceOption) (*Experiment, error) {
	var resource Experiment
	err := ctx.ReadResource("azure-native:chaos/v20210915preview:Experiment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type experimentState struct {
}

type ExperimentState struct {
}

func (ExperimentState) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentState)(nil)).Elem()
}

type experimentArgs struct {
	ExperimentName    *string              `pulumi:"experimentName"`
	Identity          *ResourceIdentity    `pulumi:"identity"`
	Location          *string              `pulumi:"location"`
	Properties        ExperimentProperties `pulumi:"properties"`
	ResourceGroupName string               `pulumi:"resourceGroupName"`
	Tags              map[string]string    `pulumi:"tags"`
}


type ExperimentArgs struct {
	ExperimentName    pulumi.StringPtrInput
	Identity          ResourceIdentityPtrInput
	Location          pulumi.StringPtrInput
	Properties        ExperimentPropertiesInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (ExperimentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentArgs)(nil)).Elem()
}

type ExperimentInput interface {
	pulumi.Input

	ToExperimentOutput() ExperimentOutput
	ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput
}

func (*Experiment) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiment)(nil)).Elem()
}

func (i *Experiment) ToExperimentOutput() ExperimentOutput {
	return i.ToExperimentOutputWithContext(context.Background())
}

func (i *Experiment) ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentOutput)
}

type ExperimentOutput struct{ *pulumi.OutputState }

func (ExperimentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiment)(nil)).Elem()
}

func (o ExperimentOutput) ToExperimentOutput() ExperimentOutput {
	return o
}

func (o ExperimentOutput) ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExperimentOutput{})
}
