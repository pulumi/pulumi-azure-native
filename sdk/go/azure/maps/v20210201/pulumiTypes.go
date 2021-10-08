


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreatorProperties struct {
	StorageUnits int `pulumi:"storageUnits"`
}





type CreatorPropertiesInput interface {
	pulumi.Input

	ToCreatorPropertiesOutput() CreatorPropertiesOutput
	ToCreatorPropertiesOutputWithContext(context.Context) CreatorPropertiesOutput
}

type CreatorPropertiesArgs struct {
	StorageUnits pulumi.IntInput `pulumi:"storageUnits"`
}

func (CreatorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatorProperties)(nil)).Elem()
}

func (i CreatorPropertiesArgs) ToCreatorPropertiesOutput() CreatorPropertiesOutput {
	return i.ToCreatorPropertiesOutputWithContext(context.Background())
}

func (i CreatorPropertiesArgs) ToCreatorPropertiesOutputWithContext(ctx context.Context) CreatorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatorPropertiesOutput)
}

func (i CreatorPropertiesArgs) ToCreatorPropertiesPtrOutput() CreatorPropertiesPtrOutput {
	return i.ToCreatorPropertiesPtrOutputWithContext(context.Background())
}

func (i CreatorPropertiesArgs) ToCreatorPropertiesPtrOutputWithContext(ctx context.Context) CreatorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatorPropertiesOutput).ToCreatorPropertiesPtrOutputWithContext(ctx)
}









type CreatorPropertiesPtrInput interface {
	pulumi.Input

	ToCreatorPropertiesPtrOutput() CreatorPropertiesPtrOutput
	ToCreatorPropertiesPtrOutputWithContext(context.Context) CreatorPropertiesPtrOutput
}

type creatorPropertiesPtrType CreatorPropertiesArgs

func CreatorPropertiesPtr(v *CreatorPropertiesArgs) CreatorPropertiesPtrInput {
	return (*creatorPropertiesPtrType)(v)
}

func (*creatorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreatorProperties)(nil)).Elem()
}

func (i *creatorPropertiesPtrType) ToCreatorPropertiesPtrOutput() CreatorPropertiesPtrOutput {
	return i.ToCreatorPropertiesPtrOutputWithContext(context.Background())
}

func (i *creatorPropertiesPtrType) ToCreatorPropertiesPtrOutputWithContext(ctx context.Context) CreatorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatorPropertiesPtrOutput)
}

type CreatorPropertiesOutput struct{ *pulumi.OutputState }

func (CreatorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatorProperties)(nil)).Elem()
}

func (o CreatorPropertiesOutput) ToCreatorPropertiesOutput() CreatorPropertiesOutput {
	return o
}

func (o CreatorPropertiesOutput) ToCreatorPropertiesOutputWithContext(ctx context.Context) CreatorPropertiesOutput {
	return o
}

func (o CreatorPropertiesOutput) ToCreatorPropertiesPtrOutput() CreatorPropertiesPtrOutput {
	return o.ToCreatorPropertiesPtrOutputWithContext(context.Background())
}

func (o CreatorPropertiesOutput) ToCreatorPropertiesPtrOutputWithContext(ctx context.Context) CreatorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreatorProperties) *CreatorProperties {
		return &v
	}).(CreatorPropertiesPtrOutput)
}

func (o CreatorPropertiesOutput) StorageUnits() pulumi.IntOutput {
	return o.ApplyT(func(v CreatorProperties) int { return v.StorageUnits }).(pulumi.IntOutput)
}

type CreatorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CreatorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreatorProperties)(nil)).Elem()
}

func (o CreatorPropertiesPtrOutput) ToCreatorPropertiesPtrOutput() CreatorPropertiesPtrOutput {
	return o
}

func (o CreatorPropertiesPtrOutput) ToCreatorPropertiesPtrOutputWithContext(ctx context.Context) CreatorPropertiesPtrOutput {
	return o
}

