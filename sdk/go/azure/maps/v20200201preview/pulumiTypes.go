


package v20200201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreatorPropertiesResponse struct {
	ProvisioningState *string `pulumi:"provisioningState"`
}





type CreatorPropertiesResponseInput interface {
	pulumi.Input

	ToCreatorPropertiesResponseOutput() CreatorPropertiesResponseOutput
	ToCreatorPropertiesResponseOutputWithContext(context.Context) CreatorPropertiesResponseOutput
}

type CreatorPropertiesResponseArgs struct {
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (CreatorPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatorPropertiesResponse)(nil)).Elem()
}

func (i CreatorPropertiesResponseArgs) ToCreatorPropertiesResponseOutput() CreatorPropertiesResponseOutput {
	return i.ToCreatorPropertiesResponseOutputWithContext(context.Background())
}

func (i CreatorPropertiesResponseArgs) ToCreatorPropertiesResponseOutputWithContext(ctx context.Context) CreatorPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatorPropertiesResponseOutput)
}

func (i CreatorPropertiesResponseArgs) ToCreatorPropertiesResponsePtrOutput() CreatorPropertiesResponsePtrOutput {
	return i.ToCreatorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CreatorPropertiesResponseArgs) ToCreatorPropertiesResponsePtrOutputWithContext(ctx context.Context) CreatorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatorPropertiesResponseOutput).ToCreatorPropertiesResponsePtrOutputWithContext(ctx)
}









type CreatorPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCreatorPropertiesResponsePtrOutput() CreatorPropertiesResponsePtrOutput
	ToCreatorPropertiesResponsePtrOutputWithContext(context.Context) CreatorPropertiesResponsePtrOutput
}

type creatorPropertiesResponsePtrType CreatorPropertiesResponseArgs

func CreatorPropertiesResponsePtr(v *CreatorPropertiesResponseArgs) CreatorPropertiesResponsePtrInput {
	return (*creatorPropertiesResponsePtrType)(v)
}

func (*creatorPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreatorPropertiesResponse)(nil)).Elem()
}

func (i *creatorPropertiesResponsePtrType) ToCreatorPropertiesResponsePtrOutput() CreatorPropertiesResponsePtrOutput {
	return i.ToCreatorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *creatorPropertiesResponsePtrType) ToCreatorPropertiesResponsePtrOutputWithContext(ctx context.Context) CreatorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatorPropertiesResponsePtrOutput)
}

type CreatorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CreatorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatorPropertiesResponse)(nil)).Elem()
}

func (o CreatorPropertiesResponseOutput) ToCreatorPropertiesResponseOutput() CreatorPropertiesResponseOutput {
	return o
}

func (o CreatorPropertiesResponseOutput) ToCreatorPropertiesResponseOutputWithContext(ctx context.Context) CreatorPropertiesResponseOutput {
	return o
}

func (o CreatorPropertiesResponseOutput) ToCreatorPropertiesResponsePtrOutput() CreatorPropertiesResponsePtrOutput {
	return o.ToCreatorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CreatorPropertiesResponseOutput) ToCreatorPropertiesResponsePtrOutputWithContext(ctx context.Context) CreatorPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreatorPropertiesResponse) *CreatorPropertiesResponse {
		return &v
	}).(CreatorPropertiesResponsePtrOutput)
}

func (o CreatorPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreatorPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type CreatorPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CreatorPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreatorPropertiesResponse)(nil)).Elem()
}

func (o CreatorPropertiesResponsePtrOutput) ToCreatorPropertiesResponsePtrOutput() CreatorPropertiesResponsePtrOutput {
	return o
}

func (o CreatorPropertiesResponsePtrOutput) ToCreatorPropertiesResponsePtrOutputWithContext(ctx context.Context) CreatorPropertiesResponsePtrOutput {
	return o
}

func (o CreatorPropertiesResponsePtrOutput) Elem() CreatorPropertiesResponseOutput {
	return o.ApplyT(func(v *CreatorPropertiesResponse) CreatorPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CreatorPropertiesResponse
		return ret
	}).(CreatorPropertiesResponseOutput)
}

func (o CreatorPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreatorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

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

type PrivateAtlasPropertiesResponse struct {
	ProvisioningState *string `pulumi:"provisioningState"`
}





type PrivateAtlasPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateAtlasPropertiesResponseOutput() PrivateAtlasPropertiesResponseOutput
	ToPrivateAtlasPropertiesResponseOutputWithContext(context.Context) PrivateAtlasPropertiesResponseOutput
}

type PrivateAtlasPropertiesResponseArgs struct {
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (PrivateAtlasPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateAtlasPropertiesResponse)(nil)).Elem()
}

