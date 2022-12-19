


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistryComponentVersion struct {
	pulumi.CustomResourceState

	ComponentVersionProperties ComponentVersionResponseOutput `pulumi:"componentVersionProperties"`
	Name                       pulumi.StringOutput            `pulumi:"name"`
	SystemData                 SystemDataResponseOutput       `pulumi:"systemData"`
	Type                       pulumi.StringOutput            `pulumi:"type"`
}


func NewRegistryComponentVersion(ctx *pulumi.Context,
	name string, args *RegistryComponentVersionArgs, opts ...pulumi.ResourceOption) (*RegistryComponentVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComponentName == nil {
		return nil, errors.New("invalid value for required argument 'ComponentName'")
	}
	if args.ComponentVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'ComponentVersionProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.ComponentVersionProperties = args.ComponentVersionProperties.ToComponentVersionTypeOutput().ApplyT(func(v ComponentVersionType) ComponentVersionType { return *v.Defaults() }).(ComponentVersionTypeOutput)
	var resource RegistryComponentVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001preview:RegistryComponentVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistryComponentVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryComponentVersionState, opts ...pulumi.ResourceOption) (*RegistryComponentVersion, error) {
	var resource RegistryComponentVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001preview:RegistryComponentVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryComponentVersionState struct {
}

type RegistryComponentVersionState struct {
}

func (RegistryComponentVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryComponentVersionState)(nil)).Elem()
}

type registryComponentVersionArgs struct {
	ComponentName              string               `pulumi:"componentName"`
	ComponentVersionProperties ComponentVersionType `pulumi:"componentVersionProperties"`
	RegistryName               string               `pulumi:"registryName"`
	ResourceGroupName          string               `pulumi:"resourceGroupName"`
	Version                    *string              `pulumi:"version"`
}


type RegistryComponentVersionArgs struct {
	ComponentName              pulumi.StringInput
	ComponentVersionProperties ComponentVersionTypeInput
	RegistryName               pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	Version                    pulumi.StringPtrInput
}

func (RegistryComponentVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryComponentVersionArgs)(nil)).Elem()
}

type RegistryComponentVersionInput interface {
	pulumi.Input

	ToRegistryComponentVersionOutput() RegistryComponentVersionOutput
	ToRegistryComponentVersionOutputWithContext(ctx context.Context) RegistryComponentVersionOutput
}

func (*RegistryComponentVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryComponentVersion)(nil)).Elem()
}

func (i *RegistryComponentVersion) ToRegistryComponentVersionOutput() RegistryComponentVersionOutput {
	return i.ToRegistryComponentVersionOutputWithContext(context.Background())
}

func (i *RegistryComponentVersion) ToRegistryComponentVersionOutputWithContext(ctx context.Context) RegistryComponentVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryComponentVersionOutput)
}

type RegistryComponentVersionOutput struct{ *pulumi.OutputState }

func (RegistryComponentVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryComponentVersion)(nil)).Elem()
}

func (o RegistryComponentVersionOutput) ToRegistryComponentVersionOutput() RegistryComponentVersionOutput {
	return o
}

func (o RegistryComponentVersionOutput) ToRegistryComponentVersionOutputWithContext(ctx context.Context) RegistryComponentVersionOutput {
	return o
}

func (o RegistryComponentVersionOutput) ComponentVersionProperties() ComponentVersionResponseOutput {
	return o.ApplyT(func(v *RegistryComponentVersion) ComponentVersionResponseOutput { return v.ComponentVersionProperties }).(ComponentVersionResponseOutput)
}

func (o RegistryComponentVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryComponentVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegistryComponentVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryComponentVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RegistryComponentVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryComponentVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryComponentVersionOutput{})
}
