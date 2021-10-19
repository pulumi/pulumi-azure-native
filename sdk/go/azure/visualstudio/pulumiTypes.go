


package visualstudio

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExtensionResourcePlan struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}





type ExtensionResourcePlanInput interface {
	pulumi.Input

	ToExtensionResourcePlanOutput() ExtensionResourcePlanOutput
	ToExtensionResourcePlanOutputWithContext(context.Context) ExtensionResourcePlanOutput
}

type ExtensionResourcePlanArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (ExtensionResourcePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionResourcePlan)(nil)).Elem()
}

func (i ExtensionResourcePlanArgs) ToExtensionResourcePlanOutput() ExtensionResourcePlanOutput {
	return i.ToExtensionResourcePlanOutputWithContext(context.Background())
}

func (i ExtensionResourcePlanArgs) ToExtensionResourcePlanOutputWithContext(ctx context.Context) ExtensionResourcePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionResourcePlanOutput)
}

func (i ExtensionResourcePlanArgs) ToExtensionResourcePlanPtrOutput() ExtensionResourcePlanPtrOutput {
	return i.ToExtensionResourcePlanPtrOutputWithContext(context.Background())
}

func (i ExtensionResourcePlanArgs) ToExtensionResourcePlanPtrOutputWithContext(ctx context.Context) ExtensionResourcePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionResourcePlanOutput).ToExtensionResourcePlanPtrOutputWithContext(ctx)
}









type ExtensionResourcePlanPtrInput interface {
	pulumi.Input

	ToExtensionResourcePlanPtrOutput() ExtensionResourcePlanPtrOutput
	ToExtensionResourcePlanPtrOutputWithContext(context.Context) ExtensionResourcePlanPtrOutput
}

type extensionResourcePlanPtrType ExtensionResourcePlanArgs

func ExtensionResourcePlanPtr(v *ExtensionResourcePlanArgs) ExtensionResourcePlanPtrInput {
	return (*extensionResourcePlanPtrType)(v)
}

func (*extensionResourcePlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionResourcePlan)(nil)).Elem()
}

func (i *extensionResourcePlanPtrType) ToExtensionResourcePlanPtrOutput() ExtensionResourcePlanPtrOutput {
	return i.ToExtensionResourcePlanPtrOutputWithContext(context.Background())
}

func (i *extensionResourcePlanPtrType) ToExtensionResourcePlanPtrOutputWithContext(ctx context.Context) ExtensionResourcePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionResourcePlanPtrOutput)
}

type ExtensionResourcePlanOutput struct{ *pulumi.OutputState }

func (ExtensionResourcePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionResourcePlan)(nil)).Elem()
}

func (o ExtensionResourcePlanOutput) ToExtensionResourcePlanOutput() ExtensionResourcePlanOutput {
	return o
}

func (o ExtensionResourcePlanOutput) ToExtensionResourcePlanOutputWithContext(ctx context.Context) ExtensionResourcePlanOutput {
	return o
}

func (o ExtensionResourcePlanOutput) ToExtensionResourcePlanPtrOutput() ExtensionResourcePlanPtrOutput {
	return o.ToExtensionResourcePlanPtrOutputWithContext(context.Background())
}

func (o ExtensionResourcePlanOutput) ToExtensionResourcePlanPtrOutputWithContext(ctx context.Context) ExtensionResourcePlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtensionResourcePlan) *ExtensionResourcePlan {
		return &v
	}).(ExtensionResourcePlanPtrOutput)
}

func (o ExtensionResourcePlanOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResourcePlan) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResourcePlan) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResourcePlan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResourcePlan) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResourcePlan) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ExtensionResourcePlanPtrOutput struct{ *pulumi.OutputState }

func (ExtensionResourcePlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionResourcePlan)(nil)).Elem()
}

func (o ExtensionResourcePlanPtrOutput) ToExtensionResourcePlanPtrOutput() ExtensionResourcePlanPtrOutput {
	return o
}

func (o ExtensionResourcePlanPtrOutput) ToExtensionResourcePlanPtrOutputWithContext(ctx context.Context) ExtensionResourcePlanPtrOutput {
	return o
}

func (o ExtensionResourcePlanPtrOutput) Elem() ExtensionResourcePlanOutput {
	return o.ApplyT(func(v *ExtensionResourcePlan) ExtensionResourcePlan {
		if v != nil {
			return *v
		}
		var ret ExtensionResourcePlan
		return ret
	}).(ExtensionResourcePlanOutput)
}

func (o ExtensionResourcePlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResourcePlan) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResourcePlan) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResourcePlan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResourcePlan) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResourcePlan) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ExtensionResourcePlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}





type ExtensionResourcePlanResponseInput interface {
	pulumi.Input

	ToExtensionResourcePlanResponseOutput() ExtensionResourcePlanResponseOutput
	ToExtensionResourcePlanResponseOutputWithContext(context.Context) ExtensionResourcePlanResponseOutput
}

type ExtensionResourcePlanResponseArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (ExtensionResourcePlanResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionResourcePlanResponse)(nil)).Elem()
}

