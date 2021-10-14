


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutoScaleVCoreSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type AutoScaleVCoreSkuInput interface {
	pulumi.Input

	ToAutoScaleVCoreSkuOutput() AutoScaleVCoreSkuOutput
	ToAutoScaleVCoreSkuOutputWithContext(context.Context) AutoScaleVCoreSkuOutput
}

type AutoScaleVCoreSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (AutoScaleVCoreSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleVCoreSku)(nil)).Elem()
}

func (i AutoScaleVCoreSkuArgs) ToAutoScaleVCoreSkuOutput() AutoScaleVCoreSkuOutput {
	return i.ToAutoScaleVCoreSkuOutputWithContext(context.Background())
}

func (i AutoScaleVCoreSkuArgs) ToAutoScaleVCoreSkuOutputWithContext(ctx context.Context) AutoScaleVCoreSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleVCoreSkuOutput)
}

func (i AutoScaleVCoreSkuArgs) ToAutoScaleVCoreSkuPtrOutput() AutoScaleVCoreSkuPtrOutput {
	return i.ToAutoScaleVCoreSkuPtrOutputWithContext(context.Background())
}

func (i AutoScaleVCoreSkuArgs) ToAutoScaleVCoreSkuPtrOutputWithContext(ctx context.Context) AutoScaleVCoreSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleVCoreSkuOutput).ToAutoScaleVCoreSkuPtrOutputWithContext(ctx)
}









type AutoScaleVCoreSkuPtrInput interface {
	pulumi.Input

	ToAutoScaleVCoreSkuPtrOutput() AutoScaleVCoreSkuPtrOutput
	ToAutoScaleVCoreSkuPtrOutputWithContext(context.Context) AutoScaleVCoreSkuPtrOutput
}

type autoScaleVCoreSkuPtrType AutoScaleVCoreSkuArgs

func AutoScaleVCoreSkuPtr(v *AutoScaleVCoreSkuArgs) AutoScaleVCoreSkuPtrInput {
	return (*autoScaleVCoreSkuPtrType)(v)
}

func (*autoScaleVCoreSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleVCoreSku)(nil)).Elem()
}

func (i *autoScaleVCoreSkuPtrType) ToAutoScaleVCoreSkuPtrOutput() AutoScaleVCoreSkuPtrOutput {
	return i.ToAutoScaleVCoreSkuPtrOutputWithContext(context.Background())
}

func (i *autoScaleVCoreSkuPtrType) ToAutoScaleVCoreSkuPtrOutputWithContext(ctx context.Context) AutoScaleVCoreSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleVCoreSkuPtrOutput)
}

type AutoScaleVCoreSkuOutput struct{ *pulumi.OutputState }

func (AutoScaleVCoreSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleVCoreSku)(nil)).Elem()
}

func (o AutoScaleVCoreSkuOutput) ToAutoScaleVCoreSkuOutput() AutoScaleVCoreSkuOutput {
	return o
}

func (o AutoScaleVCoreSkuOutput) ToAutoScaleVCoreSkuOutputWithContext(ctx context.Context) AutoScaleVCoreSkuOutput {
	return o
}

func (o AutoScaleVCoreSkuOutput) ToAutoScaleVCoreSkuPtrOutput() AutoScaleVCoreSkuPtrOutput {
	return o.ToAutoScaleVCoreSkuPtrOutputWithContext(context.Background())
}

func (o AutoScaleVCoreSkuOutput) ToAutoScaleVCoreSkuPtrOutputWithContext(ctx context.Context) AutoScaleVCoreSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScaleVCoreSku) *AutoScaleVCoreSku {
		return &v
	}).(AutoScaleVCoreSkuPtrOutput)
}

func (o AutoScaleVCoreSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o AutoScaleVCoreSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleVCoreSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o AutoScaleVCoreSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type AutoScaleVCoreSkuPtrOutput struct{ *pulumi.OutputState }

func (AutoScaleVCoreSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleVCoreSku)(nil)).Elem()
}

func (o AutoScaleVCoreSkuPtrOutput) ToAutoScaleVCoreSkuPtrOutput() AutoScaleVCoreSkuPtrOutput {
	return o
}

