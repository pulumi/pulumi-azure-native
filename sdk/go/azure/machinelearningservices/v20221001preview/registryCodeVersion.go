


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistryCodeVersion struct {
	pulumi.CustomResourceState

	CodeVersionProperties CodeVersionResponseOutput `pulumi:"codeVersionProperties"`
	Name                  pulumi.StringOutput       `pulumi:"name"`
	SystemData            SystemDataResponseOutput  `pulumi:"systemData"`
	Type                  pulumi.StringOutput       `pulumi:"type"`
}


func NewRegistryCodeVersion(ctx *pulumi.Context,
	name string, args *RegistryCodeVersionArgs, opts ...pulumi.ResourceOption) (*RegistryCodeVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CodeName == nil {
		return nil, errors.New("invalid value for required argument 'CodeName'")
	}
	if args.CodeVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'CodeVersionProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.CodeVersionProperties = args.CodeVersionProperties.ToCodeVersionTypeOutput().ApplyT(func(v CodeVersionType) CodeVersionType { return *v.Defaults() }).(CodeVersionTypeOutput)
	var resource RegistryCodeVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001preview:RegistryCodeVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistryCodeVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryCodeVersionState, opts ...pulumi.ResourceOption) (*RegistryCodeVersion, error) {
	var resource RegistryCodeVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001preview:RegistryCodeVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryCodeVersionState struct {
}

type RegistryCodeVersionState struct {
}

func (RegistryCodeVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryCodeVersionState)(nil)).Elem()
}

type registryCodeVersionArgs struct {
	CodeName              string          `pulumi:"codeName"`
	CodeVersionProperties CodeVersionType `pulumi:"codeVersionProperties"`
	RegistryName          string          `pulumi:"registryName"`
	ResourceGroupName     string          `pulumi:"resourceGroupName"`
	Version               *string         `pulumi:"version"`
}


type RegistryCodeVersionArgs struct {
	CodeName              pulumi.StringInput
	CodeVersionProperties CodeVersionTypeInput
	RegistryName          pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Version               pulumi.StringPtrInput
}

func (RegistryCodeVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryCodeVersionArgs)(nil)).Elem()
}

type RegistryCodeVersionInput interface {
	pulumi.Input

	ToRegistryCodeVersionOutput() RegistryCodeVersionOutput
	ToRegistryCodeVersionOutputWithContext(ctx context.Context) RegistryCodeVersionOutput
}

func (*RegistryCodeVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryCodeVersion)(nil)).Elem()
}

func (i *RegistryCodeVersion) ToRegistryCodeVersionOutput() RegistryCodeVersionOutput {
	return i.ToRegistryCodeVersionOutputWithContext(context.Background())
}

func (i *RegistryCodeVersion) ToRegistryCodeVersionOutputWithContext(ctx context.Context) RegistryCodeVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryCodeVersionOutput)
}

type RegistryCodeVersionOutput struct{ *pulumi.OutputState }

func (RegistryCodeVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryCodeVersion)(nil)).Elem()
}

func (o RegistryCodeVersionOutput) ToRegistryCodeVersionOutput() RegistryCodeVersionOutput {
	return o
}

func (o RegistryCodeVersionOutput) ToRegistryCodeVersionOutputWithContext(ctx context.Context) RegistryCodeVersionOutput {
	return o
}

func (o RegistryCodeVersionOutput) CodeVersionProperties() CodeVersionResponseOutput {
	return o.ApplyT(func(v *RegistryCodeVersion) CodeVersionResponseOutput { return v.CodeVersionProperties }).(CodeVersionResponseOutput)
}

func (o RegistryCodeVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryCodeVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegistryCodeVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryCodeVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RegistryCodeVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryCodeVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryCodeVersionOutput{})
}
