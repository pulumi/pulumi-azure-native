


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistryComponentContainer struct {
	pulumi.CustomResourceState

	ComponentContainerProperties ComponentContainerResponseOutput `pulumi:"componentContainerProperties"`
	Name                         pulumi.StringOutput              `pulumi:"name"`
	SystemData                   SystemDataResponseOutput         `pulumi:"systemData"`
	Type                         pulumi.StringOutput              `pulumi:"type"`
}


func NewRegistryComponentContainer(ctx *pulumi.Context,
	name string, args *RegistryComponentContainerArgs, opts ...pulumi.ResourceOption) (*RegistryComponentContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComponentContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'ComponentContainerProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.ComponentContainerProperties = args.ComponentContainerProperties.ToComponentContainerTypeOutput().ApplyT(func(v ComponentContainerType) ComponentContainerType { return *v.Defaults() }).(ComponentContainerTypeOutput)
	var resource RegistryComponentContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001preview:RegistryComponentContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistryComponentContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryComponentContainerState, opts ...pulumi.ResourceOption) (*RegistryComponentContainer, error) {
	var resource RegistryComponentContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001preview:RegistryComponentContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryComponentContainerState struct {
}

type RegistryComponentContainerState struct {
}

func (RegistryComponentContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryComponentContainerState)(nil)).Elem()
}

type registryComponentContainerArgs struct {
	ComponentContainerProperties ComponentContainerType `pulumi:"componentContainerProperties"`
	ComponentName                *string                `pulumi:"componentName"`
	RegistryName                 string                 `pulumi:"registryName"`
	ResourceGroupName            string                 `pulumi:"resourceGroupName"`
}


type RegistryComponentContainerArgs struct {
	ComponentContainerProperties ComponentContainerTypeInput
	ComponentName                pulumi.StringPtrInput
	RegistryName                 pulumi.StringInput
	ResourceGroupName            pulumi.StringInput
}

func (RegistryComponentContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryComponentContainerArgs)(nil)).Elem()
}

type RegistryComponentContainerInput interface {
	pulumi.Input

	ToRegistryComponentContainerOutput() RegistryComponentContainerOutput
	ToRegistryComponentContainerOutputWithContext(ctx context.Context) RegistryComponentContainerOutput
}

func (*RegistryComponentContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryComponentContainer)(nil)).Elem()
}

func (i *RegistryComponentContainer) ToRegistryComponentContainerOutput() RegistryComponentContainerOutput {
	return i.ToRegistryComponentContainerOutputWithContext(context.Background())
}

func (i *RegistryComponentContainer) ToRegistryComponentContainerOutputWithContext(ctx context.Context) RegistryComponentContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryComponentContainerOutput)
}

type RegistryComponentContainerOutput struct{ *pulumi.OutputState }

func (RegistryComponentContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryComponentContainer)(nil)).Elem()
}

func (o RegistryComponentContainerOutput) ToRegistryComponentContainerOutput() RegistryComponentContainerOutput {
	return o
}

func (o RegistryComponentContainerOutput) ToRegistryComponentContainerOutputWithContext(ctx context.Context) RegistryComponentContainerOutput {
	return o
}

func (o RegistryComponentContainerOutput) ComponentContainerProperties() ComponentContainerResponseOutput {
	return o.ApplyT(func(v *RegistryComponentContainer) ComponentContainerResponseOutput {
		return v.ComponentContainerProperties
	}).(ComponentContainerResponseOutput)
}

func (o RegistryComponentContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryComponentContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegistryComponentContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryComponentContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RegistryComponentContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryComponentContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryComponentContainerOutput{})
}