func (o AutoScaleVCoreSkuPtrOutput) ToAutoScaleVCoreSkuPtrOutputWithContext(ctx context.Context) AutoScaleVCoreSkuPtrOutput {
	return o
}

func (o AutoScaleVCoreSkuPtrOutput) Elem() AutoScaleVCoreSkuOutput {
	return o.ApplyT(func(v *AutoScaleVCoreSku) AutoScaleVCoreSku {
		if v != nil {
			return *v
		}
		var ret AutoScaleVCoreSku
		return ret
	}).(AutoScaleVCoreSkuOutput)
}

func (o AutoScaleVCoreSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleVCoreSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o AutoScaleVCoreSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleVCoreSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o AutoScaleVCoreSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleVCoreSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type AutoScaleVCoreSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type AutoScaleVCoreSkuResponseInput interface {
	pulumi.Input

	ToAutoScaleVCoreSkuResponseOutput() AutoScaleVCoreSkuResponseOutput
	ToAutoScaleVCoreSkuResponseOutputWithContext(context.Context) AutoScaleVCoreSkuResponseOutput
}

type AutoScaleVCoreSkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (AutoScaleVCoreSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleVCoreSkuResponse)(nil)).Elem()
}

func (i AutoScaleVCoreSkuResponseArgs) ToAutoScaleVCoreSkuResponseOutput() AutoScaleVCoreSkuResponseOutput {
	return i.ToAutoScaleVCoreSkuResponseOutputWithContext(context.Background())
}

func (i AutoScaleVCoreSkuResponseArgs) ToAutoScaleVCoreSkuResponseOutputWithContext(ctx context.Context) AutoScaleVCoreSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleVCoreSkuResponseOutput)
}

func (i AutoScaleVCoreSkuResponseArgs) ToAutoScaleVCoreSkuResponsePtrOutput() AutoScaleVCoreSkuResponsePtrOutput {
	return i.ToAutoScaleVCoreSkuResponsePtrOutputWithContext(context.Background())
}

func (i AutoScaleVCoreSkuResponseArgs) ToAutoScaleVCoreSkuResponsePtrOutputWithContext(ctx context.Context) AutoScaleVCoreSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleVCoreSkuResponseOutput).ToAutoScaleVCoreSkuResponsePtrOutputWithContext(ctx)
}









type AutoScaleVCoreSkuResponsePtrInput interface {
	pulumi.Input

	ToAutoScaleVCoreSkuResponsePtrOutput() AutoScaleVCoreSkuResponsePtrOutput
	ToAutoScaleVCoreSkuResponsePtrOutputWithContext(context.Context) AutoScaleVCoreSkuResponsePtrOutput
}

type autoScaleVCoreSkuResponsePtrType AutoScaleVCoreSkuResponseArgs

func AutoScaleVCoreSkuResponsePtr(v *AutoScaleVCoreSkuResponseArgs) AutoScaleVCoreSkuResponsePtrInput {
	return (*autoScaleVCoreSkuResponsePtrType)(v)
}

func (*autoScaleVCoreSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleVCoreSkuResponse)(nil)).Elem()
}

func (i *autoScaleVCoreSkuResponsePtrType) ToAutoScaleVCoreSkuResponsePtrOutput() AutoScaleVCoreSkuResponsePtrOutput {
	return i.ToAutoScaleVCoreSkuResponsePtrOutputWithContext(context.Background())
}

func (i *autoScaleVCoreSkuResponsePtrType) ToAutoScaleVCoreSkuResponsePtrOutputWithContext(ctx context.Context) AutoScaleVCoreSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleVCoreSkuResponsePtrOutput)
}

type AutoScaleVCoreSkuResponseOutput struct{ *pulumi.OutputState }

func (AutoScaleVCoreSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleVCoreSkuResponse)(nil)).Elem()
}

func (o AutoScaleVCoreSkuResponseOutput) ToAutoScaleVCoreSkuResponseOutput() AutoScaleVCoreSkuResponseOutput {
	return o
}

func (o AutoScaleVCoreSkuResponseOutput) ToAutoScaleVCoreSkuResponseOutputWithContext(ctx context.Context) AutoScaleVCoreSkuResponseOutput {
	return o
}

