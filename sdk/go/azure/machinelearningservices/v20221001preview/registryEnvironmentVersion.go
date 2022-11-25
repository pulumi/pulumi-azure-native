


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistryEnvironmentVersion struct {
	pulumi.CustomResourceState

	EnvironmentVersionProperties EnvironmentVersionResponseOutput `pulumi:"environmentVersionProperties"`
	Name                         pulumi.StringOutput              `pulumi:"name"`
	SystemData                   SystemDataResponseOutput         `pulumi:"systemData"`
	Type                         pulumi.StringOutput              `pulumi:"type"`
}


func NewRegistryEnvironmentVersion(ctx *pulumi.Context,
	name string, args *RegistryEnvironmentVersionArgs, opts ...pulumi.ResourceOption) (*RegistryEnvironmentVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.EnvironmentVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentVersionProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.EnvironmentVersionProperties = args.EnvironmentVersionProperties.ToEnvironmentVersionTypeOutput().ApplyT(func(v EnvironmentVersionType) EnvironmentVersionType { return *v.Defaults() }).(EnvironmentVersionTypeOutput)
	var resource RegistryEnvironmentVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001preview:RegistryEnvironmentVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistryEnvironmentVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryEnvironmentVersionState, opts ...pulumi.ResourceOption) (*RegistryEnvironmentVersion, error) {
	var resource RegistryEnvironmentVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001preview:RegistryEnvironmentVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryEnvironmentVersionState struct {
}

type RegistryEnvironmentVersionState struct {
}

func (RegistryEnvironmentVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnvironmentVersionState)(nil)).Elem()
}

type registryEnvironmentVersionArgs struct {
	EnvironmentName              string                 `pulumi:"environmentName"`
	EnvironmentVersionProperties EnvironmentVersionType `pulumi:"environmentVersionProperties"`
	RegistryName                 string                 `pulumi:"registryName"`
	ResourceGroupName            string                 `pulumi:"resourceGroupName"`
	Version                      *string                `pulumi:"version"`
}


type RegistryEnvironmentVersionArgs struct {
	EnvironmentName              pulumi.StringInput
	EnvironmentVersionProperties EnvironmentVersionTypeInput
	RegistryName                 pulumi.StringInput
	ResourceGroupName            pulumi.StringInput
	Version                      pulumi.StringPtrInput
}

func (RegistryEnvironmentVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnvironmentVersionArgs)(nil)).Elem()
}

type RegistryEnvironmentVersionInput interface {
	pulumi.Input

	ToRegistryEnvironmentVersionOutput() RegistryEnvironmentVersionOutput
	ToRegistryEnvironmentVersionOutputWithContext(ctx context.Context) RegistryEnvironmentVersionOutput
}

func (*RegistryEnvironmentVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnvironmentVersion)(nil)).Elem()
}

func (i *RegistryEnvironmentVersion) ToRegistryEnvironmentVersionOutput() RegistryEnvironmentVersionOutput {
	return i.ToRegistryEnvironmentVersionOutputWithContext(context.Background())
}

func (i *RegistryEnvironmentVersion) ToRegistryEnvironmentVersionOutputWithContext(ctx context.Context) RegistryEnvironmentVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnvironmentVersionOutput)
}

type RegistryEnvironmentVersionOutput struct{ *pulumi.OutputState }

func (RegistryEnvironmentVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnvironmentVersion)(nil)).Elem()
}

func (o RegistryEnvironmentVersionOutput) ToRegistryEnvironmentVersionOutput() RegistryEnvironmentVersionOutput {
	return o
}

func (o RegistryEnvironmentVersionOutput) ToRegistryEnvironmentVersionOutputWithContext(ctx context.Context) RegistryEnvironmentVersionOutput {
	return o
}

func (o RegistryEnvironmentVersionOutput) EnvironmentVersionProperties() EnvironmentVersionResponseOutput {
	return o.ApplyT(func(v *RegistryEnvironmentVersion) EnvironmentVersionResponseOutput {
		return v.EnvironmentVersionProperties
	}).(EnvironmentVersionResponseOutput)
}

func (o RegistryEnvironmentVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnvironmentVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegistryEnvironmentVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryEnvironmentVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RegistryEnvironmentVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnvironmentVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryEnvironmentVersionOutput{})
}
