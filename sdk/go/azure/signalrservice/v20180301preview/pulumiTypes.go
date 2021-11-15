


package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceSku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o ResourceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuResponseInput interface {
	pulumi.Input

	ToResourceSkuResponseOutput() ResourceSkuResponseOutput
	ToResourceSkuResponseOutputWithContext(context.Context) ResourceSkuResponseOutput
}

type ResourceSkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return i.ToResourceSkuResponseOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput)
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput).ToResourceSkuResponsePtrOutputWithContext(ctx)
}









type ResourceSkuResponsePtrInput interface {
	pulumi.Input

	ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput
	ToResourceSkuResponsePtrOutputWithContext(context.Context) ResourceSkuResponsePtrOutput
}

type resourceSkuResponsePtrType ResourceSkuResponseArgs

func ResourceSkuResponsePtr(v *ResourceSkuResponseArgs) ResourceSkuResponsePtrInput {
	return (*resourceSkuResponsePtrType)(v)
}

func (*resourceSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponsePtrOutput)
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSkuResponse) *ResourceSkuResponse {
		return &v
	}).(ResourceSkuResponsePtrOutput)
}

func (o ResourceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SignalRCreateOrUpdateProperties struct {
	HostNamePrefix *string `pulumi:"hostNamePrefix"`
}





type SignalRCreateOrUpdatePropertiesInput interface {
	pulumi.Input

	ToSignalRCreateOrUpdatePropertiesOutput() SignalRCreateOrUpdatePropertiesOutput
	ToSignalRCreateOrUpdatePropertiesOutputWithContext(context.Context) SignalRCreateOrUpdatePropertiesOutput
}

type SignalRCreateOrUpdatePropertiesArgs struct {
	HostNamePrefix pulumi.StringPtrInput `pulumi:"hostNamePrefix"`
}

func (SignalRCreateOrUpdatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRCreateOrUpdateProperties)(nil)).Elem()
}

func (i SignalRCreateOrUpdatePropertiesArgs) ToSignalRCreateOrUpdatePropertiesOutput() SignalRCreateOrUpdatePropertiesOutput {
	return i.ToSignalRCreateOrUpdatePropertiesOutputWithContext(context.Background())
}

func (i SignalRCreateOrUpdatePropertiesArgs) ToSignalRCreateOrUpdatePropertiesOutputWithContext(ctx context.Context) SignalRCreateOrUpdatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCreateOrUpdatePropertiesOutput)
}

func (i SignalRCreateOrUpdatePropertiesArgs) ToSignalRCreateOrUpdatePropertiesPtrOutput() SignalRCreateOrUpdatePropertiesPtrOutput {
	return i.ToSignalRCreateOrUpdatePropertiesPtrOutputWithContext(context.Background())
}

func (i SignalRCreateOrUpdatePropertiesArgs) ToSignalRCreateOrUpdatePropertiesPtrOutputWithContext(ctx context.Context) SignalRCreateOrUpdatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCreateOrUpdatePropertiesOutput).ToSignalRCreateOrUpdatePropertiesPtrOutputWithContext(ctx)
}









type SignalRCreateOrUpdatePropertiesPtrInput interface {
	pulumi.Input

	ToSignalRCreateOrUpdatePropertiesPtrOutput() SignalRCreateOrUpdatePropertiesPtrOutput
	ToSignalRCreateOrUpdatePropertiesPtrOutputWithContext(context.Context) SignalRCreateOrUpdatePropertiesPtrOutput
}

type signalRCreateOrUpdatePropertiesPtrType SignalRCreateOrUpdatePropertiesArgs

func SignalRCreateOrUpdatePropertiesPtr(v *SignalRCreateOrUpdatePropertiesArgs) SignalRCreateOrUpdatePropertiesPtrInput {
	return (*signalRCreateOrUpdatePropertiesPtrType)(v)
}

func (*signalRCreateOrUpdatePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCreateOrUpdateProperties)(nil)).Elem()
}

func (i *signalRCreateOrUpdatePropertiesPtrType) ToSignalRCreateOrUpdatePropertiesPtrOutput() SignalRCreateOrUpdatePropertiesPtrOutput {
	return i.ToSignalRCreateOrUpdatePropertiesPtrOutputWithContext(context.Background())
}

func (i *signalRCreateOrUpdatePropertiesPtrType) ToSignalRCreateOrUpdatePropertiesPtrOutputWithContext(ctx context.Context) SignalRCreateOrUpdatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCreateOrUpdatePropertiesPtrOutput)
}

type SignalRCreateOrUpdatePropertiesOutput struct{ *pulumi.OutputState }

func (SignalRCreateOrUpdatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRCreateOrUpdateProperties)(nil)).Elem()
}

func (o SignalRCreateOrUpdatePropertiesOutput) ToSignalRCreateOrUpdatePropertiesOutput() SignalRCreateOrUpdatePropertiesOutput {
	return o
}

func (o SignalRCreateOrUpdatePropertiesOutput) ToSignalRCreateOrUpdatePropertiesOutputWithContext(ctx context.Context) SignalRCreateOrUpdatePropertiesOutput {
	return o
}

func (o SignalRCreateOrUpdatePropertiesOutput) ToSignalRCreateOrUpdatePropertiesPtrOutput() SignalRCreateOrUpdatePropertiesPtrOutput {
	return o.ToSignalRCreateOrUpdatePropertiesPtrOutputWithContext(context.Background())
}

func (o SignalRCreateOrUpdatePropertiesOutput) ToSignalRCreateOrUpdatePropertiesPtrOutputWithContext(ctx context.Context) SignalRCreateOrUpdatePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignalRCreateOrUpdateProperties) *SignalRCreateOrUpdateProperties {
		return &v
	}).(SignalRCreateOrUpdatePropertiesPtrOutput)
}

func (o SignalRCreateOrUpdatePropertiesOutput) HostNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SignalRCreateOrUpdateProperties) *string { return v.HostNamePrefix }).(pulumi.StringPtrOutput)
}

type SignalRCreateOrUpdatePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SignalRCreateOrUpdatePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCreateOrUpdateProperties)(nil)).Elem()
}

func (o SignalRCreateOrUpdatePropertiesPtrOutput) ToSignalRCreateOrUpdatePropertiesPtrOutput() SignalRCreateOrUpdatePropertiesPtrOutput {
	return o
}

func (o SignalRCreateOrUpdatePropertiesPtrOutput) ToSignalRCreateOrUpdatePropertiesPtrOutputWithContext(ctx context.Context) SignalRCreateOrUpdatePropertiesPtrOutput {
	return o
}

func (o SignalRCreateOrUpdatePropertiesPtrOutput) Elem() SignalRCreateOrUpdatePropertiesOutput {
	return o.ApplyT(func(v *SignalRCreateOrUpdateProperties) SignalRCreateOrUpdateProperties {
		if v != nil {
			return *v
		}
		var ret SignalRCreateOrUpdateProperties
		return ret
	}).(SignalRCreateOrUpdatePropertiesOutput)
}

func (o SignalRCreateOrUpdatePropertiesPtrOutput) HostNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalRCreateOrUpdateProperties) *string {
		if v == nil {
			return nil
		}
		return v.HostNamePrefix
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SignalRCreateOrUpdatePropertiesOutput{})
	pulumi.RegisterOutputType(SignalRCreateOrUpdatePropertiesPtrOutput{})
}