func (i ExtensionResourcePlanResponseArgs) ToExtensionResourcePlanResponseOutput() ExtensionResourcePlanResponseOutput {
	return i.ToExtensionResourcePlanResponseOutputWithContext(context.Background())
}

func (i ExtensionResourcePlanResponseArgs) ToExtensionResourcePlanResponseOutputWithContext(ctx context.Context) ExtensionResourcePlanResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionResourcePlanResponseOutput)
}

func (i ExtensionResourcePlanResponseArgs) ToExtensionResourcePlanResponsePtrOutput() ExtensionResourcePlanResponsePtrOutput {
	return i.ToExtensionResourcePlanResponsePtrOutputWithContext(context.Background())
}

func (i ExtensionResourcePlanResponseArgs) ToExtensionResourcePlanResponsePtrOutputWithContext(ctx context.Context) ExtensionResourcePlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionResourcePlanResponseOutput).ToExtensionResourcePlanResponsePtrOutputWithContext(ctx)
}









type ExtensionResourcePlanResponsePtrInput interface {
	pulumi.Input

	ToExtensionResourcePlanResponsePtrOutput() ExtensionResourcePlanResponsePtrOutput
	ToExtensionResourcePlanResponsePtrOutputWithContext(context.Context) ExtensionResourcePlanResponsePtrOutput
}

type extensionResourcePlanResponsePtrType ExtensionResourcePlanResponseArgs

func ExtensionResourcePlanResponsePtr(v *ExtensionResourcePlanResponseArgs) ExtensionResourcePlanResponsePtrInput {
	return (*extensionResourcePlanResponsePtrType)(v)
}

func (*extensionResourcePlanResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionResourcePlanResponse)(nil)).Elem()
}

func (i *extensionResourcePlanResponsePtrType) ToExtensionResourcePlanResponsePtrOutput() ExtensionResourcePlanResponsePtrOutput {
	return i.ToExtensionResourcePlanResponsePtrOutputWithContext(context.Background())
}

func (i *extensionResourcePlanResponsePtrType) ToExtensionResourcePlanResponsePtrOutputWithContext(ctx context.Context) ExtensionResourcePlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionResourcePlanResponsePtrOutput)
}

type ExtensionResourcePlanResponseOutput struct{ *pulumi.OutputState }

func (ExtensionResourcePlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionResourcePlanResponse)(nil)).Elem()
}

func (o ExtensionResourcePlanResponseOutput) ToExtensionResourcePlanResponseOutput() ExtensionResourcePlanResponseOutput {
	return o
}

func (o ExtensionResourcePlanResponseOutput) ToExtensionResourcePlanResponseOutputWithContext(ctx context.Context) ExtensionResourcePlanResponseOutput {
	return o
}

func (o ExtensionResourcePlanResponseOutput) ToExtensionResourcePlanResponsePtrOutput() ExtensionResourcePlanResponsePtrOutput {
	return o.ToExtensionResourcePlanResponsePtrOutputWithContext(context.Background())
}

func (o ExtensionResourcePlanResponseOutput) ToExtensionResourcePlanResponsePtrOutputWithContext(ctx context.Context) ExtensionResourcePlanResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtensionResourcePlanResponse) *ExtensionResourcePlanResponse {
		return &v
	}).(ExtensionResourcePlanResponsePtrOutput)
}

func (o ExtensionResourcePlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResourcePlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResourcePlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResourcePlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResourcePlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResourcePlanResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ExtensionResourcePlanResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtensionResourcePlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionResourcePlanResponse)(nil)).Elem()
}

func (o ExtensionResourcePlanResponsePtrOutput) ToExtensionResourcePlanResponsePtrOutput() ExtensionResourcePlanResponsePtrOutput {
	return o
}

func (o ExtensionResourcePlanResponsePtrOutput) ToExtensionResourcePlanResponsePtrOutputWithContext(ctx context.Context) ExtensionResourcePlanResponsePtrOutput {
	return o
}

func (o ExtensionResourcePlanResponsePtrOutput) Elem() ExtensionResourcePlanResponseOutput {
	return o.ApplyT(func(v *ExtensionResourcePlanResponse) ExtensionResourcePlanResponse {
		if v != nil {
			return *v
		}
		var ret ExtensionResourcePlanResponse
		return ret
	}).(ExtensionResourcePlanResponseOutput)
}

func (o ExtensionResourcePlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResourcePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResourcePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResourcePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResourcePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResourcePlanResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResourcePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionResourcePlanInput)(nil)).Elem(), ExtensionResourcePlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionResourcePlanPtrInput)(nil)).Elem(), ExtensionResourcePlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionResourcePlanResponseInput)(nil)).Elem(), ExtensionResourcePlanResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionResourcePlanResponsePtrInput)(nil)).Elem(), ExtensionResourcePlanResponseArgs{})
	pulumi.RegisterOutputType(ExtensionResourcePlanOutput{})
	pulumi.RegisterOutputType(ExtensionResourcePlanPtrOutput{})
	pulumi.RegisterOutputType(ExtensionResourcePlanResponseOutput{})
	pulumi.RegisterOutputType(ExtensionResourcePlanResponsePtrOutput{})
}
