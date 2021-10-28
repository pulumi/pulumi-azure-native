


package hybridcompute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ErrorDetailResponse struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
	Target  *string               `pulumi:"target"`
}





type ErrorDetailResponseInput interface {
	pulumi.Input

	ToErrorDetailResponseOutput() ErrorDetailResponseOutput
	ToErrorDetailResponseOutputWithContext(context.Context) ErrorDetailResponseOutput
}

type ErrorDetailResponseArgs struct {
	Code    pulumi.StringInput            `pulumi:"code"`
	Details ErrorDetailResponseArrayInput `pulumi:"details"`
	Message pulumi.StringInput            `pulumi:"message"`
	Target  pulumi.StringPtrInput         `pulumi:"target"`
}

func (ErrorDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return i.ToErrorDetailResponseOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput)
}





type ErrorDetailResponseArrayInput interface {
	pulumi.Input

	ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput
	ToErrorDetailResponseArrayOutputWithContext(context.Context) ErrorDetailResponseArrayOutput
}

type ErrorDetailResponseArray []ErrorDetailResponseInput

func (ErrorDetailResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return i.ToErrorDetailResponseArrayOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseArrayOutput)
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorDetailResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
}

type HybridComputePrivateLinkScopeProperties struct {
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}





type HybridComputePrivateLinkScopePropertiesInput interface {
	pulumi.Input

	ToHybridComputePrivateLinkScopePropertiesOutput() HybridComputePrivateLinkScopePropertiesOutput
	ToHybridComputePrivateLinkScopePropertiesOutputWithContext(context.Context) HybridComputePrivateLinkScopePropertiesOutput
}

type HybridComputePrivateLinkScopePropertiesArgs struct {
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (HybridComputePrivateLinkScopePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputePrivateLinkScopeProperties)(nil)).Elem()
}

func (i HybridComputePrivateLinkScopePropertiesArgs) ToHybridComputePrivateLinkScopePropertiesOutput() HybridComputePrivateLinkScopePropertiesOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesOutputWithContext(context.Background())
}

func (i HybridComputePrivateLinkScopePropertiesArgs) ToHybridComputePrivateLinkScopePropertiesOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesOutput)
}

func (i HybridComputePrivateLinkScopePropertiesArgs) ToHybridComputePrivateLinkScopePropertiesPtrOutput() HybridComputePrivateLinkScopePropertiesPtrOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(context.Background())
}

func (i HybridComputePrivateLinkScopePropertiesArgs) ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesOutput).ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(ctx)
}









type HybridComputePrivateLinkScopePropertiesPtrInput interface {
	pulumi.Input

	ToHybridComputePrivateLinkScopePropertiesPtrOutput() HybridComputePrivateLinkScopePropertiesPtrOutput
	ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(context.Context) HybridComputePrivateLinkScopePropertiesPtrOutput
}

type hybridComputePrivateLinkScopePropertiesPtrType HybridComputePrivateLinkScopePropertiesArgs

func HybridComputePrivateLinkScopePropertiesPtr(v *HybridComputePrivateLinkScopePropertiesArgs) HybridComputePrivateLinkScopePropertiesPtrInput {
	return (*hybridComputePrivateLinkScopePropertiesPtrType)(v)
}

func (*hybridComputePrivateLinkScopePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputePrivateLinkScopeProperties)(nil)).Elem()
}

func (i *hybridComputePrivateLinkScopePropertiesPtrType) ToHybridComputePrivateLinkScopePropertiesPtrOutput() HybridComputePrivateLinkScopePropertiesPtrOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(context.Background())
}

func (i *hybridComputePrivateLinkScopePropertiesPtrType) ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesPtrOutput)
}

type HybridComputePrivateLinkScopePropertiesOutput struct{ *pulumi.OutputState }

func (HybridComputePrivateLinkScopePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputePrivateLinkScopeProperties)(nil)).Elem()
}

func (o HybridComputePrivateLinkScopePropertiesOutput) ToHybridComputePrivateLinkScopePropertiesOutput() HybridComputePrivateLinkScopePropertiesOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesOutput) ToHybridComputePrivateLinkScopePropertiesOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesOutput) ToHybridComputePrivateLinkScopePropertiesPtrOutput() HybridComputePrivateLinkScopePropertiesPtrOutput {
	return o.ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(context.Background())
}

func (o HybridComputePrivateLinkScopePropertiesOutput) ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HybridComputePrivateLinkScopeProperties) *HybridComputePrivateLinkScopeProperties {
		return &v
	}).(HybridComputePrivateLinkScopePropertiesPtrOutput)
}

func (o HybridComputePrivateLinkScopePropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputePrivateLinkScopeProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type HybridComputePrivateLinkScopePropertiesPtrOutput struct{ *pulumi.OutputState }

func (HybridComputePrivateLinkScopePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputePrivateLinkScopeProperties)(nil)).Elem()
}

func (o HybridComputePrivateLinkScopePropertiesPtrOutput) ToHybridComputePrivateLinkScopePropertiesPtrOutput() HybridComputePrivateLinkScopePropertiesPtrOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesPtrOutput) ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesPtrOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesPtrOutput) Elem() HybridComputePrivateLinkScopePropertiesOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopeProperties) HybridComputePrivateLinkScopeProperties {
		if v != nil {
			return *v
		}
		var ret HybridComputePrivateLinkScopeProperties
		return ret
	}).(HybridComputePrivateLinkScopePropertiesOutput)
}

func (o HybridComputePrivateLinkScopePropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopeProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type HybridComputePrivateLinkScopePropertiesResponse struct {
	PrivateLinkScopeId  string  `pulumi:"privateLinkScopeId"`
	ProvisioningState   string  `pulumi:"provisioningState"`
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}





type HybridComputePrivateLinkScopePropertiesResponseInput interface {
	pulumi.Input

	ToHybridComputePrivateLinkScopePropertiesResponseOutput() HybridComputePrivateLinkScopePropertiesResponseOutput
	ToHybridComputePrivateLinkScopePropertiesResponseOutputWithContext(context.Context) HybridComputePrivateLinkScopePropertiesResponseOutput
}

type HybridComputePrivateLinkScopePropertiesResponseArgs struct {
	PrivateLinkScopeId  pulumi.StringInput    `pulumi:"privateLinkScopeId"`
	ProvisioningState   pulumi.StringInput    `pulumi:"provisioningState"`
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (HybridComputePrivateLinkScopePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputePrivateLinkScopePropertiesResponse)(nil)).Elem()
}

func (i HybridComputePrivateLinkScopePropertiesResponseArgs) ToHybridComputePrivateLinkScopePropertiesResponseOutput() HybridComputePrivateLinkScopePropertiesResponseOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesResponseOutputWithContext(context.Background())
}

func (i HybridComputePrivateLinkScopePropertiesResponseArgs) ToHybridComputePrivateLinkScopePropertiesResponseOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesResponseOutput)
}

func (i HybridComputePrivateLinkScopePropertiesResponseArgs) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutput() HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i HybridComputePrivateLinkScopePropertiesResponseArgs) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesResponseOutput).ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(ctx)
}









type HybridComputePrivateLinkScopePropertiesResponsePtrInput interface {
	pulumi.Input

	ToHybridComputePrivateLinkScopePropertiesResponsePtrOutput() HybridComputePrivateLinkScopePropertiesResponsePtrOutput
	ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(context.Context) HybridComputePrivateLinkScopePropertiesResponsePtrOutput
}

type hybridComputePrivateLinkScopePropertiesResponsePtrType HybridComputePrivateLinkScopePropertiesResponseArgs

func HybridComputePrivateLinkScopePropertiesResponsePtr(v *HybridComputePrivateLinkScopePropertiesResponseArgs) HybridComputePrivateLinkScopePropertiesResponsePtrInput {
	return (*hybridComputePrivateLinkScopePropertiesResponsePtrType)(v)
}

func (*hybridComputePrivateLinkScopePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputePrivateLinkScopePropertiesResponse)(nil)).Elem()
}

func (i *hybridComputePrivateLinkScopePropertiesResponsePtrType) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutput() HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *hybridComputePrivateLinkScopePropertiesResponsePtrType) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesResponsePtrOutput)
}

type HybridComputePrivateLinkScopePropertiesResponseOutput struct{ *pulumi.OutputState }

func (HybridComputePrivateLinkScopePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputePrivateLinkScopePropertiesResponse)(nil)).Elem()
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) ToHybridComputePrivateLinkScopePropertiesResponseOutput() HybridComputePrivateLinkScopePropertiesResponseOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) ToHybridComputePrivateLinkScopePropertiesResponseOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponseOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutput() HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return o.ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HybridComputePrivateLinkScopePropertiesResponse) *HybridComputePrivateLinkScopePropertiesResponse {
		return &v
	}).(HybridComputePrivateLinkScopePropertiesResponsePtrOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) PrivateLinkScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v HybridComputePrivateLinkScopePropertiesResponse) string { return v.PrivateLinkScopeId }).(pulumi.StringOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v HybridComputePrivateLinkScopePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputePrivateLinkScopePropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type HybridComputePrivateLinkScopePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (HybridComputePrivateLinkScopePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputePrivateLinkScopePropertiesResponse)(nil)).Elem()
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutput() HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) Elem() HybridComputePrivateLinkScopePropertiesResponseOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopePropertiesResponse) HybridComputePrivateLinkScopePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret HybridComputePrivateLinkScopePropertiesResponse
		return ret
	}).(HybridComputePrivateLinkScopePropertiesResponseOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) PrivateLinkScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrivateLinkScopeId
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type LocationData struct {
	City            *string `pulumi:"city"`
	CountryOrRegion *string `pulumi:"countryOrRegion"`
	District        *string `pulumi:"district"`
	Name            string  `pulumi:"name"`
}





type LocationDataInput interface {
	pulumi.Input

	ToLocationDataOutput() LocationDataOutput
	ToLocationDataOutputWithContext(context.Context) LocationDataOutput
}

type LocationDataArgs struct {
	City            pulumi.StringPtrInput `pulumi:"city"`
	CountryOrRegion pulumi.StringPtrInput `pulumi:"countryOrRegion"`
	District        pulumi.StringPtrInput `pulumi:"district"`
	Name            pulumi.StringInput    `pulumi:"name"`
}

func (LocationDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationData)(nil)).Elem()
}

func (i LocationDataArgs) ToLocationDataOutput() LocationDataOutput {
	return i.ToLocationDataOutputWithContext(context.Background())
}

func (i LocationDataArgs) ToLocationDataOutputWithContext(ctx context.Context) LocationDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataOutput)
}

func (i LocationDataArgs) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return i.ToLocationDataPtrOutputWithContext(context.Background())
}

func (i LocationDataArgs) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataOutput).ToLocationDataPtrOutputWithContext(ctx)
}









type LocationDataPtrInput interface {
	pulumi.Input

	ToLocationDataPtrOutput() LocationDataPtrOutput
	ToLocationDataPtrOutputWithContext(context.Context) LocationDataPtrOutput
}

type locationDataPtrType LocationDataArgs

func LocationDataPtr(v *LocationDataArgs) LocationDataPtrInput {
	return (*locationDataPtrType)(v)
}

func (*locationDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationData)(nil)).Elem()
}

func (i *locationDataPtrType) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return i.ToLocationDataPtrOutputWithContext(context.Background())
}

func (i *locationDataPtrType) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataPtrOutput)
}

type LocationDataOutput struct{ *pulumi.OutputState }