func (o AutoScaleVCoreSkuResponseOutput) ToAutoScaleVCoreSkuResponsePtrOutput() AutoScaleVCoreSkuResponsePtrOutput {
	return o.ToAutoScaleVCoreSkuResponsePtrOutputWithContext(context.Background())
}

func (o AutoScaleVCoreSkuResponseOutput) ToAutoScaleVCoreSkuResponsePtrOutputWithContext(ctx context.Context) AutoScaleVCoreSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScaleVCoreSkuResponse) *AutoScaleVCoreSkuResponse {
		return &v
	}).(AutoScaleVCoreSkuResponsePtrOutput)
}

func (o AutoScaleVCoreSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o AutoScaleVCoreSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleVCoreSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AutoScaleVCoreSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type AutoScaleVCoreSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoScaleVCoreSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleVCoreSkuResponse)(nil)).Elem()
}

func (o AutoScaleVCoreSkuResponsePtrOutput) ToAutoScaleVCoreSkuResponsePtrOutput() AutoScaleVCoreSkuResponsePtrOutput {
	return o
}

func (o AutoScaleVCoreSkuResponsePtrOutput) ToAutoScaleVCoreSkuResponsePtrOutputWithContext(ctx context.Context) AutoScaleVCoreSkuResponsePtrOutput {
	return o
}

func (o AutoScaleVCoreSkuResponsePtrOutput) Elem() AutoScaleVCoreSkuResponseOutput {
	return o.ApplyT(func(v *AutoScaleVCoreSkuResponse) AutoScaleVCoreSkuResponse {
		if v != nil {
			return *v
		}
		var ret AutoScaleVCoreSkuResponse
		return ret
	}).(AutoScaleVCoreSkuResponseOutput)
}

func (o AutoScaleVCoreSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleVCoreSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o AutoScaleVCoreSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleVCoreSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o AutoScaleVCoreSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleVCoreSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type CapacitySku struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type CapacitySkuInput interface {
	pulumi.Input

	ToCapacitySkuOutput() CapacitySkuOutput
	ToCapacitySkuOutputWithContext(context.Context) CapacitySkuOutput
}

type CapacitySkuArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (CapacitySkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacitySku)(nil)).Elem()
}

func (i CapacitySkuArgs) ToCapacitySkuOutput() CapacitySkuOutput {
	return i.ToCapacitySkuOutputWithContext(context.Background())
}

func (i CapacitySkuArgs) ToCapacitySkuOutputWithContext(ctx context.Context) CapacitySkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacitySkuOutput)
}

func (i CapacitySkuArgs) ToCapacitySkuPtrOutput() CapacitySkuPtrOutput {
	return i.ToCapacitySkuPtrOutputWithContext(context.Background())
}

func (i CapacitySkuArgs) ToCapacitySkuPtrOutputWithContext(ctx context.Context) CapacitySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacitySkuOutput).ToCapacitySkuPtrOutputWithContext(ctx)
}









type CapacitySkuPtrInput interface {
	pulumi.Input

	ToCapacitySkuPtrOutput() CapacitySkuPtrOutput
	ToCapacitySkuPtrOutputWithContext(context.Context) CapacitySkuPtrOutput
}

type capacitySkuPtrType CapacitySkuArgs

func CapacitySkuPtr(v *CapacitySkuArgs) CapacitySkuPtrInput {
	return (*capacitySkuPtrType)(v)
}

func (*capacitySkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacitySku)(nil)).Elem()
}

func (i *capacitySkuPtrType) ToCapacitySkuPtrOutput() CapacitySkuPtrOutput {
	return i.ToCapacitySkuPtrOutputWithContext(context.Background())
}

func (i *capacitySkuPtrType) ToCapacitySkuPtrOutputWithContext(ctx context.Context) CapacitySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacitySkuPtrOutput)
}

type CapacitySkuOutput struct{ *pulumi.OutputState }

func (CapacitySkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacitySku)(nil)).Elem()
}

func (o CapacitySkuOutput) ToCapacitySkuOutput() CapacitySkuOutput {
	return o
}

func (o CapacitySkuOutput) ToCapacitySkuOutputWithContext(ctx context.Context) CapacitySkuOutput {
	return o
}

func (o CapacitySkuOutput) ToCapacitySkuPtrOutput() CapacitySkuPtrOutput {
	return o.ToCapacitySkuPtrOutputWithContext(context.Background())
}