func (o CreatorPropertiesPtrOutput) Elem() CreatorPropertiesOutput {
	return o.ApplyT(func(v *CreatorProperties) CreatorProperties {
		if v != nil {
			return *v
		}
		var ret CreatorProperties
		return ret
	}).(CreatorPropertiesOutput)
}

func (o CreatorPropertiesPtrOutput) StorageUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CreatorProperties) *int {
		if v == nil {
			return nil
		}
		return &v.StorageUnits
	}).(pulumi.IntPtrOutput)
}

type CreatorPropertiesResponse struct {
	ProvisioningState string `pulumi:"provisioningState"`
	StorageUnits      int    `pulumi:"storageUnits"`
}





type CreatorPropertiesResponseInput interface {
	pulumi.Input

	ToCreatorPropertiesResponseOutput() CreatorPropertiesResponseOutput
	ToCreatorPropertiesResponseOutputWithContext(context.Context) CreatorPropertiesResponseOutput
}

type CreatorPropertiesResponseArgs struct {
	ProvisioningState pulumi.StringInput `pulumi:"provisioningState"`
	StorageUnits      pulumi.IntInput    `pulumi:"storageUnits"`
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

func (o CreatorPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v CreatorPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CreatorPropertiesResponseOutput) StorageUnits() pulumi.IntOutput {
	return o.ApplyT(func(v CreatorPropertiesResponse) int { return v.StorageUnits }).(pulumi.IntOutput)
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
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o CreatorPropertiesResponsePtrOutput) StorageUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CreatorPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.StorageUnits
	}).(pulumi.IntPtrOutput)
}

type MapsAccountProperties struct {
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
}





type MapsAccountPropertiesInput interface {
	pulumi.Input

	ToMapsAccountPropertiesOutput() MapsAccountPropertiesOutput
	ToMapsAccountPropertiesOutputWithContext(context.Context) MapsAccountPropertiesOutput
}

type MapsAccountPropertiesArgs struct {
	DisableLocalAuth pulumi.BoolPtrInput `pulumi:"disableLocalAuth"`
}

func (MapsAccountPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MapsAccountProperties)(nil)).Elem()
}

func (i MapsAccountPropertiesArgs) ToMapsAccountPropertiesOutput() MapsAccountPropertiesOutput {
	return i.ToMapsAccountPropertiesOutputWithContext(context.Background())
}

func (i MapsAccountPropertiesArgs) ToMapsAccountPropertiesOutputWithContext(ctx context.Context) MapsAccountPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MapsAccountPropertiesOutput)
}

func (i MapsAccountPropertiesArgs) ToMapsAccountPropertiesPtrOutput() MapsAccountPropertiesPtrOutput {
	return i.ToMapsAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i MapsAccountPropertiesArgs) ToMapsAccountPropertiesPtrOutputWithContext(ctx context.Context) MapsAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MapsAccountPropertiesOutput).ToMapsAccountPropertiesPtrOutputWithContext(ctx)
}









type MapsAccountPropertiesPtrInput interface {
	pulumi.Input

	ToMapsAccountPropertiesPtrOutput() MapsAccountPropertiesPtrOutput
	ToMapsAccountPropertiesPtrOutputWithContext(context.Context) MapsAccountPropertiesPtrOutput
}

type mapsAccountPropertiesPtrType MapsAccountPropertiesArgs

func MapsAccountPropertiesPtr(v *MapsAccountPropertiesArgs) MapsAccountPropertiesPtrInput {
	return (*mapsAccountPropertiesPtrType)(v)
}

func (*mapsAccountPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MapsAccountProperties)(nil)).Elem()
}

func (i *mapsAccountPropertiesPtrType) ToMapsAccountPropertiesPtrOutput() MapsAccountPropertiesPtrOutput {
	return i.ToMapsAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i *mapsAccountPropertiesPtrType) ToMapsAccountPropertiesPtrOutputWithContext(ctx context.Context) MapsAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MapsAccountPropertiesPtrOutput)
}

type MapsAccountPropertiesOutput struct{ *pulumi.OutputState }