func (LocationDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationData)(nil)).Elem()
}

func (o LocationDataOutput) ToLocationDataOutput() LocationDataOutput {
	return o
}

func (o LocationDataOutput) ToLocationDataOutputWithContext(ctx context.Context) LocationDataOutput {
	return o
}

func (o LocationDataOutput) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return o.ToLocationDataPtrOutputWithContext(context.Background())
}

func (o LocationDataOutput) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocationData) *LocationData {
		return &v
	}).(LocationDataPtrOutput)
}

func (o LocationDataOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o LocationDataOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

func (o LocationDataOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.District }).(pulumi.StringPtrOutput)
}

func (o LocationDataOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LocationData) string { return v.Name }).(pulumi.StringOutput)
}

type LocationDataPtrOutput struct{ *pulumi.OutputState }

func (LocationDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationData)(nil)).Elem()
}

func (o LocationDataPtrOutput) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return o
}

func (o LocationDataPtrOutput) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return o
}

func (o LocationDataPtrOutput) Elem() LocationDataOutput {
	return o.ApplyT(func(v *LocationData) LocationData {
		if v != nil {
			return *v
		}
		var ret LocationData
		return ret
	}).(LocationDataOutput)
}

func (o LocationDataPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataPtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataPtrOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.District
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type LocationDataResponse struct {
	City            *string `pulumi:"city"`
	CountryOrRegion *string `pulumi:"countryOrRegion"`
	District        *string `pulumi:"district"`
	Name            string  `pulumi:"name"`
}





type LocationDataResponseInput interface {
	pulumi.Input

	ToLocationDataResponseOutput() LocationDataResponseOutput
	ToLocationDataResponseOutputWithContext(context.Context) LocationDataResponseOutput
}

type LocationDataResponseArgs struct {
	City            pulumi.StringPtrInput `pulumi:"city"`
	CountryOrRegion pulumi.StringPtrInput `pulumi:"countryOrRegion"`
	District        pulumi.StringPtrInput `pulumi:"district"`
	Name            pulumi.StringInput    `pulumi:"name"`
}

func (LocationDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationDataResponse)(nil)).Elem()
}

func (i LocationDataResponseArgs) ToLocationDataResponseOutput() LocationDataResponseOutput {
	return i.ToLocationDataResponseOutputWithContext(context.Background())
}

func (i LocationDataResponseArgs) ToLocationDataResponseOutputWithContext(ctx context.Context) LocationDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataResponseOutput)
}

func (i LocationDataResponseArgs) ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput {
	return i.ToLocationDataResponsePtrOutputWithContext(context.Background())
}

func (i LocationDataResponseArgs) ToLocationDataResponsePtrOutputWithContext(ctx context.Context) LocationDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataResponseOutput).ToLocationDataResponsePtrOutputWithContext(ctx)
}









type LocationDataResponsePtrInput interface {
	pulumi.Input

	ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput
	ToLocationDataResponsePtrOutputWithContext(context.Context) LocationDataResponsePtrOutput
}

type locationDataResponsePtrType LocationDataResponseArgs

func LocationDataResponsePtr(v *LocationDataResponseArgs) LocationDataResponsePtrInput {
	return (*locationDataResponsePtrType)(v)
}

func (*locationDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationDataResponse)(nil)).Elem()
}

func (i *locationDataResponsePtrType) ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput {
	return i.ToLocationDataResponsePtrOutputWithContext(context.Background())
}

func (i *locationDataResponsePtrType) ToLocationDataResponsePtrOutputWithContext(ctx context.Context) LocationDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataResponsePtrOutput)
}

type LocationDataResponseOutput struct{ *pulumi.OutputState }

func (LocationDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationDataResponse)(nil)).Elem()
}

func (o LocationDataResponseOutput) ToLocationDataResponseOutput() LocationDataResponseOutput {
	return o
}

func (o LocationDataResponseOutput) ToLocationDataResponseOutputWithContext(ctx context.Context) LocationDataResponseOutput {
	return o
}

func (o LocationDataResponseOutput) ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput {
	return o.ToLocationDataResponsePtrOutputWithContext(context.Background())
}

func (o LocationDataResponseOutput) ToLocationDataResponsePtrOutputWithContext(ctx context.Context) LocationDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocationDataResponse) *LocationDataResponse {
		return &v
	}).(LocationDataResponsePtrOutput)
}

func (o LocationDataResponseOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationDataResponse) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o LocationDataResponseOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationDataResponse) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

func (o LocationDataResponseOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationDataResponse) *string { return v.District }).(pulumi.StringPtrOutput)
}

func (o LocationDataResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LocationDataResponse) string { return v.Name }).(pulumi.StringOutput)
}

type LocationDataResponsePtrOutput struct{ *pulumi.OutputState }

func (LocationDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationDataResponse)(nil)).Elem()
}

func (o LocationDataResponsePtrOutput) ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput {
	return o
}

func (o LocationDataResponsePtrOutput) ToLocationDataResponsePtrOutputWithContext(ctx context.Context) LocationDataResponsePtrOutput {
	return o
}

func (o LocationDataResponsePtrOutput) Elem() LocationDataResponseOutput {
	return o.ApplyT(func(v *LocationDataResponse) LocationDataResponse {
		if v != nil {
			return *v
		}
		var ret LocationDataResponse
		return ret
	}).(LocationDataResponseOutput)
}

func (o LocationDataResponsePtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataResponsePtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataResponsePtrOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.District
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewResponse struct {
	Name               string                                      `pulumi:"name"`
	Status             *MachineExtensionInstanceViewResponseStatus `pulumi:"status"`
	Type               string                                      `pulumi:"type"`
	TypeHandlerVersion string                                      `pulumi:"typeHandlerVersion"`
}





type MachineExtensionInstanceViewResponseInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseOutput() MachineExtensionInstanceViewResponseOutput
	ToMachineExtensionInstanceViewResponseOutputWithContext(context.Context) MachineExtensionInstanceViewResponseOutput
}