func (o CapacitySkuOutput) ToCapacitySkuPtrOutputWithContext(ctx context.Context) CapacitySkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CapacitySku) *CapacitySku {
		return &v
	}).(CapacitySkuPtrOutput)
}

func (o CapacitySkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CapacitySku) string { return v.Name }).(pulumi.StringOutput)
}

func (o CapacitySkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapacitySku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type CapacitySkuPtrOutput struct{ *pulumi.OutputState }

func (CapacitySkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacitySku)(nil)).Elem()
}

func (o CapacitySkuPtrOutput) ToCapacitySkuPtrOutput() CapacitySkuPtrOutput {
	return o
}

func (o CapacitySkuPtrOutput) ToCapacitySkuPtrOutputWithContext(ctx context.Context) CapacitySkuPtrOutput {
	return o
}

func (o CapacitySkuPtrOutput) Elem() CapacitySkuOutput {
	return o.ApplyT(func(v *CapacitySku) CapacitySku {
		if v != nil {
			return *v
		}
		var ret CapacitySku
		return ret
	}).(CapacitySkuOutput)
}

func (o CapacitySkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapacitySku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CapacitySkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapacitySku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type CapacitySkuResponse struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type CapacitySkuResponseInput interface {
	pulumi.Input

	ToCapacitySkuResponseOutput() CapacitySkuResponseOutput
	ToCapacitySkuResponseOutputWithContext(context.Context) CapacitySkuResponseOutput
}

type CapacitySkuResponseArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (CapacitySkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacitySkuResponse)(nil)).Elem()
}

func (i CapacitySkuResponseArgs) ToCapacitySkuResponseOutput() CapacitySkuResponseOutput {
	return i.ToCapacitySkuResponseOutputWithContext(context.Background())
}

func (i CapacitySkuResponseArgs) ToCapacitySkuResponseOutputWithContext(ctx context.Context) CapacitySkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacitySkuResponseOutput)
}

func (i CapacitySkuResponseArgs) ToCapacitySkuResponsePtrOutput() CapacitySkuResponsePtrOutput {
	return i.ToCapacitySkuResponsePtrOutputWithContext(context.Background())
}

func (i CapacitySkuResponseArgs) ToCapacitySkuResponsePtrOutputWithContext(ctx context.Context) CapacitySkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacitySkuResponseOutput).ToCapacitySkuResponsePtrOutputWithContext(ctx)
}









type CapacitySkuResponsePtrInput interface {
	pulumi.Input

	ToCapacitySkuResponsePtrOutput() CapacitySkuResponsePtrOutput
	ToCapacitySkuResponsePtrOutputWithContext(context.Context) CapacitySkuResponsePtrOutput
}

type capacitySkuResponsePtrType CapacitySkuResponseArgs

func CapacitySkuResponsePtr(v *CapacitySkuResponseArgs) CapacitySkuResponsePtrInput {
	return (*capacitySkuResponsePtrType)(v)
}

func (*capacitySkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacitySkuResponse)(nil)).Elem()
}

func (i *capacitySkuResponsePtrType) ToCapacitySkuResponsePtrOutput() CapacitySkuResponsePtrOutput {
	return i.ToCapacitySkuResponsePtrOutputWithContext(context.Background())
}

func (i *capacitySkuResponsePtrType) ToCapacitySkuResponsePtrOutputWithContext(ctx context.Context) CapacitySkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacitySkuResponsePtrOutput)
}

type CapacitySkuResponseOutput struct{ *pulumi.OutputState }

func (CapacitySkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacitySkuResponse)(nil)).Elem()
}

func (o CapacitySkuResponseOutput) ToCapacitySkuResponseOutput() CapacitySkuResponseOutput {
	return o
}

func (o CapacitySkuResponseOutput) ToCapacitySkuResponseOutputWithContext(ctx context.Context) CapacitySkuResponseOutput {
	return o
}

func (o CapacitySkuResponseOutput) ToCapacitySkuResponsePtrOutput() CapacitySkuResponsePtrOutput {
	return o.ToCapacitySkuResponsePtrOutputWithContext(context.Background())
}

