


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistryCodeContainer struct {
	pulumi.CustomResourceState

	CodeContainerProperties CodeContainerResponseOutput `pulumi:"codeContainerProperties"`
	Name                    pulumi.StringOutput         `pulumi:"name"`
	SystemData              SystemDataResponseOutput    `pulumi:"systemData"`
	Type                    pulumi.StringOutput         `pulumi:"type"`
}


func NewRegistryCodeContainer(ctx *pulumi.Context,
	name string, args *RegistryCodeContainerArgs, opts ...pulumi.ResourceOption) (*RegistryCodeContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CodeContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'CodeContainerProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.CodeContainerProperties = args.CodeContainerProperties.ToCodeContainerTypeOutput().ApplyT(func(v CodeContainerType) CodeContainerType { return *v.Defaults() }).(CodeContainerTypeOutput)
	var resource RegistryCodeContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001preview:RegistryCodeContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistryCodeContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryCodeContainerState, opts ...pulumi.ResourceOption) (*RegistryCodeContainer, error) {
	var resource RegistryCodeContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001preview:RegistryCodeContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryCodeContainerState struct {
}

type RegistryCodeContainerState struct {
}

func (RegistryCodeContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryCodeContainerState)(nil)).Elem()
}

type registryCodeContainerArgs struct {
	CodeContainerProperties CodeContainerType `pulumi:"codeContainerProperties"`
	CodeName                *string           `pulumi:"codeName"`
	RegistryName            string            `pulumi:"registryName"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
}


type RegistryCodeContainerArgs struct {
	CodeContainerProperties CodeContainerTypeInput
	CodeName                pulumi.StringPtrInput
	RegistryName            pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
}

func (RegistryCodeContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryCodeContainerArgs)(nil)).Elem()
}

type RegistryCodeContainerInput interface {
	pulumi.Input

	ToRegistryCodeContainerOutput() RegistryCodeContainerOutput
	ToRegistryCodeContainerOutputWithContext(ctx context.Context) RegistryCodeContainerOutput
}

func (*RegistryCodeContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryCodeContainer)(nil)).Elem()
}

func (i *RegistryCodeContainer) ToRegistryCodeContainerOutput() RegistryCodeContainerOutput {
	return i.ToRegistryCodeContainerOutputWithContext(context.Background())
}

func (i *RegistryCodeContainer) ToRegistryCodeContainerOutputWithContext(ctx context.Context) RegistryCodeContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryCodeContainerOutput)
}

type RegistryCodeContainerOutput struct{ *pulumi.OutputState }

func (RegistryCodeContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryCodeContainer)(nil)).Elem()
}

func (o RegistryCodeContainerOutput) ToRegistryCodeContainerOutput() RegistryCodeContainerOutput {
	return o
}

func (o RegistryCodeContainerOutput) ToRegistryCodeContainerOutputWithContext(ctx context.Context) RegistryCodeContainerOutput {
	return o
}

func (o RegistryCodeContainerOutput) CodeContainerProperties() CodeContainerResponseOutput {
	return o.ApplyT(func(v *RegistryCodeContainer) CodeContainerResponseOutput { return v.CodeContainerProperties }).(CodeContainerResponseOutput)
}

func (o RegistryCodeContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryCodeContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegistryCodeContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryCodeContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RegistryCodeContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryCodeContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryCodeContainerOutput{})
}