type MachineExtensionInstanceViewResponseArgs struct {
	Name               pulumi.StringInput                                 `pulumi:"name"`
	Status             MachineExtensionInstanceViewResponseStatusPtrInput `pulumi:"status"`
	Type               pulumi.StringInput                                 `pulumi:"type"`
	TypeHandlerVersion pulumi.StringInput                                 `pulumi:"typeHandlerVersion"`
}

func (MachineExtensionInstanceViewResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (i MachineExtensionInstanceViewResponseArgs) ToMachineExtensionInstanceViewResponseOutput() MachineExtensionInstanceViewResponseOutput {
	return i.ToMachineExtensionInstanceViewResponseOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseArgs) ToMachineExtensionInstanceViewResponseOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseOutput)
}





type MachineExtensionInstanceViewResponseArrayInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseArrayOutput() MachineExtensionInstanceViewResponseArrayOutput
	ToMachineExtensionInstanceViewResponseArrayOutputWithContext(context.Context) MachineExtensionInstanceViewResponseArrayOutput
}

type MachineExtensionInstanceViewResponseArray []MachineExtensionInstanceViewResponseInput

func (MachineExtensionInstanceViewResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (i MachineExtensionInstanceViewResponseArray) ToMachineExtensionInstanceViewResponseArrayOutput() MachineExtensionInstanceViewResponseArrayOutput {
	return i.ToMachineExtensionInstanceViewResponseArrayOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseArray) ToMachineExtensionInstanceViewResponseArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseArrayOutput)
}

type MachineExtensionInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseOutput) ToMachineExtensionInstanceViewResponseOutput() MachineExtensionInstanceViewResponseOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseOutput) ToMachineExtensionInstanceViewResponseOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) *MachineExtensionInstanceViewResponseStatus {
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type MachineExtensionInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseArrayOutput) ToMachineExtensionInstanceViewResponseArrayOutput() MachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseArrayOutput) ToMachineExtensionInstanceViewResponseArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) MachineExtensionInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineExtensionInstanceViewResponse {
		return vs[0].([]MachineExtensionInstanceViewResponse)[vs[1].(int)]
	}).(MachineExtensionInstanceViewResponseOutput)
}

type MachineExtensionInstanceViewResponseStatus struct {
	Code          string `pulumi:"code"`
	DisplayStatus string `pulumi:"displayStatus"`
	Level         string `pulumi:"level"`
	Message       string `pulumi:"message"`
	Time          string `pulumi:"time"`
}





type MachineExtensionInstanceViewResponseStatusInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput
	ToMachineExtensionInstanceViewResponseStatusOutputWithContext(context.Context) MachineExtensionInstanceViewResponseStatusOutput
}

type MachineExtensionInstanceViewResponseStatusArgs struct {
	Code          pulumi.StringInput `pulumi:"code"`
	DisplayStatus pulumi.StringInput `pulumi:"displayStatus"`
	Level         pulumi.StringInput `pulumi:"level"`
	Message       pulumi.StringInput `pulumi:"message"`
	Time          pulumi.StringInput `pulumi:"time"`
}

func (MachineExtensionInstanceViewResponseStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput {
	return i.ToMachineExtensionInstanceViewResponseStatusOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseStatusOutput)
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return i.ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseStatusOutput).ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx)
}









type MachineExtensionInstanceViewResponseStatusPtrInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput
	ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput
}

type machineExtensionInstanceViewResponseStatusPtrType MachineExtensionInstanceViewResponseStatusArgs

func MachineExtensionInstanceViewResponseStatusPtr(v *MachineExtensionInstanceViewResponseStatusArgs) MachineExtensionInstanceViewResponseStatusPtrInput {
	return (*machineExtensionInstanceViewResponseStatusPtrType)(v)
}

func (*machineExtensionInstanceViewResponseStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (i *machineExtensionInstanceViewResponseStatusPtrType) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return i.ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Background())
}

func (i *machineExtensionInstanceViewResponseStatusPtrType) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

type MachineExtensionInstanceViewResponseStatusOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Background())
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionInstanceViewResponseStatus) *MachineExtensionInstanceViewResponseStatus {
		return &v
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Code }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) DisplayStatus() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.DisplayStatus }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Level }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Message }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Time }).(pulumi.StringOutput)
}

type MachineExtensionInstanceViewResponseStatusPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Elem() MachineExtensionInstanceViewResponseStatusOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) MachineExtensionInstanceViewResponseStatus {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceViewResponseStatus
		return ret
	}).(MachineExtensionInstanceViewResponseStatusOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Level
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Time
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionPropertiesResponseInstanceView struct {
	Name               string                                      `pulumi:"name"`
	Status             *MachineExtensionInstanceViewResponseStatus `pulumi:"status"`
	Type               string                                      `pulumi:"type"`
	TypeHandlerVersion string                                      `pulumi:"typeHandlerVersion"`
}





type MachineExtensionPropertiesResponseInstanceViewInput interface {
	pulumi.Input

	ToMachineExtensionPropertiesResponseInstanceViewOutput() MachineExtensionPropertiesResponseInstanceViewOutput
	ToMachineExtensionPropertiesResponseInstanceViewOutputWithContext(context.Context) MachineExtensionPropertiesResponseInstanceViewOutput
}

type MachineExtensionPropertiesResponseInstanceViewArgs struct {
	Name               pulumi.StringInput                                 `pulumi:"name"`
	Status             MachineExtensionInstanceViewResponseStatusPtrInput `pulumi:"status"`
	Type               pulumi.StringInput                                 `pulumi:"type"`
	TypeHandlerVersion pulumi.StringInput                                 `pulumi:"typeHandlerVersion"`
}

func (MachineExtensionPropertiesResponseInstanceViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionPropertiesResponseInstanceView)(nil)).Elem()
}