func (o CapacitySkuResponseOutput) ToCapacitySkuResponsePtrOutputWithContext(ctx context.Context) CapacitySkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CapacitySkuResponse) *CapacitySkuResponse {
		return &v
	}).(CapacitySkuResponsePtrOutput)
}

func (o CapacitySkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CapacitySkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CapacitySkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapacitySkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type CapacitySkuResponsePtrOutput struct{ *pulumi.OutputState }

func (CapacitySkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacitySkuResponse)(nil)).Elem()
}

func (o CapacitySkuResponsePtrOutput) ToCapacitySkuResponsePtrOutput() CapacitySkuResponsePtrOutput {
	return o
}

func (o CapacitySkuResponsePtrOutput) ToCapacitySkuResponsePtrOutputWithContext(ctx context.Context) CapacitySkuResponsePtrOutput {
	return o
}

func (o CapacitySkuResponsePtrOutput) Elem() CapacitySkuResponseOutput {
	return o.ApplyT(func(v *CapacitySkuResponse) CapacitySkuResponse {
		if v != nil {
			return *v
		}
		var ret CapacitySkuResponse
		return ret
	}).(CapacitySkuResponseOutput)
}

func (o CapacitySkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapacitySkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CapacitySkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapacitySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type DedicatedCapacityAdministrators struct {
	Members []string `pulumi:"members"`
}





type DedicatedCapacityAdministratorsInput interface {
	pulumi.Input

	ToDedicatedCapacityAdministratorsOutput() DedicatedCapacityAdministratorsOutput
	ToDedicatedCapacityAdministratorsOutputWithContext(context.Context) DedicatedCapacityAdministratorsOutput
}

type DedicatedCapacityAdministratorsArgs struct {
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (DedicatedCapacityAdministratorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministrators)(nil)).Elem()
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsOutput() DedicatedCapacityAdministratorsOutput {
	return i.ToDedicatedCapacityAdministratorsOutputWithContext(context.Background())
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsOutput)
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return i.ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Background())
}

func (i DedicatedCapacityAdministratorsArgs) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsOutput).ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx)
}









type DedicatedCapacityAdministratorsPtrInput interface {
	pulumi.Input

	ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput
	ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Context) DedicatedCapacityAdministratorsPtrOutput
}

type dedicatedCapacityAdministratorsPtrType DedicatedCapacityAdministratorsArgs

func DedicatedCapacityAdministratorsPtr(v *DedicatedCapacityAdministratorsArgs) DedicatedCapacityAdministratorsPtrInput {
	return (*dedicatedCapacityAdministratorsPtrType)(v)
}

func (*dedicatedCapacityAdministratorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministrators)(nil)).Elem()
}

func (i *dedicatedCapacityAdministratorsPtrType) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return i.ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Background())
}

func (i *dedicatedCapacityAdministratorsPtrType) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsPtrOutput)
}

type DedicatedCapacityAdministratorsOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministrators)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsOutput() DedicatedCapacityAdministratorsOutput {
	return o
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsOutput {
	return o
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return o.ToDedicatedCapacityAdministratorsPtrOutputWithContext(context.Background())
}

func (o DedicatedCapacityAdministratorsOutput) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DedicatedCapacityAdministrators) *DedicatedCapacityAdministrators {
		return &v
	}).(DedicatedCapacityAdministratorsPtrOutput)
}

func (o DedicatedCapacityAdministratorsOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DedicatedCapacityAdministrators) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type DedicatedCapacityAdministratorsPtrOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministrators)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsPtrOutput) ToDedicatedCapacityAdministratorsPtrOutput() DedicatedCapacityAdministratorsPtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsPtrOutput) ToDedicatedCapacityAdministratorsPtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsPtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsPtrOutput) Elem() DedicatedCapacityAdministratorsOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministrators) DedicatedCapacityAdministrators {
		if v != nil {
			return *v
		}
		var ret DedicatedCapacityAdministrators
		return ret
	}).(DedicatedCapacityAdministratorsOutput)
}

func (o DedicatedCapacityAdministratorsPtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministrators) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

type DedicatedCapacityAdministratorsResponse struct {
	Members []string `pulumi:"members"`
}





type DedicatedCapacityAdministratorsResponseInput interface {
	pulumi.Input

	ToDedicatedCapacityAdministratorsResponseOutput() DedicatedCapacityAdministratorsResponseOutput
	ToDedicatedCapacityAdministratorsResponseOutputWithContext(context.Context) DedicatedCapacityAdministratorsResponseOutput
}

