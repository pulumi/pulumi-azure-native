


package v20211031preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Distro string

const (
	DistroAKSEdge = Distro("AKSEdge")
)

func (Distro) ElementType() reflect.Type {
	return reflect.TypeOf((*Distro)(nil)).Elem()
}

func (e Distro) ToDistroOutput() DistroOutput {
	return pulumi.ToOutput(e).(DistroOutput)
}

func (e Distro) ToDistroOutputWithContext(ctx context.Context) DistroOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DistroOutput)
}

func (e Distro) ToDistroPtrOutput() DistroPtrOutput {
	return e.ToDistroPtrOutputWithContext(context.Background())
}

func (e Distro) ToDistroPtrOutputWithContext(ctx context.Context) DistroPtrOutput {
	return Distro(e).ToDistroOutputWithContext(ctx).ToDistroPtrOutputWithContext(ctx)
}

func (e Distro) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Distro) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Distro) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Distro) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DistroOutput struct{ *pulumi.OutputState }

func (DistroOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Distro)(nil)).Elem()
}

func (o DistroOutput) ToDistroOutput() DistroOutput {
	return o
}

func (o DistroOutput) ToDistroOutputWithContext(ctx context.Context) DistroOutput {
	return o
}

func (o DistroOutput) ToDistroPtrOutput() DistroPtrOutput {
	return o.ToDistroPtrOutputWithContext(context.Background())
}

func (o DistroOutput) ToDistroPtrOutputWithContext(ctx context.Context) DistroPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Distro) *Distro {
		return &v
	}).(DistroPtrOutput)
}

func (o DistroOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DistroOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Distro) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DistroOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DistroOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Distro) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DistroPtrOutput struct{ *pulumi.OutputState }

func (DistroPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Distro)(nil)).Elem()
}

func (o DistroPtrOutput) ToDistroPtrOutput() DistroPtrOutput {
	return o
}

func (o DistroPtrOutput) ToDistroPtrOutputWithContext(ctx context.Context) DistroPtrOutput {
	return o
}

func (o DistroPtrOutput) Elem() DistroOutput {
	return o.ApplyT(func(v *Distro) Distro {
		if v != nil {
			return *v
		}
		var ret Distro
		return ret
	}).(DistroOutput)
}

func (o DistroPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DistroPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Distro) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DistroInput interface {
	pulumi.Input

	ToDistroOutput() DistroOutput
	ToDistroOutputWithContext(context.Context) DistroOutput
}

var distroPtrType = reflect.TypeOf((**Distro)(nil)).Elem()

type DistroPtrInput interface {
	pulumi.Input

	ToDistroPtrOutput() DistroPtrOutput
	ToDistroPtrOutputWithContext(context.Context) DistroPtrOutput
}

type distroPtr string

func DistroPtr(v string) DistroPtrInput {
	return (*distroPtr)(&v)
}

func (*distroPtr) ElementType() reflect.Type {
	return distroPtrType
}

func (in *distroPtr) ToDistroPtrOutput() DistroPtrOutput {
	return pulumi.ToOutput(in).(DistroPtrOutput)
}

func (in *distroPtr) ToDistroPtrOutputWithContext(ctx context.Context) DistroPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DistroPtrOutput)
}

type Provider string

const (
	ProviderVMWare = Provider("VMWare")
	ProviderHCI    = Provider("HCI")
	ProviderSCVMM  = Provider("SCVMM")
)

func (Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (e Provider) ToProviderOutput() ProviderOutput {
	return pulumi.ToOutput(e).(ProviderOutput)
}

func (e Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProviderOutput)
}

func (e Provider) ToProviderPtrOutput() ProviderPtrOutput {
	return e.ToProviderPtrOutputWithContext(context.Background())
}

func (e Provider) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return Provider(e).ToProviderOutputWithContext(ctx).ToProviderPtrOutputWithContext(ctx)
}

func (e Provider) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Provider) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Provider) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Provider) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o.ToProviderPtrOutputWithContext(context.Background())
}

func (o ProviderOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Provider) *Provider {
		return &v
	}).(ProviderPtrOutput)
}

func (o ProviderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProviderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Provider) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProviderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProviderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Provider) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProviderPtrOutput struct{ *pulumi.OutputState }

func (ProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderPtrOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) Elem() ProviderOutput {
	return o.ApplyT(func(v *Provider) Provider {
		if v != nil {
			return *v
		}
		var ret Provider
		return ret
	}).(ProviderOutput)
}

func (o ProviderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProviderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Provider) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(context.Context) ProviderOutput
}

var providerPtrType = reflect.TypeOf((**Provider)(nil)).Elem()

type ProviderPtrInput interface {
	pulumi.Input

	ToProviderPtrOutput() ProviderPtrOutput
	ToProviderPtrOutputWithContext(context.Context) ProviderPtrOutput
}

type providerPtr string

func ProviderPtr(v string) ProviderPtrInput {
	return (*providerPtr)(&v)
}

func (*providerPtr) ElementType() reflect.Type {
	return providerPtrType
}

func (in *providerPtr) ToProviderPtrOutput() ProviderPtrOutput {
	return pulumi.ToOutput(in).(ProviderPtrOutput)
}

func (in *providerPtr) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProviderPtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DistroOutput{})
	pulumi.RegisterOutputType(DistroPtrOutput{})
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
