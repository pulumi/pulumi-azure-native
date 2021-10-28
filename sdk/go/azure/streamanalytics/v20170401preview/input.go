


package v20170401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Input struct {
	pulumi.CustomResourceState

	Name       pulumi.StringPtrOutput `pulumi:"name"`
	Properties pulumi.AnyOutput       `pulumi:"properties"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewInput(ctx *pulumi.Context,
	name string, args *InputArgs, opts ...pulumi.ResourceOption) (*Input, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobName == nil {
		return nil, errors.New("invalid value for required argument 'JobName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:streamanalytics/v20170401preview:Input"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics:Input"),
		},
		{
			Type: pulumi.String("azure-nextgen:streamanalytics:Input"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20160301:Input"),
		},
		{
			Type: pulumi.String("azure-nextgen:streamanalytics/v20160301:Input"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20200301:Input"),
		},
		{
			Type: pulumi.String("azure-nextgen:streamanalytics/v20200301:Input"),
		},
	})
	opts = append(opts, aliases)
	var resource Input
	err := ctx.RegisterResource("azure-native:streamanalytics/v20170401preview:Input", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetInput(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InputState, opts ...pulumi.ResourceOption) (*Input, error) {
	var resource Input
	err := ctx.ReadResource("azure-native:streamanalytics/v20170401preview:Input", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type inputState struct {
}

type InputState struct {
}

func (InputState) ElementType() reflect.Type {
	return reflect.TypeOf((*inputState)(nil)).Elem()
}

type inputArgs struct {
	InputName         *string     `pulumi:"inputName"`
	JobName           string      `pulumi:"jobName"`
	Name              *string     `pulumi:"name"`
	Properties        interface{} `pulumi:"properties"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
}


type InputArgs struct {
	InputName         pulumi.StringPtrInput
	JobName           pulumi.StringInput
	Name              pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
}

func (InputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inputArgs)(nil)).Elem()
}

type InputInput interface {
	pulumi.Input

	ToInputOutput() InputOutput
	ToInputOutputWithContext(ctx context.Context) InputOutput
}

func (*Input) ElementType() reflect.Type {
	return reflect.TypeOf((*Input)(nil))
}

func (i *Input) ToInputOutput() InputOutput {
	return i.ToInputOutputWithContext(context.Background())
}

func (i *Input) ToInputOutputWithContext(ctx context.Context) InputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputOutput)
}

type InputOutput struct{ *pulumi.OutputState }

func (InputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Input)(nil))
}

func (o InputOutput) ToInputOutput() InputOutput {
	return o
}

func (o InputOutput) ToInputOutputWithContext(ctx context.Context) InputOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InputInput)(nil)).Elem(), &Input{})
	pulumi.RegisterOutputType(InputOutput{})
}
