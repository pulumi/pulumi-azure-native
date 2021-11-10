


package portal

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Console struct {
	pulumi.CustomResourceState

	Properties ConsolePropertiesResponseOutput `pulumi:"properties"`
}


func NewConsole(ctx *pulumi.Context,
	name string, args *ConsoleArgs, opts ...pulumi.ResourceOption) (*Console, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:portal/v20181001:Console"),
		},
	})
	opts = append(opts, aliases)
	var resource Console
	err := ctx.RegisterResource("azure-native:portal:Console", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConsole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsoleState, opts ...pulumi.ResourceOption) (*Console, error) {
	var resource Console
	err := ctx.ReadResource("azure-native:portal:Console", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type consoleState struct {
}

type ConsoleState struct {
}

func (ConsoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*consoleState)(nil)).Elem()
}

type consoleArgs struct {
	ConsoleName *string                 `pulumi:"consoleName"`
	Properties  ConsoleCreateProperties `pulumi:"properties"`
}


type ConsoleArgs struct {
	ConsoleName pulumi.StringPtrInput
	Properties  ConsoleCreatePropertiesInput
}

func (ConsoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consoleArgs)(nil)).Elem()
}

type ConsoleInput interface {
	pulumi.Input

	ToConsoleOutput() ConsoleOutput
	ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput
}

func (*Console) ElementType() reflect.Type {
	return reflect.TypeOf((*Console)(nil))
}

func (i *Console) ToConsoleOutput() ConsoleOutput {
	return i.ToConsoleOutputWithContext(context.Background())
}

func (i *Console) ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleOutput)
}

type ConsoleOutput struct{ *pulumi.OutputState }

func (ConsoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Console)(nil))
}

func (o ConsoleOutput) ToConsoleOutput() ConsoleOutput {
	return o
}

func (o ConsoleOutput) ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConsoleOutput{})
}
