


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Registry struct {
	pulumi.CustomResourceState

	Identity           ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind               pulumi.StringPtrOutput                  `pulumi:"kind"`
	Location           pulumi.StringOutput                     `pulumi:"location"`
	Name               pulumi.StringOutput                     `pulumi:"name"`
	RegistryProperties RegistryResponseOutput                  `pulumi:"registryProperties"`
	Sku                SkuResponsePtrOutput                    `pulumi:"sku"`
	SystemData         SystemDataResponseOutput                `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput                  `pulumi:"tags"`
	Type               pulumi.StringOutput                     `pulumi:"type"`
}


func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryProperties == nil {
		return nil, errors.New("invalid value for required argument 'RegistryProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Registry
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001preview:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001preview:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryState struct {
}

type RegistryState struct {
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	Identity           *ManagedServiceIdentity `pulumi:"identity"`
	Kind               *string                 `pulumi:"kind"`
	Location           *string                 `pulumi:"location"`
	RegistryName       *string                 `pulumi:"registryName"`
	RegistryProperties RegistryType            `pulumi:"registryProperties"`
	ResourceGroupName  string                  `pulumi:"resourceGroupName"`
	Sku                *Sku                    `pulumi:"sku"`
	Tags               map[string]string       `pulumi:"tags"`
}


type RegistryArgs struct {
	Identity           ManagedServiceIdentityPtrInput
	Kind               pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	RegistryName       pulumi.StringPtrInput
	RegistryProperties RegistryTypeInput
	ResourceGroupName  pulumi.StringInput
	Sku                SkuPtrInput
	Tags               pulumi.StringMapInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}

type RegistryInput interface {
	pulumi.Input

	ToRegistryOutput() RegistryOutput
	ToRegistryOutputWithContext(ctx context.Context) RegistryOutput
}

func (*Registry) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (i *Registry) ToRegistryOutput() RegistryOutput {
	return i.ToRegistryOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryOutput)
}

type RegistryOutput struct{ *pulumi.OutputState }

func (RegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (o RegistryOutput) ToRegistryOutput() RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return o
}

func (o RegistryOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Registry) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o RegistryOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o RegistryOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o RegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegistryOutput) RegistryProperties() RegistryResponseOutput {
	return o.ApplyT(func(v *Registry) RegistryResponseOutput { return v.RegistryProperties }).(RegistryResponseOutput)
}

func (o RegistryOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Registry) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o RegistryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Registry) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RegistryOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryOutput{})
}
