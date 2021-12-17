


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

func (o CreatorPropertiesOutput) StorageUnits() pulumi.IntOutput {
	return o.ApplyT(func(v CreatorProperties) int { return v.StorageUnits }).(pulumi.IntOutput)
}

type CreatorPropertiesResponse struct {
	ProvisioningState string `pulumi:"provisioningState"`
	StorageUnits      int    `pulumi:"storageUnits"`
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

func (o CreatorPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v CreatorPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CreatorPropertiesResponseOutput) StorageUnits() pulumi.IntOutput {
	return o.ApplyT(func(v CreatorPropertiesResponse) int { return v.StorageUnits }).(pulumi.IntOutput)
}

type MapsAccountProperties struct {
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
}


func (val *MapsAccountProperties) Defaults() *MapsAccountProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DisableLocalAuth) {
		disableLocalAuth_ := false
		tmp.DisableLocalAuth = &disableLocalAuth_
	}
	return &tmp
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


func (val *MapsAccountPropertiesResponse) Defaults() *MapsAccountPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DisableLocalAuth) {
		disableLocalAuth_ := false
		tmp.DisableLocalAuth = &disableLocalAuth_
	}
	return &tmp
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

func (o MapsAccountPropertiesResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o MapsAccountPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MapsAccountPropertiesResponseOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) string { return v.UniqueId }).(pulumi.StringOutput)
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

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
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

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
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

func init() {
	pulumi.RegisterOutputType(CreatorPropertiesOutput{})
	pulumi.RegisterOutputType(CreatorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
