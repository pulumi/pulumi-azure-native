


package v20170101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MapsAccountPropertiesResponse struct {
	XMsClientId *string `pulumi:"xMsClientId"`
}





type MapsAccountPropertiesResponseInput interface {
	pulumi.Input

	ToMapsAccountPropertiesResponseOutput() MapsAccountPropertiesResponseOutput
	ToMapsAccountPropertiesResponseOutputWithContext(context.Context) MapsAccountPropertiesResponseOutput
}

type MapsAccountPropertiesResponseArgs struct {
	XMsClientId pulumi.StringPtrInput `pulumi:"xMsClientId"`
}

func (MapsAccountPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MapsAccountPropertiesResponse)(nil)).Elem()
}

func (i MapsAccountPropertiesResponseArgs) ToMapsAccountPropertiesResponseOutput() MapsAccountPropertiesResponseOutput {
	return i.ToMapsAccountPropertiesResponseOutputWithContext(context.Background())
}

func (i MapsAccountPropertiesResponseArgs) ToMapsAccountPropertiesResponseOutputWithContext(ctx context.Context) MapsAccountPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MapsAccountPropertiesResponseOutput)
}

func (i MapsAccountPropertiesResponseArgs) ToMapsAccountPropertiesResponsePtrOutput() MapsAccountPropertiesResponsePtrOutput {
	return i.ToMapsAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MapsAccountPropertiesResponseArgs) ToMapsAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) MapsAccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MapsAccountPropertiesResponseOutput).ToMapsAccountPropertiesResponsePtrOutputWithContext(ctx)
}









type MapsAccountPropertiesResponsePtrInput interface {
	pulumi.Input

	ToMapsAccountPropertiesResponsePtrOutput() MapsAccountPropertiesResponsePtrOutput
	ToMapsAccountPropertiesResponsePtrOutputWithContext(context.Context) MapsAccountPropertiesResponsePtrOutput
}

type mapsAccountPropertiesResponsePtrType MapsAccountPropertiesResponseArgs

func MapsAccountPropertiesResponsePtr(v *MapsAccountPropertiesResponseArgs) MapsAccountPropertiesResponsePtrInput {
	return (*mapsAccountPropertiesResponsePtrType)(v)
}

func (*mapsAccountPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MapsAccountPropertiesResponse)(nil)).Elem()
}

func (i *mapsAccountPropertiesResponsePtrType) ToMapsAccountPropertiesResponsePtrOutput() MapsAccountPropertiesResponsePtrOutput {
	return i.ToMapsAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *mapsAccountPropertiesResponsePtrType) ToMapsAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) MapsAccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MapsAccountPropertiesResponsePtrOutput)
}

type MapsAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MapsAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MapsAccountPropertiesResponse)(nil)).Elem()
}

func (o MapsAccountPropertiesResponseOutput) ToMapsAccountPropertiesResponseOutput() MapsAccountPropertiesResponseOutput {
	return o
}

func (o MapsAccountPropertiesResponseOutput) ToMapsAccountPropertiesResponseOutputWithContext(ctx context.Context) MapsAccountPropertiesResponseOutput {
	return o
}

func (o MapsAccountPropertiesResponseOutput) ToMapsAccountPropertiesResponsePtrOutput() MapsAccountPropertiesResponsePtrOutput {
	return o.ToMapsAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MapsAccountPropertiesResponseOutput) ToMapsAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) MapsAccountPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MapsAccountPropertiesResponse) *MapsAccountPropertiesResponse {
		return &v
	}).(MapsAccountPropertiesResponsePtrOutput)
}

func (o MapsAccountPropertiesResponseOutput) XMsClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) *string { return v.XMsClientId }).(pulumi.StringPtrOutput)
}

type MapsAccountPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MapsAccountPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MapsAccountPropertiesResponse)(nil)).Elem()
}

func (o MapsAccountPropertiesResponsePtrOutput) ToMapsAccountPropertiesResponsePtrOutput() MapsAccountPropertiesResponsePtrOutput {
	return o
}

func (o MapsAccountPropertiesResponsePtrOutput) ToMapsAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) MapsAccountPropertiesResponsePtrOutput {
	return o
}

func (o MapsAccountPropertiesResponsePtrOutput) Elem() MapsAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *MapsAccountPropertiesResponse) MapsAccountPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MapsAccountPropertiesResponse
		return ret
	}).(MapsAccountPropertiesResponseOutput)
}

func (o MapsAccountPropertiesResponsePtrOutput) XMsClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MapsAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.XMsClientId
	}).(pulumi.StringPtrOutput)
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
	Tier string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
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

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
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

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MapsAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