func (i MachineExtensionPropertiesResponseInstanceViewArgs) ToMachineExtensionPropertiesResponseInstanceViewOutput() MachineExtensionPropertiesResponseInstanceViewOutput {
	return i.ToMachineExtensionPropertiesResponseInstanceViewOutputWithContext(context.Background())
}

func (i MachineExtensionPropertiesResponseInstanceViewArgs) ToMachineExtensionPropertiesResponseInstanceViewOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionPropertiesResponseInstanceViewOutput)
}

func (i MachineExtensionPropertiesResponseInstanceViewArgs) ToMachineExtensionPropertiesResponseInstanceViewPtrOutput() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return i.ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(context.Background())
}

func (i MachineExtensionPropertiesResponseInstanceViewArgs) ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionPropertiesResponseInstanceViewOutput).ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(ctx)
}









type MachineExtensionPropertiesResponseInstanceViewPtrInput interface {
	pulumi.Input

	ToMachineExtensionPropertiesResponseInstanceViewPtrOutput() MachineExtensionPropertiesResponseInstanceViewPtrOutput
	ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(context.Context) MachineExtensionPropertiesResponseInstanceViewPtrOutput
}

type machineExtensionPropertiesResponseInstanceViewPtrType MachineExtensionPropertiesResponseInstanceViewArgs

func MachineExtensionPropertiesResponseInstanceViewPtr(v *MachineExtensionPropertiesResponseInstanceViewArgs) MachineExtensionPropertiesResponseInstanceViewPtrInput {
	return (*machineExtensionPropertiesResponseInstanceViewPtrType)(v)
}

func (*machineExtensionPropertiesResponseInstanceViewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionPropertiesResponseInstanceView)(nil)).Elem()
}

func (i *machineExtensionPropertiesResponseInstanceViewPtrType) ToMachineExtensionPropertiesResponseInstanceViewPtrOutput() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return i.ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(context.Background())
}

func (i *machineExtensionPropertiesResponseInstanceViewPtrType) ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionPropertiesResponseInstanceViewPtrOutput)
}

type MachineExtensionPropertiesResponseInstanceViewOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesResponseInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionPropertiesResponseInstanceView)(nil)).Elem()
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) ToMachineExtensionPropertiesResponseInstanceViewOutput() MachineExtensionPropertiesResponseInstanceViewOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) ToMachineExtensionPropertiesResponseInstanceViewOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) ToMachineExtensionPropertiesResponseInstanceViewPtrOutput() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o.ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(context.Background())
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionPropertiesResponseInstanceView) *MachineExtensionPropertiesResponseInstanceView {
		return &v
	}).(MachineExtensionPropertiesResponseInstanceViewPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.Name }).(pulumi.StringOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) *MachineExtensionInstanceViewResponseStatus {
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.Type }).(pulumi.StringOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type MachineExtensionPropertiesResponseInstanceViewPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesResponseInstanceViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionPropertiesResponseInstanceView)(nil)).Elem()
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) ToMachineExtensionPropertiesResponseInstanceViewPtrOutput() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Elem() MachineExtensionPropertiesResponseInstanceViewOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) MachineExtensionPropertiesResponseInstanceView {
		if v != nil {
			return *v
		}
		var ret MachineExtensionPropertiesResponseInstanceView
		return ret
	}).(MachineExtensionPropertiesResponseInstanceViewOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *MachineExtensionInstanceViewResponseStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type MachineIdentity struct {
	Type *string `pulumi:"type"`
}





type MachineIdentityInput interface {
	pulumi.Input

	ToMachineIdentityOutput() MachineIdentityOutput
	ToMachineIdentityOutputWithContext(context.Context) MachineIdentityOutput
}

type MachineIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (MachineIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineIdentity)(nil)).Elem()
}

func (i MachineIdentityArgs) ToMachineIdentityOutput() MachineIdentityOutput {
	return i.ToMachineIdentityOutputWithContext(context.Background())
}

func (i MachineIdentityArgs) ToMachineIdentityOutputWithContext(ctx context.Context) MachineIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineIdentityOutput)
}

func (i MachineIdentityArgs) ToMachineIdentityPtrOutput() MachineIdentityPtrOutput {
	return i.ToMachineIdentityPtrOutputWithContext(context.Background())
}

func (i MachineIdentityArgs) ToMachineIdentityPtrOutputWithContext(ctx context.Context) MachineIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineIdentityOutput).ToMachineIdentityPtrOutputWithContext(ctx)
}









type MachineIdentityPtrInput interface {
	pulumi.Input

	ToMachineIdentityPtrOutput() MachineIdentityPtrOutput
	ToMachineIdentityPtrOutputWithContext(context.Context) MachineIdentityPtrOutput
}

type machineIdentityPtrType MachineIdentityArgs

func MachineIdentityPtr(v *MachineIdentityArgs) MachineIdentityPtrInput {
	return (*machineIdentityPtrType)(v)
}

func (*machineIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineIdentity)(nil)).Elem()
}

func (i *machineIdentityPtrType) ToMachineIdentityPtrOutput() MachineIdentityPtrOutput {
	return i.ToMachineIdentityPtrOutputWithContext(context.Background())
}

func (i *machineIdentityPtrType) ToMachineIdentityPtrOutputWithContext(ctx context.Context) MachineIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineIdentityPtrOutput)
}

type MachineIdentityOutput struct{ *pulumi.OutputState }

func (MachineIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineIdentity)(nil)).Elem()
}

func (o MachineIdentityOutput) ToMachineIdentityOutput() MachineIdentityOutput {
	return o
}

func (o MachineIdentityOutput) ToMachineIdentityOutputWithContext(ctx context.Context) MachineIdentityOutput {
	return o
}