type DedicatedCapacityAdministratorsResponseArgs struct {
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (DedicatedCapacityAdministratorsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministratorsResponse)(nil)).Elem()
}

func (i DedicatedCapacityAdministratorsResponseArgs) ToDedicatedCapacityAdministratorsResponseOutput() DedicatedCapacityAdministratorsResponseOutput {
	return i.ToDedicatedCapacityAdministratorsResponseOutputWithContext(context.Background())
}

func (i DedicatedCapacityAdministratorsResponseArgs) ToDedicatedCapacityAdministratorsResponseOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsResponseOutput)
}

func (i DedicatedCapacityAdministratorsResponseArgs) ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput {
	return i.ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (i DedicatedCapacityAdministratorsResponseArgs) ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsResponseOutput).ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx)
}









type DedicatedCapacityAdministratorsResponsePtrInput interface {
	pulumi.Input

	ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput
	ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(context.Context) DedicatedCapacityAdministratorsResponsePtrOutput
}

type dedicatedCapacityAdministratorsResponsePtrType DedicatedCapacityAdministratorsResponseArgs

func DedicatedCapacityAdministratorsResponsePtr(v *DedicatedCapacityAdministratorsResponseArgs) DedicatedCapacityAdministratorsResponsePtrInput {
	return (*dedicatedCapacityAdministratorsResponsePtrType)(v)
}

func (*dedicatedCapacityAdministratorsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministratorsResponse)(nil)).Elem()
}

func (i *dedicatedCapacityAdministratorsResponsePtrType) ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput {
	return i.ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (i *dedicatedCapacityAdministratorsResponsePtrType) ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCapacityAdministratorsResponsePtrOutput)
}

type DedicatedCapacityAdministratorsResponseOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedCapacityAdministratorsResponse)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsResponseOutput) ToDedicatedCapacityAdministratorsResponseOutput() DedicatedCapacityAdministratorsResponseOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponseOutput) ToDedicatedCapacityAdministratorsResponseOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponseOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponseOutput) ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput {
	return o.ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (o DedicatedCapacityAdministratorsResponseOutput) ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DedicatedCapacityAdministratorsResponse) *DedicatedCapacityAdministratorsResponse {
		return &v
	}).(DedicatedCapacityAdministratorsResponsePtrOutput)
}

func (o DedicatedCapacityAdministratorsResponseOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DedicatedCapacityAdministratorsResponse) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type DedicatedCapacityAdministratorsResponsePtrOutput struct{ *pulumi.OutputState }

func (DedicatedCapacityAdministratorsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCapacityAdministratorsResponse)(nil)).Elem()
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) ToDedicatedCapacityAdministratorsResponsePtrOutput() DedicatedCapacityAdministratorsResponsePtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) ToDedicatedCapacityAdministratorsResponsePtrOutputWithContext(ctx context.Context) DedicatedCapacityAdministratorsResponsePtrOutput {
	return o
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) Elem() DedicatedCapacityAdministratorsResponseOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministratorsResponse) DedicatedCapacityAdministratorsResponse {
		if v != nil {
			return *v
		}
		var ret DedicatedCapacityAdministratorsResponse
		return ret
	}).(DedicatedCapacityAdministratorsResponseOutput)
}

func (o DedicatedCapacityAdministratorsResponsePtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DedicatedCapacityAdministratorsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

type SystemData struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataInput interface {
	pulumi.Input

	ToSystemDataOutput() SystemDataOutput
	ToSystemDataOutputWithContext(context.Context) SystemDataOutput
}

type SystemDataArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemData)(nil)).Elem()
}

func (i SystemDataArgs) ToSystemDataOutput() SystemDataOutput {
	return i.ToSystemDataOutputWithContext(context.Background())
}

func (i SystemDataArgs) ToSystemDataOutputWithContext(ctx context.Context) SystemDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataOutput)
}

func (i SystemDataArgs) ToSystemDataPtrOutput() SystemDataPtrOutput {
	return i.ToSystemDataPtrOutputWithContext(context.Background())
}

