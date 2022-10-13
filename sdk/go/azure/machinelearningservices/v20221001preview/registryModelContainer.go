


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistryModelContainer struct {
	pulumi.CustomResourceState

	ModelContainerProperties ModelContainerResponseOutput `pulumi:"modelContainerProperties"`
	Name                     pulumi.StringOutput          `pulumi:"name"`
	SystemData               SystemDataResponseOutput     `pulumi:"systemData"`
	Type                     pulumi.StringOutput          `pulumi:"type"`
}


func NewRegistryModelContainer(ctx *pulumi.Context,
	name string, args *RegistryModelContainerArgs, opts ...pulumi.ResourceOption) (*RegistryModelContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModelContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'ModelContainerProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.ModelContainerProperties = args.ModelContainerProperties.ToModelContainerTypeOutput().ApplyT(func(v ModelContainerType) ModelContainerType { return *v.Defaults() }).(ModelContainerTypeOutput)
	var resource RegistryModelContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001preview:RegistryModelContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistryModelContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryModelContainerState, opts ...pulumi.ResourceOption) (*RegistryModelContainer, error) {
	var resource RegistryModelContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001preview:RegistryModelContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryModelContainerState struct {
}

type RegistryModelContainerState struct {
}

func (RegistryModelContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryModelContainerState)(nil)).Elem()
}

type registryModelContainerArgs struct {
	ModelContainerProperties ModelContainerType `pulumi:"modelContainerProperties"`
	ModelName                *string            `pulumi:"modelName"`
	RegistryName             string             `pulumi:"registryName"`
	ResourceGroupName        string             `pulumi:"resourceGroupName"`
}


type RegistryModelContainerArgs struct {
	ModelContainerProperties ModelContainerTypeInput
	ModelName                pulumi.StringPtrInput
	RegistryName             pulumi.StringInput
	ResourceGroupName        pulumi.StringInput
}

func (RegistryModelContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryModelContainerArgs)(nil)).Elem()
}

type RegistryModelContainerInput interface {
	pulumi.Input

	ToRegistryModelContainerOutput() RegistryModelContainerOutput
	ToRegistryModelContainerOutputWithContext(ctx context.Context) RegistryModelContainerOutput
}

func (*RegistryModelContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryModelContainer)(nil)).Elem()
}

func (i *RegistryModelContainer) ToRegistryModelContainerOutput() RegistryModelContainerOutput {
	return i.ToRegistryModelContainerOutputWithContext(context.Background())
}

func (i *RegistryModelContainer) ToRegistryModelContainerOutputWithContext(ctx context.Context) RegistryModelContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryModelContainerOutput)
}

type RegistryModelContainerOutput struct{ *pulumi.OutputState }

func (RegistryModelContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryModelContainer)(nil)).Elem()
}

func (o RegistryModelContainerOutput) ToRegistryModelContainerOutput() RegistryModelContainerOutput {
	return o
}

func (o RegistryModelContainerOutput) ToRegistryModelContainerOutputWithContext(ctx context.Context) RegistryModelContainerOutput {
	return o
}

func (o RegistryModelContainerOutput) ModelContainerProperties() ModelContainerResponseOutput {
	return o.ApplyT(func(v *RegistryModelContainer) ModelContainerResponseOutput { return v.ModelContainerProperties }).(ModelContainerResponseOutput)
}

func (o RegistryModelContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegistryModelContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryModelContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RegistryModelContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryModelContainerOutput{})
}