func (MapsAccountPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MapsAccountProperties)(nil)).Elem()
}

func (o MapsAccountPropertiesOutput) ToMapsAccountPropertiesOutput() MapsAccountPropertiesOutput {
	return o
}

func (o MapsAccountPropertiesOutput) ToMapsAccountPropertiesOutputWithContext(ctx context.Context) MapsAccountPropertiesOutput {
	return o
}

func (o MapsAccountPropertiesOutput) ToMapsAccountPropertiesPtrOutput() MapsAccountPropertiesPtrOutput {
	return o.ToMapsAccountPropertiesPtrOutputWithContext(context.Background())
}

func (o MapsAccountPropertiesOutput) ToMapsAccountPropertiesPtrOutputWithContext(ctx context.Context) MapsAccountPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MapsAccountProperties) *MapsAccountProperties {
		return &v
	}).(MapsAccountPropertiesPtrOutput)
}

func (o MapsAccountPropertiesOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MapsAccountProperties) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

type MapsAccountPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MapsAccountPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MapsAccountProperties)(nil)).Elem()
}

func (o MapsAccountPropertiesPtrOutput) ToMapsAccountPropertiesPtrOutput() MapsAccountPropertiesPtrOutput {
	return o
}

func (o MapsAccountPropertiesPtrOutput) ToMapsAccountPropertiesPtrOutputWithContext(ctx context.Context) MapsAccountPropertiesPtrOutput {
	return o
}

func (o MapsAccountPropertiesPtrOutput) Elem() MapsAccountPropertiesOutput {
	return o.ApplyT(func(v *MapsAccountProperties) MapsAccountProperties {
		if v != nil {
			return *v
		}
		var ret MapsAccountProperties
		return ret
	}).(MapsAccountPropertiesOutput)
}

func (o MapsAccountPropertiesPtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MapsAccountProperties) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

type MapsAccountPropertiesResponse struct {
	DisableLocalAuth  *bool  `pulumi:"disableLocalAuth"`
	ProvisioningState string `pulumi:"provisioningState"`
	UniqueId          string `pulumi:"uniqueId"`
}





type MapsAccountPropertiesResponseInput interface {
	pulumi.Input

	ToMapsAccountPropertiesResponseOutput() MapsAccountPropertiesResponseOutput
	ToMapsAccountPropertiesResponseOutputWithContext(context.Context) MapsAccountPropertiesResponseOutput
}

type MapsAccountPropertiesResponseArgs struct {
	DisableLocalAuth  pulumi.BoolPtrInput `pulumi:"disableLocalAuth"`
	ProvisioningState pulumi.StringInput  `pulumi:"provisioningState"`
	UniqueId          pulumi.StringInput  `pulumi:"uniqueId"`
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

func (o MapsAccountPropertiesResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o MapsAccountPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MapsAccountPropertiesResponseOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) string { return v.UniqueId }).(pulumi.StringOutput)
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

func (o MapsAccountPropertiesResponsePtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MapsAccountPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

func (o MapsAccountPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MapsAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MapsAccountPropertiesResponsePtrOutput) UniqueId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MapsAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UniqueId
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
	pulumi.RegisterInputType(reflect.TypeOf((*CreatorPropertiesInput)(nil)).Elem(), CreatorPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreatorPropertiesPtrInput)(nil)).Elem(), CreatorPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreatorPropertiesResponseInput)(nil)).Elem(), CreatorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreatorPropertiesResponsePtrInput)(nil)).Elem(), CreatorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MapsAccountPropertiesInput)(nil)).Elem(), MapsAccountPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MapsAccountPropertiesPtrInput)(nil)).Elem(), MapsAccountPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MapsAccountPropertiesResponseInput)(nil)).Elem(), MapsAccountPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MapsAccountPropertiesResponsePtrInput)(nil)).Elem(), MapsAccountPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterOutputType(CreatorPropertiesOutput{})
	pulumi.RegisterOutputType(CreatorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreatorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CreatorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
