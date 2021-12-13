


package powerbidedicated

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

func (o AutoScaleVCoreSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o AutoScaleVCoreSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleVCoreSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o AutoScaleVCoreSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type AutoScaleVCoreSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
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

func (o AutoScaleVCoreSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o AutoScaleVCoreSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScaleVCoreSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AutoScaleVCoreSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleVCoreSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o CapacitySkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CapacitySku) string { return v.Name }).(pulumi.StringOutput)
}

func (o CapacitySkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapacitySku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type CapacitySkuResponse struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
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

func (o CapacitySkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CapacitySkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CapacitySkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapacitySkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(AutoScaleVCoreSkuOutput{})
	pulumi.RegisterOutputType(AutoScaleVCoreSkuResponseOutput{})
	pulumi.RegisterOutputType(CapacitySkuOutput{})
	pulumi.RegisterOutputType(CapacitySkuResponseOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(DedicatedCapacityAdministratorsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataOutput{})
	pulumi.RegisterOutputType(SystemDataPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