func (o MachineIdentityOutput) ToMachineIdentityPtrOutput() MachineIdentityPtrOutput {
	return o.ToMachineIdentityPtrOutputWithContext(context.Background())
}

func (o MachineIdentityOutput) ToMachineIdentityPtrOutputWithContext(ctx context.Context) MachineIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineIdentity) *MachineIdentity {
		return &v
	}).(MachineIdentityPtrOutput)
}

func (o MachineIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type MachineIdentityPtrOutput struct{ *pulumi.OutputState }

func (MachineIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineIdentity)(nil)).Elem()
}

func (o MachineIdentityPtrOutput) ToMachineIdentityPtrOutput() MachineIdentityPtrOutput {
	return o
}

func (o MachineIdentityPtrOutput) ToMachineIdentityPtrOutputWithContext(ctx context.Context) MachineIdentityPtrOutput {
	return o
}

func (o MachineIdentityPtrOutput) Elem() MachineIdentityOutput {
	return o.ApplyT(func(v *MachineIdentity) MachineIdentity {
		if v != nil {
			return *v
		}
		var ret MachineIdentity
		return ret
	}).(MachineIdentityOutput)
}

func (o MachineIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type MachinePropertiesResponseOsProfile struct {
	ComputerName string `pulumi:"computerName"`
}





type MachinePropertiesResponseOsProfileInput interface {
	pulumi.Input

	ToMachinePropertiesResponseOsProfileOutput() MachinePropertiesResponseOsProfileOutput
	ToMachinePropertiesResponseOsProfileOutputWithContext(context.Context) MachinePropertiesResponseOsProfileOutput
}

type MachinePropertiesResponseOsProfileArgs struct {
	ComputerName pulumi.StringInput `pulumi:"computerName"`
}

func (MachinePropertiesResponseOsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachinePropertiesResponseOsProfile)(nil)).Elem()
}

func (i MachinePropertiesResponseOsProfileArgs) ToMachinePropertiesResponseOsProfileOutput() MachinePropertiesResponseOsProfileOutput {
	return i.ToMachinePropertiesResponseOsProfileOutputWithContext(context.Background())
}

func (i MachinePropertiesResponseOsProfileArgs) ToMachinePropertiesResponseOsProfileOutputWithContext(ctx context.Context) MachinePropertiesResponseOsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePropertiesResponseOsProfileOutput)
}

func (i MachinePropertiesResponseOsProfileArgs) ToMachinePropertiesResponseOsProfilePtrOutput() MachinePropertiesResponseOsProfilePtrOutput {
	return i.ToMachinePropertiesResponseOsProfilePtrOutputWithContext(context.Background())
}

func (i MachinePropertiesResponseOsProfileArgs) ToMachinePropertiesResponseOsProfilePtrOutputWithContext(ctx context.Context) MachinePropertiesResponseOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePropertiesResponseOsProfileOutput).ToMachinePropertiesResponseOsProfilePtrOutputWithContext(ctx)
}









type MachinePropertiesResponseOsProfilePtrInput interface {
	pulumi.Input

	ToMachinePropertiesResponseOsProfilePtrOutput() MachinePropertiesResponseOsProfilePtrOutput
	ToMachinePropertiesResponseOsProfilePtrOutputWithContext(context.Context) MachinePropertiesResponseOsProfilePtrOutput
}

type machinePropertiesResponseOsProfilePtrType MachinePropertiesResponseOsProfileArgs

func MachinePropertiesResponseOsProfilePtr(v *MachinePropertiesResponseOsProfileArgs) MachinePropertiesResponseOsProfilePtrInput {
	return (*machinePropertiesResponseOsProfilePtrType)(v)
}

func (*machinePropertiesResponseOsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachinePropertiesResponseOsProfile)(nil)).Elem()
}

func (i *machinePropertiesResponseOsProfilePtrType) ToMachinePropertiesResponseOsProfilePtrOutput() MachinePropertiesResponseOsProfilePtrOutput {
	return i.ToMachinePropertiesResponseOsProfilePtrOutputWithContext(context.Background())
}

func (i *machinePropertiesResponseOsProfilePtrType) ToMachinePropertiesResponseOsProfilePtrOutputWithContext(ctx context.Context) MachinePropertiesResponseOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePropertiesResponseOsProfilePtrOutput)
}

type MachinePropertiesResponseOsProfileOutput struct{ *pulumi.OutputState }

func (MachinePropertiesResponseOsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachinePropertiesResponseOsProfile)(nil)).Elem()
}

func (o MachinePropertiesResponseOsProfileOutput) ToMachinePropertiesResponseOsProfileOutput() MachinePropertiesResponseOsProfileOutput {
	return o
}

func (o MachinePropertiesResponseOsProfileOutput) ToMachinePropertiesResponseOsProfileOutputWithContext(ctx context.Context) MachinePropertiesResponseOsProfileOutput {
	return o
}

func (o MachinePropertiesResponseOsProfileOutput) ToMachinePropertiesResponseOsProfilePtrOutput() MachinePropertiesResponseOsProfilePtrOutput {
	return o.ToMachinePropertiesResponseOsProfilePtrOutputWithContext(context.Background())
}

func (o MachinePropertiesResponseOsProfileOutput) ToMachinePropertiesResponseOsProfilePtrOutputWithContext(ctx context.Context) MachinePropertiesResponseOsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachinePropertiesResponseOsProfile) *MachinePropertiesResponseOsProfile {
		return &v
	}).(MachinePropertiesResponseOsProfilePtrOutput)
}

func (o MachinePropertiesResponseOsProfileOutput) ComputerName() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponseOsProfile) string { return v.ComputerName }).(pulumi.StringOutput)
}

type MachinePropertiesResponseOsProfilePtrOutput struct{ *pulumi.OutputState }

func (MachinePropertiesResponseOsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachinePropertiesResponseOsProfile)(nil)).Elem()
}

func (o MachinePropertiesResponseOsProfilePtrOutput) ToMachinePropertiesResponseOsProfilePtrOutput() MachinePropertiesResponseOsProfilePtrOutput {
	return o
}

func (o MachinePropertiesResponseOsProfilePtrOutput) ToMachinePropertiesResponseOsProfilePtrOutputWithContext(ctx context.Context) MachinePropertiesResponseOsProfilePtrOutput {
	return o
}

func (o MachinePropertiesResponseOsProfilePtrOutput) Elem() MachinePropertiesResponseOsProfileOutput {
	return o.ApplyT(func(v *MachinePropertiesResponseOsProfile) MachinePropertiesResponseOsProfile {
		if v != nil {
			return *v
		}
		var ret MachinePropertiesResponseOsProfile
		return ret
	}).(MachinePropertiesResponseOsProfileOutput)
}

func (o MachinePropertiesResponseOsProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponseOsProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ComputerName
	}).(pulumi.StringPtrOutput)
}

type MachineResponseIdentity struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type MachineResponseIdentityInput interface {
	pulumi.Input

	ToMachineResponseIdentityOutput() MachineResponseIdentityOutput
	ToMachineResponseIdentityOutputWithContext(context.Context) MachineResponseIdentityOutput
}

type MachineResponseIdentityArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (MachineResponseIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineResponseIdentity)(nil)).Elem()
}

func (i MachineResponseIdentityArgs) ToMachineResponseIdentityOutput() MachineResponseIdentityOutput {
	return i.ToMachineResponseIdentityOutputWithContext(context.Background())
}

func (i MachineResponseIdentityArgs) ToMachineResponseIdentityOutputWithContext(ctx context.Context) MachineResponseIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineResponseIdentityOutput)
}

func (i MachineResponseIdentityArgs) ToMachineResponseIdentityPtrOutput() MachineResponseIdentityPtrOutput {
	return i.ToMachineResponseIdentityPtrOutputWithContext(context.Background())
}

func (i MachineResponseIdentityArgs) ToMachineResponseIdentityPtrOutputWithContext(ctx context.Context) MachineResponseIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineResponseIdentityOutput).ToMachineResponseIdentityPtrOutputWithContext(ctx)
}









type MachineResponseIdentityPtrInput interface {
	pulumi.Input

	ToMachineResponseIdentityPtrOutput() MachineResponseIdentityPtrOutput
	ToMachineResponseIdentityPtrOutputWithContext(context.Context) MachineResponseIdentityPtrOutput
}

type machineResponseIdentityPtrType MachineResponseIdentityArgs

func MachineResponseIdentityPtr(v *MachineResponseIdentityArgs) MachineResponseIdentityPtrInput {
	return (*machineResponseIdentityPtrType)(v)
}

func (*machineResponseIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineResponseIdentity)(nil)).Elem()
}

func (i *machineResponseIdentityPtrType) ToMachineResponseIdentityPtrOutput() MachineResponseIdentityPtrOutput {
	return i.ToMachineResponseIdentityPtrOutputWithContext(context.Background())
}

func (i *machineResponseIdentityPtrType) ToMachineResponseIdentityPtrOutputWithContext(ctx context.Context) MachineResponseIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineResponseIdentityPtrOutput)
}

type MachineResponseIdentityOutput struct{ *pulumi.OutputState }

func (MachineResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineResponseIdentity)(nil)).Elem()
}

func (o MachineResponseIdentityOutput) ToMachineResponseIdentityOutput() MachineResponseIdentityOutput {
	return o
}

func (o MachineResponseIdentityOutput) ToMachineResponseIdentityOutputWithContext(ctx context.Context) MachineResponseIdentityOutput {
	return o
}

func (o MachineResponseIdentityOutput) ToMachineResponseIdentityPtrOutput() MachineResponseIdentityPtrOutput {
	return o.ToMachineResponseIdentityPtrOutputWithContext(context.Background())
}

func (o MachineResponseIdentityOutput) ToMachineResponseIdentityPtrOutputWithContext(ctx context.Context) MachineResponseIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineResponseIdentity) *MachineResponseIdentity {
		return &v
	}).(MachineResponseIdentityPtrOutput)
}

func (o MachineResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v MachineResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o MachineResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v MachineResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o MachineResponseIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineResponseIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type MachineResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (MachineResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineResponseIdentity)(nil)).Elem()
}

func (o MachineResponseIdentityPtrOutput) ToMachineResponseIdentityPtrOutput() MachineResponseIdentityPtrOutput {
	return o
}

func (o MachineResponseIdentityPtrOutput) ToMachineResponseIdentityPtrOutputWithContext(ctx context.Context) MachineResponseIdentityPtrOutput {
	return o
}

func (o MachineResponseIdentityPtrOutput) Elem() MachineResponseIdentityOutput {
	return o.ApplyT(func(v *MachineResponseIdentity) MachineResponseIdentity {
		if v != nil {
			return *v
		}
		var ret MachineResponseIdentity
		return ret
	}).(MachineResponseIdentityOutput)
}

func (o MachineResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o MachineResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o MachineResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionProperties struct {
	PrivateEndpoint                   *PrivateEndpointProperty                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	PrivateEndpoint                   PrivateEndpointPropertyPtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyPtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return i.ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput)
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput).ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput
	ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPtrOutput
}

type privateEndpointConnectionPropertiesPtrType PrivateEndpointConnectionPropertiesArgs

func PrivateEndpointConnectionPropertiesPtr(v *PrivateEndpointConnectionPropertiesArgs) PrivateEndpointConnectionPropertiesPtrInput {
	return (*privateEndpointConnectionPropertiesPtrType)(v)
}

