


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistryModelVersion struct {
	pulumi.CustomResourceState

	ModelVersionProperties ModelVersionResponseOutput `pulumi:"modelVersionProperties"`
	Name                   pulumi.StringOutput        `pulumi:"name"`
	SystemData             SystemDataResponseOutput   `pulumi:"systemData"`
	Type                   pulumi.StringOutput        `pulumi:"type"`
}


func NewRegistryModelVersion(ctx *pulumi.Context,
	name string, args *RegistryModelVersionArgs, opts ...pulumi.ResourceOption) (*RegistryModelVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModelName == nil {
		return nil, errors.New("invalid value for required argument 'ModelName'")
	}
	if args.ModelVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'ModelVersionProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.ModelVersionProperties = args.ModelVersionProperties.ToModelVersionTypeOutput().ApplyT(func(v ModelVersionType) ModelVersionType { return *v.Defaults() }).(ModelVersionTypeOutput)
	var resource RegistryModelVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001preview:RegistryModelVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistryModelVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryModelVersionState, opts ...pulumi.ResourceOption) (*RegistryModelVersion, error) {
	var resource RegistryModelVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001preview:RegistryModelVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryModelVersionState struct {
}

type RegistryModelVersionState struct {
}

func (RegistryModelVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryModelVersionState)(nil)).Elem()
}

type registryModelVersionArgs struct {
	ModelName              string           `pulumi:"modelName"`
	ModelVersionProperties ModelVersionType `pulumi:"modelVersionProperties"`
	RegistryName           string           `pulumi:"registryName"`
	ResourceGroupName      string           `pulumi:"resourceGroupName"`
	Version                *string          `pulumi:"version"`
}


type RegistryModelVersionArgs struct {
	ModelName              pulumi.StringInput
	ModelVersionProperties ModelVersionTypeInput
	RegistryName           pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	Version                pulumi.StringPtrInput
}

func (RegistryModelVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryModelVersionArgs)(nil)).Elem()
}

type RegistryModelVersionInput interface {
	pulumi.Input

	ToRegistryModelVersionOutput() RegistryModelVersionOutput
	ToRegistryModelVersionOutputWithContext(ctx context.Context) RegistryModelVersionOutput
}

func (*RegistryModelVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryModelVersion)(nil)).Elem()
}

func (i *RegistryModelVersion) ToRegistryModelVersionOutput() RegistryModelVersionOutput {
	return i.ToRegistryModelVersionOutputWithContext(context.Background())
}

func (i *RegistryModelVersion) ToRegistryModelVersionOutputWithContext(ctx context.Context) RegistryModelVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryModelVersionOutput)
}

type RegistryModelVersionOutput struct{ *pulumi.OutputState }

func (RegistryModelVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryModelVersion)(nil)).Elem()
}

func (o RegistryModelVersionOutput) ToRegistryModelVersionOutput() RegistryModelVersionOutput {
	return o
}

func (o RegistryModelVersionOutput) ToRegistryModelVersionOutputWithContext(ctx context.Context) RegistryModelVersionOutput {
	return o
}

func (o RegistryModelVersionOutput) ModelVersionProperties() ModelVersionResponseOutput {
	return o.ApplyT(func(v *RegistryModelVersion) ModelVersionResponseOutput { return v.ModelVersionProperties }).(ModelVersionResponseOutput)
}

func (o RegistryModelVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegistryModelVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryModelVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RegistryModelVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryModelVersionOutput{})
}
