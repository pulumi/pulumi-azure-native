


package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Experiment struct {
	pulumi.CustomResourceState

	Description   pulumi.StringPtrOutput              `pulumi:"description"`
	EnabledState  pulumi.StringPtrOutput              `pulumi:"enabledState"`
	EndpointA     ExperimentEndpointResponsePtrOutput `pulumi:"endpointA"`
	EndpointB     ExperimentEndpointResponsePtrOutput `pulumi:"endpointB"`
	Location      pulumi.StringPtrOutput              `pulumi:"location"`
	Name          pulumi.StringOutput                 `pulumi:"name"`
	ResourceState pulumi.StringOutput                 `pulumi:"resourceState"`
	ScriptFileUri pulumi.StringOutput                 `pulumi:"scriptFileUri"`
	Status        pulumi.StringOutput                 `pulumi:"status"`
	Tags          pulumi.StringMapOutput              `pulumi:"tags"`
	Type          pulumi.StringOutput                 `pulumi:"type"`
}


func NewExperiment(ctx *pulumi.Context,
	name string, args *ExperimentArgs, opts ...pulumi.ResourceOption) (*Experiment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:Experiment"),
		},
	})
	opts = append(opts, aliases)
	var resource Experiment
	err := ctx.RegisterResource("azure-native:network/v20191101:Experiment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExperiment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExperimentState, opts ...pulumi.ResourceOption) (*Experiment, error) {
	var resource Experiment
	err := ctx.ReadResource("azure-native:network/v20191101:Experiment", name, id, state, &resource, opts...)
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
	Description       *string             `pulumi:"description"`
	EnabledState      *string             `pulumi:"enabledState"`
	EndpointA         *ExperimentEndpoint `pulumi:"endpointA"`
	EndpointB         *ExperimentEndpoint `pulumi:"endpointB"`
	ExperimentName    *string             `pulumi:"experimentName"`
	Location          *string             `pulumi:"location"`
	ProfileName       string              `pulumi:"profileName"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Tags              map[string]string   `pulumi:"tags"`
}


type ExperimentArgs struct {
	Description       pulumi.StringPtrInput
	EnabledState      pulumi.StringPtrInput
	EndpointA         ExperimentEndpointPtrInput
	EndpointB         ExperimentEndpointPtrInput
	ExperimentName    pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ProfileName       pulumi.StringInput
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

func (o ExperimentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ExperimentOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringPtrOutput { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o ExperimentOutput) EndpointA() ExperimentEndpointResponsePtrOutput {
	return o.ApplyT(func(v *Experiment) ExperimentEndpointResponsePtrOutput { return v.EndpointA }).(ExperimentEndpointResponsePtrOutput)
}

func (o ExperimentOutput) EndpointB() ExperimentEndpointResponsePtrOutput {
	return o.ApplyT(func(v *Experiment) ExperimentEndpointResponsePtrOutput { return v.EndpointB }).(ExperimentEndpointResponsePtrOutput)
}

func (o ExperimentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ExperimentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExperimentOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o ExperimentOutput) ScriptFileUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.ScriptFileUri }).(pulumi.StringOutput)
}

func (o ExperimentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ExperimentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ExperimentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExperimentOutput{})
}
