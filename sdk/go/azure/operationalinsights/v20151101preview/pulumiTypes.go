


package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MachineReferenceWithHints struct {
	Id   string `pulumi:"id"`
	Kind string `pulumi:"kind"`
}





type MachineReferenceWithHintsInput interface {
	pulumi.Input

	ToMachineReferenceWithHintsOutput() MachineReferenceWithHintsOutput
	ToMachineReferenceWithHintsOutputWithContext(context.Context) MachineReferenceWithHintsOutput
}

type MachineReferenceWithHintsArgs struct {
	Id   pulumi.StringInput `pulumi:"id"`
	Kind pulumi.StringInput `pulumi:"kind"`
}

func (MachineReferenceWithHintsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineReferenceWithHints)(nil)).Elem()
}

func (i MachineReferenceWithHintsArgs) ToMachineReferenceWithHintsOutput() MachineReferenceWithHintsOutput {
	return i.ToMachineReferenceWithHintsOutputWithContext(context.Background())
}

func (i MachineReferenceWithHintsArgs) ToMachineReferenceWithHintsOutputWithContext(ctx context.Context) MachineReferenceWithHintsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineReferenceWithHintsOutput)
}





type MachineReferenceWithHintsArrayInput interface {
	pulumi.Input

	ToMachineReferenceWithHintsArrayOutput() MachineReferenceWithHintsArrayOutput
	ToMachineReferenceWithHintsArrayOutputWithContext(context.Context) MachineReferenceWithHintsArrayOutput
}

type MachineReferenceWithHintsArray []MachineReferenceWithHintsInput

func (MachineReferenceWithHintsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineReferenceWithHints)(nil)).Elem()
}

func (i MachineReferenceWithHintsArray) ToMachineReferenceWithHintsArrayOutput() MachineReferenceWithHintsArrayOutput {
	return i.ToMachineReferenceWithHintsArrayOutputWithContext(context.Background())
}

func (i MachineReferenceWithHintsArray) ToMachineReferenceWithHintsArrayOutputWithContext(ctx context.Context) MachineReferenceWithHintsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineReferenceWithHintsArrayOutput)
}

type MachineReferenceWithHintsOutput struct{ *pulumi.OutputState }

func (MachineReferenceWithHintsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineReferenceWithHints)(nil)).Elem()
}

func (o MachineReferenceWithHintsOutput) ToMachineReferenceWithHintsOutput() MachineReferenceWithHintsOutput {
	return o
}

func (o MachineReferenceWithHintsOutput) ToMachineReferenceWithHintsOutputWithContext(ctx context.Context) MachineReferenceWithHintsOutput {
	return o
}

func (o MachineReferenceWithHintsOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHints) string { return v.Id }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHints) string { return v.Kind }).(pulumi.StringOutput)
}

type MachineReferenceWithHintsArrayOutput struct{ *pulumi.OutputState }

func (MachineReferenceWithHintsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineReferenceWithHints)(nil)).Elem()
}

func (o MachineReferenceWithHintsArrayOutput) ToMachineReferenceWithHintsArrayOutput() MachineReferenceWithHintsArrayOutput {
	return o
}

func (o MachineReferenceWithHintsArrayOutput) ToMachineReferenceWithHintsArrayOutputWithContext(ctx context.Context) MachineReferenceWithHintsArrayOutput {
	return o
}

func (o MachineReferenceWithHintsArrayOutput) Index(i pulumi.IntInput) MachineReferenceWithHintsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineReferenceWithHints {
		return vs[0].([]MachineReferenceWithHints)[vs[1].(int)]
	}).(MachineReferenceWithHintsOutput)
}

type MachineReferenceWithHintsResponse struct {
	DisplayNameHint string `pulumi:"displayNameHint"`
	Id              string `pulumi:"id"`
	Kind            string `pulumi:"kind"`
	Name            string `pulumi:"name"`
	OsFamilyHint    string `pulumi:"osFamilyHint"`
	Type            string `pulumi:"type"`
}

type MachineReferenceWithHintsResponseOutput struct{ *pulumi.OutputState }

func (MachineReferenceWithHintsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineReferenceWithHintsResponse)(nil)).Elem()
}

func (o MachineReferenceWithHintsResponseOutput) ToMachineReferenceWithHintsResponseOutput() MachineReferenceWithHintsResponseOutput {
	return o
}

func (o MachineReferenceWithHintsResponseOutput) ToMachineReferenceWithHintsResponseOutputWithContext(ctx context.Context) MachineReferenceWithHintsResponseOutput {
	return o
}

func (o MachineReferenceWithHintsResponseOutput) DisplayNameHint() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.DisplayNameHint }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsResponseOutput) OsFamilyHint() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.OsFamilyHint }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.Type }).(pulumi.StringOutput)
}

type MachineReferenceWithHintsResponseArrayOutput struct{ *pulumi.OutputState }

func (MachineReferenceWithHintsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineReferenceWithHintsResponse)(nil)).Elem()
}

func (o MachineReferenceWithHintsResponseArrayOutput) ToMachineReferenceWithHintsResponseArrayOutput() MachineReferenceWithHintsResponseArrayOutput {
	return o
}

func (o MachineReferenceWithHintsResponseArrayOutput) ToMachineReferenceWithHintsResponseArrayOutputWithContext(ctx context.Context) MachineReferenceWithHintsResponseArrayOutput {
	return o
}

func (o MachineReferenceWithHintsResponseArrayOutput) Index(i pulumi.IntInput) MachineReferenceWithHintsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineReferenceWithHintsResponse {
		return vs[0].([]MachineReferenceWithHintsResponse)[vs[1].(int)]
	}).(MachineReferenceWithHintsResponseOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineReferenceWithHintsOutput{})
	pulumi.RegisterOutputType(MachineReferenceWithHintsArrayOutput{})
	pulumi.RegisterOutputType(MachineReferenceWithHintsResponseOutput{})
	pulumi.RegisterOutputType(MachineReferenceWithHintsResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