func (i PrivateAtlasPropertiesResponseArgs) ToPrivateAtlasPropertiesResponseOutput() PrivateAtlasPropertiesResponseOutput {
	return i.ToPrivateAtlasPropertiesResponseOutputWithContext(context.Background())
}

func (i PrivateAtlasPropertiesResponseArgs) ToPrivateAtlasPropertiesResponseOutputWithContext(ctx context.Context) PrivateAtlasPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateAtlasPropertiesResponseOutput)
}

func (i PrivateAtlasPropertiesResponseArgs) ToPrivateAtlasPropertiesResponsePtrOutput() PrivateAtlasPropertiesResponsePtrOutput {
	return i.ToPrivateAtlasPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PrivateAtlasPropertiesResponseArgs) ToPrivateAtlasPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateAtlasPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateAtlasPropertiesResponseOutput).ToPrivateAtlasPropertiesResponsePtrOutputWithContext(ctx)
}









type PrivateAtlasPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPrivateAtlasPropertiesResponsePtrOutput() PrivateAtlasPropertiesResponsePtrOutput
	ToPrivateAtlasPropertiesResponsePtrOutputWithContext(context.Context) PrivateAtlasPropertiesResponsePtrOutput
}

type privateAtlasPropertiesResponsePtrType PrivateAtlasPropertiesResponseArgs

func PrivateAtlasPropertiesResponsePtr(v *PrivateAtlasPropertiesResponseArgs) PrivateAtlasPropertiesResponsePtrInput {
	return (*privateAtlasPropertiesResponsePtrType)(v)
}

func (*privateAtlasPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateAtlasPropertiesResponse)(nil)).Elem()
}

func (i *privateAtlasPropertiesResponsePtrType) ToPrivateAtlasPropertiesResponsePtrOutput() PrivateAtlasPropertiesResponsePtrOutput {
	return i.ToPrivateAtlasPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *privateAtlasPropertiesResponsePtrType) ToPrivateAtlasPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateAtlasPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateAtlasPropertiesResponsePtrOutput)
}

type PrivateAtlasPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateAtlasPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateAtlasPropertiesResponse)(nil)).Elem()
}

func (o PrivateAtlasPropertiesResponseOutput) ToPrivateAtlasPropertiesResponseOutput() PrivateAtlasPropertiesResponseOutput {
	return o
}

func (o PrivateAtlasPropertiesResponseOutput) ToPrivateAtlasPropertiesResponseOutputWithContext(ctx context.Context) PrivateAtlasPropertiesResponseOutput {
	return o
}

func (o PrivateAtlasPropertiesResponseOutput) ToPrivateAtlasPropertiesResponsePtrOutput() PrivateAtlasPropertiesResponsePtrOutput {
	return o.ToPrivateAtlasPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PrivateAtlasPropertiesResponseOutput) ToPrivateAtlasPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateAtlasPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateAtlasPropertiesResponse) *PrivateAtlasPropertiesResponse {
		return &v
	}).(PrivateAtlasPropertiesResponsePtrOutput)
}

func (o PrivateAtlasPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateAtlasPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type PrivateAtlasPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateAtlasPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateAtlasPropertiesResponse)(nil)).Elem()
}

func (o PrivateAtlasPropertiesResponsePtrOutput) ToPrivateAtlasPropertiesResponsePtrOutput() PrivateAtlasPropertiesResponsePtrOutput {
	return o
}

func (o PrivateAtlasPropertiesResponsePtrOutput) ToPrivateAtlasPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateAtlasPropertiesResponsePtrOutput {
	return o
}

func (o PrivateAtlasPropertiesResponsePtrOutput) Elem() PrivateAtlasPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateAtlasPropertiesResponse) PrivateAtlasPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PrivateAtlasPropertiesResponse
		return ret
	}).(PrivateAtlasPropertiesResponseOutput)
}

func (o PrivateAtlasPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateAtlasPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
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
	pulumi.RegisterInputType(reflect.TypeOf((*CreatorPropertiesResponseInput)(nil)).Elem(), CreatorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreatorPropertiesResponsePtrInput)(nil)).Elem(), CreatorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MapsAccountPropertiesResponseInput)(nil)).Elem(), MapsAccountPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MapsAccountPropertiesResponsePtrInput)(nil)).Elem(), MapsAccountPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateAtlasPropertiesResponseInput)(nil)).Elem(), PrivateAtlasPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateAtlasPropertiesResponsePtrInput)(nil)).Elem(), PrivateAtlasPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterOutputType(CreatorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CreatorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateAtlasPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateAtlasPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
