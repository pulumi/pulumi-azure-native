


package portal

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConsoleWithLocation struct {
	pulumi.CustomResourceState

	Properties ConsolePropertiesResponseOutput `pulumi:"properties"`
}


func NewConsoleWithLocation(ctx *pulumi.Context,
	name string, args *ConsoleWithLocationArgs, opts ...pulumi.ResourceOption) (*ConsoleWithLocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:portal:ConsoleWithLocation"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20181001:ConsoleWithLocation"),
		},
		{
			Type: pulumi.String("azure-nextgen:portal/v20181001:ConsoleWithLocation"),
		},
	})
	opts = append(opts, aliases)
	var resource ConsoleWithLocation
	err := ctx.RegisterResource("azure-native:portal:ConsoleWithLocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConsoleWithLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsoleWithLocationState, opts ...pulumi.ResourceOption) (*ConsoleWithLocation, error) {
	var resource ConsoleWithLocation
	err := ctx.ReadResource("azure-native:portal:ConsoleWithLocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type consoleWithLocationState struct {
}

type ConsoleWithLocationState struct {
}

func (ConsoleWithLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*consoleWithLocationState)(nil)).Elem()
}

type consoleWithLocationArgs struct {
	ConsoleName *string `pulumi:"consoleName"`
	Location    string  `pulumi:"location"`
}


type ConsoleWithLocationArgs struct {
	ConsoleName pulumi.StringPtrInput
	Location    pulumi.StringInput
}

func (ConsoleWithLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consoleWithLocationArgs)(nil)).Elem()
}

type ConsoleWithLocationInput interface {
	pulumi.Input

	ToConsoleWithLocationOutput() ConsoleWithLocationOutput
	ToConsoleWithLocationOutputWithContext(ctx context.Context) ConsoleWithLocationOutput
}

func (*ConsoleWithLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleWithLocation)(nil))
}

func (i *ConsoleWithLocation) ToConsoleWithLocationOutput() ConsoleWithLocationOutput {
	return i.ToConsoleWithLocationOutputWithContext(context.Background())
}

func (i *ConsoleWithLocation) ToConsoleWithLocationOutputWithContext(ctx context.Context) ConsoleWithLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleWithLocationOutput)
}

type ConsoleWithLocationOutput struct{ *pulumi.OutputState }

func (ConsoleWithLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleWithLocation)(nil))
}

func (o ConsoleWithLocationOutput) ToConsoleWithLocationOutput() ConsoleWithLocationOutput {
	return o
}

func (o ConsoleWithLocationOutput) ToConsoleWithLocationOutputWithContext(ctx context.Context) ConsoleWithLocationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConsoleWithLocationOutput{})
}
