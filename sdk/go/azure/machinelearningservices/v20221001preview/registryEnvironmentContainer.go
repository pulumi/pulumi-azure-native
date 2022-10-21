


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistryEnvironmentContainer struct {
	pulumi.CustomResourceState

	EnvironmentContainerProperties EnvironmentContainerResponseOutput `pulumi:"environmentContainerProperties"`
	Name                           pulumi.StringOutput                `pulumi:"name"`
	SystemData                     SystemDataResponseOutput           `pulumi:"systemData"`
	Type                           pulumi.StringOutput                `pulumi:"type"`
}


func NewRegistryEnvironmentContainer(ctx *pulumi.Context,
	name string, args *RegistryEnvironmentContainerArgs, opts ...pulumi.ResourceOption) (*RegistryEnvironmentContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentContainerProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.EnvironmentContainerProperties = args.EnvironmentContainerProperties.ToEnvironmentContainerTypeOutput().ApplyT(func(v EnvironmentContainerType) EnvironmentContainerType { return *v.Defaults() }).(EnvironmentContainerTypeOutput)
	var resource RegistryEnvironmentContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001preview:RegistryEnvironmentContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistryEnvironmentContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryEnvironmentContainerState, opts ...pulumi.ResourceOption) (*RegistryEnvironmentContainer, error) {
	var resource RegistryEnvironmentContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001preview:RegistryEnvironmentContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryEnvironmentContainerState struct {
}

type RegistryEnvironmentContainerState struct {
}

func (RegistryEnvironmentContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnvironmentContainerState)(nil)).Elem()
}

type registryEnvironmentContainerArgs struct {
	EnvironmentContainerProperties EnvironmentContainerType `pulumi:"environmentContainerProperties"`
	EnvironmentName                *string                  `pulumi:"environmentName"`
	RegistryName                   string                   `pulumi:"registryName"`
	ResourceGroupName              string                   `pulumi:"resourceGroupName"`
}


type RegistryEnvironmentContainerArgs struct {
	EnvironmentContainerProperties EnvironmentContainerTypeInput
	EnvironmentName                pulumi.StringPtrInput
	RegistryName                   pulumi.StringInput
	ResourceGroupName              pulumi.StringInput
}

func (RegistryEnvironmentContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnvironmentContainerArgs)(nil)).Elem()
}

type RegistryEnvironmentContainerInput interface {
	pulumi.Input

	ToRegistryEnvironmentContainerOutput() RegistryEnvironmentContainerOutput
	ToRegistryEnvironmentContainerOutputWithContext(ctx context.Context) RegistryEnvironmentContainerOutput
}

func (*RegistryEnvironmentContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnvironmentContainer)(nil)).Elem()
}

func (i *RegistryEnvironmentContainer) ToRegistryEnvironmentContainerOutput() RegistryEnvironmentContainerOutput {
	return i.ToRegistryEnvironmentContainerOutputWithContext(context.Background())
}

func (i *RegistryEnvironmentContainer) ToRegistryEnvironmentContainerOutputWithContext(ctx context.Context) RegistryEnvironmentContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnvironmentContainerOutput)
}

type RegistryEnvironmentContainerOutput struct{ *pulumi.OutputState }

func (RegistryEnvironmentContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnvironmentContainer)(nil)).Elem()
}

func (o RegistryEnvironmentContainerOutput) ToRegistryEnvironmentContainerOutput() RegistryEnvironmentContainerOutput {
	return o
}

func (o RegistryEnvironmentContainerOutput) ToRegistryEnvironmentContainerOutputWithContext(ctx context.Context) RegistryEnvironmentContainerOutput {
	return o
}

func (o RegistryEnvironmentContainerOutput) EnvironmentContainerProperties() EnvironmentContainerResponseOutput {
	return o.ApplyT(func(v *RegistryEnvironmentContainer) EnvironmentContainerResponseOutput {
		return v.EnvironmentContainerProperties
	}).(EnvironmentContainerResponseOutput)
}

func (o RegistryEnvironmentContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnvironmentContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegistryEnvironmentContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryEnvironmentContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RegistryEnvironmentContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnvironmentContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryEnvironmentContainerOutput{})
}