func (i SystemDataArgs) ToSystemDataPtrOutputWithContext(ctx context.Context) SystemDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataOutput).ToSystemDataPtrOutputWithContext(ctx)
}









type SystemDataPtrInput interface {
	pulumi.Input

	ToSystemDataPtrOutput() SystemDataPtrOutput
	ToSystemDataPtrOutputWithContext(context.Context) SystemDataPtrOutput
}

type systemDataPtrType SystemDataArgs

func SystemDataPtr(v *SystemDataArgs) SystemDataPtrInput {
	return (*systemDataPtrType)(v)
}

func (*systemDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemData)(nil)).Elem()
}

func (i *systemDataPtrType) ToSystemDataPtrOutput() SystemDataPtrOutput {
	return i.ToSystemDataPtrOutputWithContext(context.Background())
}

func (i *systemDataPtrType) ToSystemDataPtrOutputWithContext(ctx context.Context) SystemDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataPtrOutput)
}

type SystemDataOutput struct{ *pulumi.OutputState }

func (SystemDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemData)(nil)).Elem()
}

func (o SystemDataOutput) ToSystemDataOutput() SystemDataOutput {
	return o
}

func (o SystemDataOutput) ToSystemDataOutputWithContext(ctx context.Context) SystemDataOutput {
	return o
}

func (o SystemDataOutput) ToSystemDataPtrOutput() SystemDataPtrOutput {
	return o.ToSystemDataPtrOutputWithContext(context.Background())
}

func (o SystemDataOutput) ToSystemDataPtrOutputWithContext(ctx context.Context) SystemDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemData) *SystemData {
		return &v
	}).(SystemDataPtrOutput)
}

func (o SystemDataOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemData) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataPtrOutput struct{ *pulumi.OutputState }

func (SystemDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemData)(nil)).Elem()
}

func (o SystemDataPtrOutput) ToSystemDataPtrOutput() SystemDataPtrOutput {
	return o
}

func (o SystemDataPtrOutput) ToSystemDataPtrOutputWithContext(ctx context.Context) SystemDataPtrOutput {
	return o
}

func (o SystemDataPtrOutput) Elem() SystemDataOutput {
	return o.ApplyT(func(v *SystemData) SystemData {
		if v != nil {
			return *v
		}
		var ret SystemData
		return ret
	}).(SystemDataOutput)
}

func (o SystemDataPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataPtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataPtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataPtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataPtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataPtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoScaleVCoreSkuInput)(nil)).Elem(), AutoScaleVCoreSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoScaleVCoreSkuPtrInput)(nil)).Elem(), AutoScaleVCoreSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoScaleVCoreSkuResponseInput)(nil)).Elem(), AutoScaleVCoreSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoScaleVCoreSkuResponsePtrInput)(nil)).Elem(), AutoScaleVCoreSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CapacitySkuInput)(nil)).Elem(), CapacitySkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CapacitySkuPtrInput)(nil)).Elem(), CapacitySkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CapacitySkuResponseInput)(nil)).Elem(), CapacitySkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CapacitySkuResponsePtrInput)(nil)).Elem(), CapacitySkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCapacityAdministratorsInput)(nil)).Elem(), DedicatedCapacityAdministratorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCapacityAdministratorsPtrInput)(nil)).Elem(), DedicatedCapacityAdministratorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCapacityAdministratorsResponseInput)(nil)).Elem(), DedicatedCapacityAdministratorsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedCapacityAdministratorsResponsePtrInput)(nil)).Elem(), DedicatedCapacityAdministratorsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataInput)(nil)).Elem(), SystemDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataPtrInput)(nil)).Elem(), SystemDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterOutputType(AutoScaleVCoreSkuOutput{})
	pulumi.RegisterOutputType(AutoScaleVCoreSkuPtrOutput{})
	pulumi.RegisterOutputType(AutoScaleVCoreSkuResponseOutput{})
	pulumi.RegisterOutputType(AutoScaleVCoreSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(CapacitySkuOutput{})
	pulumi.RegisterOutputType(CapacitySkuPtrOutput{})
	pulumi.RegisterOutputType(CapacitySkuResponseOutput{})
	pulumi.RegisterOutputType(CapacitySkuResponsePtrOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataOutput{})
	pulumi.RegisterOutputType(SystemDataPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
