


package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Function struct {
	pulumi.CustomResourceState

	Name       pulumi.StringPtrOutput                 `pulumi:"name"`
	Properties ScalarFunctionPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
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
			Type: pulumi.String("azure-native:streamanalytics/v20160301:Function"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20170401preview:Function"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20200301:Function"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20211001preview:Function"),
		},
	})
	opts = append(opts, aliases)
	var resource Function
	err := ctx.RegisterResource("azure-native:streamanalytics:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("azure-native:streamanalytics:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type functionState struct {
}

type FunctionState struct {
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	FunctionName      *string                   `pulumi:"functionName"`
	JobName           string                    `pulumi:"jobName"`
	Name              *string                   `pulumi:"name"`
	Properties        *ScalarFunctionProperties `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
}


type FunctionArgs struct {
	FunctionName      pulumi.StringPtrInput
	JobName           pulumi.StringInput
	Name              pulumi.StringPtrInput
	Properties        ScalarFunctionPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

type FunctionInput interface {
	pulumi.Input

	ToFunctionOutput() FunctionOutput
	ToFunctionOutputWithContext(ctx context.Context) FunctionOutput
}

func (*Function) ElementType() reflect.Type {
	return reflect.TypeOf((*Function)(nil))
}

func (i *Function) ToFunctionOutput() FunctionOutput {
	return i.ToFunctionOutputWithContext(context.Background())
}

func (i *Function) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutput)
}

type FunctionOutput struct{ *pulumi.OutputState }

func (FunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Function)(nil))
}

func (o FunctionOutput) ToFunctionOutput() FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FunctionOutput{})
}
