


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ModelContainer struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput          `pulumi:"name"`
	Properties ModelContainerResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput     `pulumi:"systemData"`
	Type       pulumi.StringOutput          `pulumi:"type"`
}


func NewModelContainer(ctx *pulumi.Context,
	name string, args *ModelContainerArgs, opts ...pulumi.ResourceOption) (*ModelContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:ModelContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource ModelContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices:ModelContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetModelContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelContainerState, opts ...pulumi.ResourceOption) (*ModelContainer, error) {
	var resource ModelContainer
	err := ctx.ReadResource("azure-native:machinelearningservices:ModelContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type modelContainerState struct {
}

type ModelContainerState struct {
}

func (ModelContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelContainerState)(nil)).Elem()
}

type modelContainerArgs struct {
	Name              *string            `pulumi:"name"`
	Properties        ModelContainerType `pulumi:"properties"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	WorkspaceName     string             `pulumi:"workspaceName"`
}


type ModelContainerArgs struct {
	Name              pulumi.StringPtrInput
	Properties        ModelContainerTypeInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (ModelContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelContainerArgs)(nil)).Elem()
}

type ModelContainerInput interface {
	pulumi.Input

	ToModelContainerOutput() ModelContainerOutput
	ToModelContainerOutputWithContext(ctx context.Context) ModelContainerOutput
}

func (*ModelContainer) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelContainer)(nil))
}

func (i *ModelContainer) ToModelContainerOutput() ModelContainerOutput {
	return i.ToModelContainerOutputWithContext(context.Background())
}

func (i *ModelContainer) ToModelContainerOutputWithContext(ctx context.Context) ModelContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelContainerOutput)
}

type ModelContainerOutput struct{ *pulumi.OutputState }

func (ModelContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelContainer)(nil))
}

func (o ModelContainerOutput) ToModelContainerOutput() ModelContainerOutput {
	return o
}

func (o ModelContainerOutput) ToModelContainerOutputWithContext(ctx context.Context) ModelContainerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ModelContainerOutput{})
}