func (*privateEndpointConnectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPtrOutput)
}

type PrivateEndpointConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionProperties {
		return &v
	}).(PrivateEndpointConnectionPropertiesPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateEndpoint() PrivateEndpointPropertyPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *PrivateEndpointProperty { return v.PrivateEndpoint }).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionStateProperty {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateEndpointConnectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) Elem() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) PrivateEndpointConnectionProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionProperties
		return ret
	}).(PrivateEndpointConnectionPropertiesOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateEndpoint() PrivateEndpointPropertyPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateEndpointProperty {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionStateProperty {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                             `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput
	ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponseOutput
}

type PrivateEndpointConnectionPropertiesResponseArgs struct {
	PrivateEndpoint                   PrivateEndpointPropertyResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                        `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput).ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput
	ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput
}

type privateEndpointConnectionPropertiesResponsePtrType PrivateEndpointConnectionPropertiesResponseArgs

func PrivateEndpointConnectionPropertiesResponsePtr(v *PrivateEndpointConnectionPropertiesResponseArgs) PrivateEndpointConnectionPropertiesResponsePtrInput {
	return (*privateEndpointConnectionPropertiesResponsePtrType)(v)
}

func (*privateEndpointConnectionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponse {
		return &v
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointPropertyResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) Elem() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) PrivateEndpointConnectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponse
		return ret
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointPropertyResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointProperty struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput
	ToPrivateEndpointPropertyOutputWithContext(context.Context) PrivateEndpointPropertyOutput
}

type PrivateEndpointPropertyArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return i.ToPrivateEndpointPropertyOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput)
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput).ToPrivateEndpointPropertyPtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput
	ToPrivateEndpointPropertyPtrOutputWithContext(context.Context) PrivateEndpointPropertyPtrOutput
}

type privateEndpointPropertyPtrType PrivateEndpointPropertyArgs

func PrivateEndpointPropertyPtr(v *PrivateEndpointPropertyArgs) PrivateEndpointPropertyPtrInput {
	return (*privateEndpointPropertyPtrType)(v)
}

func (*privateEndpointPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyPtrOutput)
}

type PrivateEndpointPropertyOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointProperty) *PrivateEndpointProperty {
		return &v
	}).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointPropertyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointProperty) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) Elem() PrivateEndpointPropertyOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) PrivateEndpointProperty {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointProperty
		return ret
	}).(PrivateEndpointPropertyOutput)
}

func (o PrivateEndpointPropertyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponse struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyResponseInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput
	ToPrivateEndpointPropertyResponseOutputWithContext(context.Context) PrivateEndpointPropertyResponseOutput
}

type PrivateEndpointPropertyResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return i.ToPrivateEndpointPropertyResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput)
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput).ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput
	ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Context) PrivateEndpointPropertyResponsePtrOutput
}

type privateEndpointPropertyResponsePtrType PrivateEndpointPropertyResponseArgs

func PrivateEndpointPropertyResponsePtr(v *PrivateEndpointPropertyResponseArgs) PrivateEndpointPropertyResponsePtrInput {
	return (*privateEndpointPropertyResponsePtrType)(v)
}

func (*privateEndpointPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponsePtrOutput)
}

type PrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointPropertyResponse) *PrivateEndpointPropertyResponse {
		return &v
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointPropertyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointPropertyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) Elem() PrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) PrivateEndpointPropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointPropertyResponse
		return ret
	}).(PrivateEndpointPropertyResponseOutput)
}

func (o PrivateEndpointPropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateProperty struct {
	Description string `pulumi:"description"`
	Status      string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput
	ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyOutput
}

type PrivateLinkServiceConnectionStatePropertyArgs struct {
	Description pulumi.StringInput `pulumi:"description"`
	Status      pulumi.StringInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput).ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput
	ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput
}

type privateLinkServiceConnectionStatePropertyPtrType PrivateLinkServiceConnectionStatePropertyArgs

func PrivateLinkServiceConnectionStatePropertyPtr(v *PrivateLinkServiceConnectionStatePropertyArgs) PrivateLinkServiceConnectionStatePropertyPtrInput {
	return (*privateLinkServiceConnectionStatePropertyPtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateProperty) *PrivateLinkServiceConnectionStateProperty {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) PrivateLinkServiceConnectionStateProperty {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateProperty
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput
	ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput
}

type PrivateLinkServiceConnectionStatePropertyResponseArgs struct {
	ActionsRequired pulumi.StringInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput `pulumi:"description"`
	Status          pulumi.StringInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput).ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
}

type privateLinkServiceConnectionStatePropertyResponsePtrType PrivateLinkServiceConnectionStatePropertyResponseArgs

func PrivateLinkServiceConnectionStatePropertyResponsePtr(v *PrivateLinkServiceConnectionStatePropertyResponseArgs) PrivateLinkServiceConnectionStatePropertyResponsePtrInput {
	return (*privateLinkServiceConnectionStatePropertyResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStatePropertyResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) PrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
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
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesPtrOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesResponseOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LocationDataOutput{})
	pulumi.RegisterOutputType(LocationDataPtrOutput{})
	pulumi.RegisterOutputType(LocationDataResponseOutput{})
	pulumi.RegisterOutputType(LocationDataResponsePtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionPropertiesResponseInstanceViewOutput{})
	pulumi.RegisterOutputType(MachineExtensionPropertiesResponseInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(MachineIdentityOutput{})
	pulumi.RegisterOutputType(MachineIdentityPtrOutput{})
	pulumi.RegisterOutputType(MachinePropertiesResponseOsProfileOutput{})
	pulumi.RegisterOutputType(MachinePropertiesResponseOsProfilePtrOutput{})
	pulumi.RegisterOutputType(MachineResponseIdentityOutput{})
	pulumi.RegisterOutputType(MachineResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
