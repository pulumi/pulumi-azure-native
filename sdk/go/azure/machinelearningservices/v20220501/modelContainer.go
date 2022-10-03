


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ModelContainer struct {
	pulumi.CustomResourceState

	ModelContainerProperties ModelContainerResponseOutput `pulumi:"modelContainerProperties"`
	Name                     pulumi.StringOutput          `pulumi:"name"`
	SystemData               SystemDataResponseOutput     `pulumi:"systemData"`
	Type                     pulumi.StringOutput          `pulumi:"type"`
}


func NewModelContainer(ctx *pulumi.Context,
	name string, args *ModelContainerArgs, opts ...pulumi.ResourceOption) (*ModelContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModelContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'ModelContainerProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ModelContainerProperties = args.ModelContainerProperties.ToModelContainerTypeOutput().ApplyT(func(v ModelContainerType) ModelContainerType { return *v.Defaults() }).(ModelContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:ModelContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:ModelContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource ModelContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220501:ModelContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetModelContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelContainerState, opts ...pulumi.ResourceOption) (*ModelContainer, error) {
	var resource ModelContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220501:ModelContainer", name, id, state, &resource, opts...)
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
	ModelContainerProperties ModelContainerType `pulumi:"modelContainerProperties"`
	Name                     *string            `pulumi:"name"`
	ResourceGroupName        string             `pulumi:"resourceGroupName"`
	WorkspaceName            string             `pulumi:"workspaceName"`
}


type ModelContainerArgs struct {
	ModelContainerProperties ModelContainerTypeInput
	Name                     pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	WorkspaceName            pulumi.StringInput
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
	return reflect.TypeOf((**ModelContainer)(nil)).Elem()
}

func (i *ModelContainer) ToModelContainerOutput() ModelContainerOutput {
	return i.ToModelContainerOutputWithContext(context.Background())
}

func (i *ModelContainer) ToModelContainerOutputWithContext(ctx context.Context) ModelContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelContainerOutput)
}

type ModelContainerOutput struct{ *pulumi.OutputState }

func (ModelContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelContainer)(nil)).Elem()
}

func (o ModelContainerOutput) ToModelContainerOutput() ModelContainerOutput {
	return o
}

func (o ModelContainerOutput) ToModelContainerOutputWithContext(ctx context.Context) ModelContainerOutput {
	return o
}

func (o ModelContainerOutput) ModelContainerProperties() ModelContainerResponseOutput {
	return o.ApplyT(func(v *ModelContainer) ModelContainerResponseOutput { return v.ModelContainerProperties }).(ModelContainerResponseOutput)
}

func (o ModelContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ModelContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ModelContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ModelContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ModelContainerOutput{})
}
