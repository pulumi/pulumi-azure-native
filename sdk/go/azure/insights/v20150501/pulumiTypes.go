


package v20150501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationInsightsComponentAnalyticsItemProperties struct {
	FunctionAlias *string `pulumi:"functionAlias"`
}





type ApplicationInsightsComponentAnalyticsItemPropertiesInput interface {
	pulumi.Input

	ToApplicationInsightsComponentAnalyticsItemPropertiesOutput() ApplicationInsightsComponentAnalyticsItemPropertiesOutput
	ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesOutput
}

type ApplicationInsightsComponentAnalyticsItemPropertiesArgs struct {
	FunctionAlias pulumi.StringPtrInput `pulumi:"functionAlias"`
}

func (ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentAnalyticsItemProperties)(nil)).Elem()
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesOutput() ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesOutput)
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesOutput).ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentAnalyticsItemPropertiesPtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput
	ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput
}

type applicationInsightsComponentAnalyticsItemPropertiesPtrType ApplicationInsightsComponentAnalyticsItemPropertiesArgs

func ApplicationInsightsComponentAnalyticsItemPropertiesPtr(v *ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ApplicationInsightsComponentAnalyticsItemPropertiesPtrInput {
	return (*applicationInsightsComponentAnalyticsItemPropertiesPtrType)(v)
}

func (*applicationInsightsComponentAnalyticsItemPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentAnalyticsItemProperties)(nil)).Elem()
}

func (i *applicationInsightsComponentAnalyticsItemPropertiesPtrType) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentAnalyticsItemPropertiesPtrType) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput)
}

type ApplicationInsightsComponentAnalyticsItemPropertiesOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentAnalyticsItemPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentAnalyticsItemProperties)(nil)).Elem()
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesOutput() ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return o.ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentAnalyticsItemProperties) *ApplicationInsightsComponentAnalyticsItemProperties {
		return &v
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput)
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentAnalyticsItemProperties) *string { return v.FunctionAlias }).(pulumi.StringPtrOutput)
}

type ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentAnalyticsItemProperties)(nil)).Elem()
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput) Elem() ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentAnalyticsItemProperties) ApplicationInsightsComponentAnalyticsItemProperties {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentAnalyticsItemProperties
		return ret
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesOutput)
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentAnalyticsItemProperties) *string {
		if v == nil {
			return nil
		}
		return v.FunctionAlias
	}).(pulumi.StringPtrOutput)
}

type ApplicationInsightsComponentAnalyticsItemPropertiesResponse struct {
	FunctionAlias *string `pulumi:"functionAlias"`
}





type ApplicationInsightsComponentAnalyticsItemPropertiesResponseInput interface {
	pulumi.Input

	ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput
	ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutputWithContext(context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput
}

type ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs struct {
	FunctionAlias pulumi.StringPtrInput `pulumi:"functionAlias"`
}

func (ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentAnalyticsItemPropertiesResponse)(nil)).Elem()
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput)
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput).ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput
	ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput
}

type applicationInsightsComponentAnalyticsItemPropertiesResponsePtrType ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs

func ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtr(v *ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrInput {
	return (*applicationInsightsComponentAnalyticsItemPropertiesResponsePtrType)(v)
}

func (*applicationInsightsComponentAnalyticsItemPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentAnalyticsItemPropertiesResponse)(nil)).Elem()
}

func (i *applicationInsightsComponentAnalyticsItemPropertiesResponsePtrType) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentAnalyticsItemPropertiesResponsePtrType) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput)
}

type ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentAnalyticsItemPropertiesResponse)(nil)).Elem()
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o.ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentAnalyticsItemPropertiesResponse) *ApplicationInsightsComponentAnalyticsItemPropertiesResponse {
		return &v
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput)
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentAnalyticsItemPropertiesResponse) *string { return v.FunctionAlias }).(pulumi.StringPtrOutput)
}

type ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentAnalyticsItemPropertiesResponse)(nil)).Elem()
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) Elem() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentAnalyticsItemPropertiesResponse) ApplicationInsightsComponentAnalyticsItemPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentAnalyticsItemPropertiesResponse
		return ret
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput)
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentAnalyticsItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FunctionAlias
	}).(pulumi.StringPtrOutput)
}

type ApplicationInsightsComponentDataVolumeCap struct {
	Cap                                  *float64 `pulumi:"cap"`
	StopSendNotificationWhenHitCap       *bool    `pulumi:"stopSendNotificationWhenHitCap"`
	StopSendNotificationWhenHitThreshold *bool    `pulumi:"stopSendNotificationWhenHitThreshold"`
	WarningThreshold                     *int     `pulumi:"warningThreshold"`
}





type ApplicationInsightsComponentDataVolumeCapInput interface {
	pulumi.Input

	ToApplicationInsightsComponentDataVolumeCapOutput() ApplicationInsightsComponentDataVolumeCapOutput
	ToApplicationInsightsComponentDataVolumeCapOutputWithContext(context.Context) ApplicationInsightsComponentDataVolumeCapOutput
}

type ApplicationInsightsComponentDataVolumeCapArgs struct {
	Cap                                  pulumi.Float64PtrInput `pulumi:"cap"`
	StopSendNotificationWhenHitCap       pulumi.BoolPtrInput    `pulumi:"stopSendNotificationWhenHitCap"`
	StopSendNotificationWhenHitThreshold pulumi.BoolPtrInput    `pulumi:"stopSendNotificationWhenHitThreshold"`
	WarningThreshold                     pulumi.IntPtrInput     `pulumi:"warningThreshold"`
}

func (ApplicationInsightsComponentDataVolumeCapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentDataVolumeCap)(nil)).Elem()
}

func (i ApplicationInsightsComponentDataVolumeCapArgs) ToApplicationInsightsComponentDataVolumeCapOutput() ApplicationInsightsComponentDataVolumeCapOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentDataVolumeCapArgs) ToApplicationInsightsComponentDataVolumeCapOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapOutput)
}

func (i ApplicationInsightsComponentDataVolumeCapArgs) ToApplicationInsightsComponentDataVolumeCapPtrOutput() ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentDataVolumeCapArgs) ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapOutput).ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentDataVolumeCapPtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentDataVolumeCapPtrOutput() ApplicationInsightsComponentDataVolumeCapPtrOutput
	ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(context.Context) ApplicationInsightsComponentDataVolumeCapPtrOutput
}

type applicationInsightsComponentDataVolumeCapPtrType ApplicationInsightsComponentDataVolumeCapArgs

func ApplicationInsightsComponentDataVolumeCapPtr(v *ApplicationInsightsComponentDataVolumeCapArgs) ApplicationInsightsComponentDataVolumeCapPtrInput {
	return (*applicationInsightsComponentDataVolumeCapPtrType)(v)
}

func (*applicationInsightsComponentDataVolumeCapPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentDataVolumeCap)(nil)).Elem()
}

func (i *applicationInsightsComponentDataVolumeCapPtrType) ToApplicationInsightsComponentDataVolumeCapPtrOutput() ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentDataVolumeCapPtrType) ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapPtrOutput)
}

type ApplicationInsightsComponentDataVolumeCapOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentDataVolumeCapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentDataVolumeCap)(nil)).Elem()
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) ToApplicationInsightsComponentDataVolumeCapOutput() ApplicationInsightsComponentDataVolumeCapOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) ToApplicationInsightsComponentDataVolumeCapOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) ToApplicationInsightsComponentDataVolumeCapPtrOutput() ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return o.ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentDataVolumeCap) *ApplicationInsightsComponentDataVolumeCap {
		return &v
	}).(ApplicationInsightsComponentDataVolumeCapPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) Cap() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCap) *float64 { return v.Cap }).(pulumi.Float64PtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) StopSendNotificationWhenHitCap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCap) *bool { return v.StopSendNotificationWhenHitCap }).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) StopSendNotificationWhenHitThreshold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCap) *bool { return v.StopSendNotificationWhenHitThreshold }).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) WarningThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCap) *int { return v.WarningThreshold }).(pulumi.IntPtrOutput)
}

type ApplicationInsightsComponentDataVolumeCapPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentDataVolumeCapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentDataVolumeCap)(nil)).Elem()
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) ToApplicationInsightsComponentDataVolumeCapPtrOutput() ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) Elem() ApplicationInsightsComponentDataVolumeCapOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCap) ApplicationInsightsComponentDataVolumeCap {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentDataVolumeCap
		return ret
	}).(ApplicationInsightsComponentDataVolumeCapOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) Cap() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCap) *float64 {
		if v == nil {
			return nil
		}
		return v.Cap
	}).(pulumi.Float64PtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) StopSendNotificationWhenHitCap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCap) *bool {
		if v == nil {
			return nil
		}
		return v.StopSendNotificationWhenHitCap
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) StopSendNotificationWhenHitThreshold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCap) *bool {
		if v == nil {
			return nil
		}
		return v.StopSendNotificationWhenHitThreshold
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) WarningThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCap) *int {
		if v == nil {
			return nil
		}
		return v.WarningThreshold
	}).(pulumi.IntPtrOutput)
}

type ApplicationInsightsComponentDataVolumeCapResponse struct {
	Cap                                  *float64 `pulumi:"cap"`
	MaxHistoryCap                        float64  `pulumi:"maxHistoryCap"`
	ResetTime                            int      `pulumi:"resetTime"`
	StopSendNotificationWhenHitCap       *bool    `pulumi:"stopSendNotificationWhenHitCap"`
	StopSendNotificationWhenHitThreshold *bool    `pulumi:"stopSendNotificationWhenHitThreshold"`
	WarningThreshold                     *int     `pulumi:"warningThreshold"`
}





type ApplicationInsightsComponentDataVolumeCapResponseInput interface {
	pulumi.Input

	ToApplicationInsightsComponentDataVolumeCapResponseOutput() ApplicationInsightsComponentDataVolumeCapResponseOutput
	ToApplicationInsightsComponentDataVolumeCapResponseOutputWithContext(context.Context) ApplicationInsightsComponentDataVolumeCapResponseOutput
}

type ApplicationInsightsComponentDataVolumeCapResponseArgs struct {
	Cap                                  pulumi.Float64PtrInput `pulumi:"cap"`
	MaxHistoryCap                        pulumi.Float64Input    `pulumi:"maxHistoryCap"`
	ResetTime                            pulumi.IntInput        `pulumi:"resetTime"`
	StopSendNotificationWhenHitCap       pulumi.BoolPtrInput    `pulumi:"stopSendNotificationWhenHitCap"`
	StopSendNotificationWhenHitThreshold pulumi.BoolPtrInput    `pulumi:"stopSendNotificationWhenHitThreshold"`
	WarningThreshold                     pulumi.IntPtrInput     `pulumi:"warningThreshold"`
}

func (ApplicationInsightsComponentDataVolumeCapResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentDataVolumeCapResponse)(nil)).Elem()
}

func (i ApplicationInsightsComponentDataVolumeCapResponseArgs) ToApplicationInsightsComponentDataVolumeCapResponseOutput() ApplicationInsightsComponentDataVolumeCapResponseOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapResponseOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentDataVolumeCapResponseArgs) ToApplicationInsightsComponentDataVolumeCapResponseOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapResponseOutput)
}

func (i ApplicationInsightsComponentDataVolumeCapResponseArgs) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutput() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentDataVolumeCapResponseArgs) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapResponseOutput).ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentDataVolumeCapResponsePtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentDataVolumeCapResponsePtrOutput() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput
	ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(context.Context) ApplicationInsightsComponentDataVolumeCapResponsePtrOutput
}

type applicationInsightsComponentDataVolumeCapResponsePtrType ApplicationInsightsComponentDataVolumeCapResponseArgs

func ApplicationInsightsComponentDataVolumeCapResponsePtr(v *ApplicationInsightsComponentDataVolumeCapResponseArgs) ApplicationInsightsComponentDataVolumeCapResponsePtrInput {
	return (*applicationInsightsComponentDataVolumeCapResponsePtrType)(v)
}

func (*applicationInsightsComponentDataVolumeCapResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentDataVolumeCapResponse)(nil)).Elem()
}

func (i *applicationInsightsComponentDataVolumeCapResponsePtrType) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutput() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentDataVolumeCapResponsePtrType) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapResponsePtrOutput)
}

type ApplicationInsightsComponentDataVolumeCapResponseOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentDataVolumeCapResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentDataVolumeCapResponse)(nil)).Elem()
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) ToApplicationInsightsComponentDataVolumeCapResponseOutput() ApplicationInsightsComponentDataVolumeCapResponseOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) ToApplicationInsightsComponentDataVolumeCapResponseOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponseOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutput() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return o.ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentDataVolumeCapResponse) *ApplicationInsightsComponentDataVolumeCapResponse {
		return &v
	}).(ApplicationInsightsComponentDataVolumeCapResponsePtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) Cap() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) *float64 { return v.Cap }).(pulumi.Float64PtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) MaxHistoryCap() pulumi.Float64Output {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) float64 { return v.MaxHistoryCap }).(pulumi.Float64Output)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) ResetTime() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) int { return v.ResetTime }).(pulumi.IntOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) StopSendNotificationWhenHitCap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) *bool {
		return v.StopSendNotificationWhenHitCap
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) StopSendNotificationWhenHitThreshold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) *bool {
		return v.StopSendNotificationWhenHitThreshold
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) WarningThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) *int { return v.WarningThreshold }).(pulumi.IntPtrOutput)
}

type ApplicationInsightsComponentDataVolumeCapResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentDataVolumeCapResponse)(nil)).Elem()
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutput() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) Elem() ApplicationInsightsComponentDataVolumeCapResponseOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) ApplicationInsightsComponentDataVolumeCapResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentDataVolumeCapResponse
		return ret
	}).(ApplicationInsightsComponentDataVolumeCapResponseOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) Cap() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Cap
	}).(pulumi.Float64PtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) MaxHistoryCap() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MaxHistoryCap
	}).(pulumi.Float64PtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) ResetTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ResetTime
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) StopSendNotificationWhenHitCap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *bool {
		if v == nil {
			return nil
		}
		return v.StopSendNotificationWhenHitCap
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) StopSendNotificationWhenHitThreshold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *bool {
		if v == nil {
			return nil
		}
		return v.StopSendNotificationWhenHitThreshold
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) WarningThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *int {
		if v == nil {
			return nil
		}
		return v.WarningThreshold
	}).(pulumi.IntPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions struct {
	Description                *string `pulumi:"description"`
	DisplayName                *string `pulumi:"displayName"`
	HelpUrl                    *string `pulumi:"helpUrl"`
	IsEnabledByDefault         *bool   `pulumi:"isEnabledByDefault"`
	IsHidden                   *bool   `pulumi:"isHidden"`
	IsInPreview                *bool   `pulumi:"isInPreview"`
	Name                       *string `pulumi:"name"`
	SupportsEmailNotifications *bool   `pulumi:"supportsEmailNotifications"`
}





type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsInput interface {
	pulumi.Input

	ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput
	ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutputWithContext(context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput
}

type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs struct {
	Description                pulumi.StringPtrInput `pulumi:"description"`
	DisplayName                pulumi.StringPtrInput `pulumi:"displayName"`
	HelpUrl                    pulumi.StringPtrInput `pulumi:"helpUrl"`
	IsEnabledByDefault         pulumi.BoolPtrInput   `pulumi:"isEnabledByDefault"`
	IsHidden                   pulumi.BoolPtrInput   `pulumi:"isHidden"`
	IsInPreview                pulumi.BoolPtrInput   `pulumi:"isInPreview"`
	Name                       pulumi.StringPtrInput `pulumi:"name"`
	SupportsEmailNotifications pulumi.BoolPtrInput   `pulumi:"supportsEmailNotifications"`
}

func (ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions)(nil)).Elem()
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput)
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput).ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput
	ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput
}

type applicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrType ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs

func ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtr(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrInput {
	return (*applicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrType)(v)
}

func (*applicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions)(nil)).Elem()
}

func (i *applicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrType) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrType) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return o.ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions {
		return &v
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) Elem() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions
		return ret
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions struct {
	Description                *string `pulumi:"description"`
	DisplayName                *string `pulumi:"displayName"`
	HelpUrl                    *string `pulumi:"helpUrl"`
	IsEnabledByDefault         *bool   `pulumi:"isEnabledByDefault"`
	IsHidden                   *bool   `pulumi:"isHidden"`
	IsInPreview                *bool   `pulumi:"isInPreview"`
	Name                       *string `pulumi:"name"`
	SupportsEmailNotifications *bool   `pulumi:"supportsEmailNotifications"`
}





type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsInput interface {
	pulumi.Input

	ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput
	ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutputWithContext(context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput
}

type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs struct {
	Description                pulumi.StringPtrInput `pulumi:"description"`
	DisplayName                pulumi.StringPtrInput `pulumi:"displayName"`
	HelpUrl                    pulumi.StringPtrInput `pulumi:"helpUrl"`
	IsEnabledByDefault         pulumi.BoolPtrInput   `pulumi:"isEnabledByDefault"`
	IsHidden                   pulumi.BoolPtrInput   `pulumi:"isHidden"`
	IsInPreview                pulumi.BoolPtrInput   `pulumi:"isInPreview"`
	Name                       pulumi.StringPtrInput `pulumi:"name"`
	SupportsEmailNotifications pulumi.BoolPtrInput   `pulumi:"supportsEmailNotifications"`
}

func (ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions)(nil)).Elem()
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput)
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput).ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput
	ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput
}

type applicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrType ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs

func ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtr(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrInput {
	return (*applicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrType)(v)
}

func (*applicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions)(nil)).Elem()
}

func (i *applicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrType) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrType) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return o.ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions {
		return &v
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) Elem() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions
		return ret
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

type PrivateLinkScopedResourceResponse struct {
	ResourceId *string `pulumi:"resourceId"`
	ScopeId    *string `pulumi:"scopeId"`
}





type PrivateLinkScopedResourceResponseInput interface {
	pulumi.Input

	ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput
	ToPrivateLinkScopedResourceResponseOutputWithContext(context.Context) PrivateLinkScopedResourceResponseOutput
}

type PrivateLinkScopedResourceResponseArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	ScopeId    pulumi.StringPtrInput `pulumi:"scopeId"`
}

func (PrivateLinkScopedResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (i PrivateLinkScopedResourceResponseArgs) ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput {
	return i.ToPrivateLinkScopedResourceResponseOutputWithContext(context.Background())
}

func (i PrivateLinkScopedResourceResponseArgs) ToPrivateLinkScopedResourceResponseOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopedResourceResponseOutput)
}





type PrivateLinkScopedResourceResponseArrayInput interface {
	pulumi.Input

	ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput
	ToPrivateLinkScopedResourceResponseArrayOutputWithContext(context.Context) PrivateLinkScopedResourceResponseArrayOutput
}

type PrivateLinkScopedResourceResponseArray []PrivateLinkScopedResourceResponseInput

func (PrivateLinkScopedResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (i PrivateLinkScopedResourceResponseArray) ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput {
	return i.ToPrivateLinkScopedResourceResponseArrayOutputWithContext(context.Background())
}

func (i PrivateLinkScopedResourceResponseArray) ToPrivateLinkScopedResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopedResourceResponseArrayOutput)
}

type PrivateLinkScopedResourceResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkScopedResourceResponseOutput) ScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ScopeId }).(pulumi.StringPtrOutput)
}

type PrivateLinkScopedResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkScopedResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkScopedResourceResponse {
		return vs[0].([]PrivateLinkScopedResourceResponse)[vs[1].(int)]
	}).(PrivateLinkScopedResourceResponseOutput)
}

type WebTestGeolocation struct {
	Location *string `pulumi:"location"`
}





type WebTestGeolocationInput interface {
	pulumi.Input

	ToWebTestGeolocationOutput() WebTestGeolocationOutput
	ToWebTestGeolocationOutputWithContext(context.Context) WebTestGeolocationOutput
}

type WebTestGeolocationArgs struct {
	Location pulumi.StringPtrInput `pulumi:"location"`
}

func (WebTestGeolocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestGeolocation)(nil)).Elem()
}

func (i WebTestGeolocationArgs) ToWebTestGeolocationOutput() WebTestGeolocationOutput {
	return i.ToWebTestGeolocationOutputWithContext(context.Background())
}

func (i WebTestGeolocationArgs) ToWebTestGeolocationOutputWithContext(ctx context.Context) WebTestGeolocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestGeolocationOutput)
}





type WebTestGeolocationArrayInput interface {
	pulumi.Input

	ToWebTestGeolocationArrayOutput() WebTestGeolocationArrayOutput
	ToWebTestGeolocationArrayOutputWithContext(context.Context) WebTestGeolocationArrayOutput
}

type WebTestGeolocationArray []WebTestGeolocationInput

func (WebTestGeolocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTestGeolocation)(nil)).Elem()
}

func (i WebTestGeolocationArray) ToWebTestGeolocationArrayOutput() WebTestGeolocationArrayOutput {
	return i.ToWebTestGeolocationArrayOutputWithContext(context.Background())
}

func (i WebTestGeolocationArray) ToWebTestGeolocationArrayOutputWithContext(ctx context.Context) WebTestGeolocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestGeolocationArrayOutput)
}

type WebTestGeolocationOutput struct{ *pulumi.OutputState }

func (WebTestGeolocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestGeolocation)(nil)).Elem()
}

func (o WebTestGeolocationOutput) ToWebTestGeolocationOutput() WebTestGeolocationOutput {
	return o
}

func (o WebTestGeolocationOutput) ToWebTestGeolocationOutputWithContext(ctx context.Context) WebTestGeolocationOutput {
	return o
}

func (o WebTestGeolocationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestGeolocation) *string { return v.Location }).(pulumi.StringPtrOutput)
}

type WebTestGeolocationArrayOutput struct{ *pulumi.OutputState }

func (WebTestGeolocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTestGeolocation)(nil)).Elem()
}

func (o WebTestGeolocationArrayOutput) ToWebTestGeolocationArrayOutput() WebTestGeolocationArrayOutput {
	return o
}

func (o WebTestGeolocationArrayOutput) ToWebTestGeolocationArrayOutputWithContext(ctx context.Context) WebTestGeolocationArrayOutput {
	return o
}

func (o WebTestGeolocationArrayOutput) Index(i pulumi.IntInput) WebTestGeolocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebTestGeolocation {
		return vs[0].([]WebTestGeolocation)[vs[1].(int)]
	}).(WebTestGeolocationOutput)
}

type WebTestGeolocationResponse struct {
	Location *string `pulumi:"location"`
}





type WebTestGeolocationResponseInput interface {
	pulumi.Input

	ToWebTestGeolocationResponseOutput() WebTestGeolocationResponseOutput
	ToWebTestGeolocationResponseOutputWithContext(context.Context) WebTestGeolocationResponseOutput
}

type WebTestGeolocationResponseArgs struct {
	Location pulumi.StringPtrInput `pulumi:"location"`
}

func (WebTestGeolocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestGeolocationResponse)(nil)).Elem()
}

func (i WebTestGeolocationResponseArgs) ToWebTestGeolocationResponseOutput() WebTestGeolocationResponseOutput {
	return i.ToWebTestGeolocationResponseOutputWithContext(context.Background())
}

func (i WebTestGeolocationResponseArgs) ToWebTestGeolocationResponseOutputWithContext(ctx context.Context) WebTestGeolocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestGeolocationResponseOutput)
}





type WebTestGeolocationResponseArrayInput interface {
	pulumi.Input

	ToWebTestGeolocationResponseArrayOutput() WebTestGeolocationResponseArrayOutput
	ToWebTestGeolocationResponseArrayOutputWithContext(context.Context) WebTestGeolocationResponseArrayOutput
}

type WebTestGeolocationResponseArray []WebTestGeolocationResponseInput

func (WebTestGeolocationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTestGeolocationResponse)(nil)).Elem()
}

func (i WebTestGeolocationResponseArray) ToWebTestGeolocationResponseArrayOutput() WebTestGeolocationResponseArrayOutput {
	return i.ToWebTestGeolocationResponseArrayOutputWithContext(context.Background())
}

func (i WebTestGeolocationResponseArray) ToWebTestGeolocationResponseArrayOutputWithContext(ctx context.Context) WebTestGeolocationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestGeolocationResponseArrayOutput)
}

type WebTestGeolocationResponseOutput struct{ *pulumi.OutputState }

func (WebTestGeolocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestGeolocationResponse)(nil)).Elem()
}

func (o WebTestGeolocationResponseOutput) ToWebTestGeolocationResponseOutput() WebTestGeolocationResponseOutput {
	return o
}

func (o WebTestGeolocationResponseOutput) ToWebTestGeolocationResponseOutputWithContext(ctx context.Context) WebTestGeolocationResponseOutput {
	return o
}

func (o WebTestGeolocationResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestGeolocationResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

type WebTestGeolocationResponseArrayOutput struct{ *pulumi.OutputState }

func (WebTestGeolocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTestGeolocationResponse)(nil)).Elem()
}

func (o WebTestGeolocationResponseArrayOutput) ToWebTestGeolocationResponseArrayOutput() WebTestGeolocationResponseArrayOutput {
	return o
}

func (o WebTestGeolocationResponseArrayOutput) ToWebTestGeolocationResponseArrayOutputWithContext(ctx context.Context) WebTestGeolocationResponseArrayOutput {
	return o
}

func (o WebTestGeolocationResponseArrayOutput) Index(i pulumi.IntInput) WebTestGeolocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebTestGeolocationResponse {
		return vs[0].([]WebTestGeolocationResponse)[vs[1].(int)]
	}).(WebTestGeolocationResponseOutput)
}

type WebTestPropertiesConfiguration struct {
	WebTest *string `pulumi:"webTest"`
}





type WebTestPropertiesConfigurationInput interface {
	pulumi.Input

	ToWebTestPropertiesConfigurationOutput() WebTestPropertiesConfigurationOutput
	ToWebTestPropertiesConfigurationOutputWithContext(context.Context) WebTestPropertiesConfigurationOutput
}

type WebTestPropertiesConfigurationArgs struct {
	WebTest pulumi.StringPtrInput `pulumi:"webTest"`
}

func (WebTestPropertiesConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesConfiguration)(nil)).Elem()
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationOutput() WebTestPropertiesConfigurationOutput {
	return i.ToWebTestPropertiesConfigurationOutputWithContext(context.Background())
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesConfigurationOutput)
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput {
	return i.ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Background())
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesConfigurationOutput).ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx)
}









type WebTestPropertiesConfigurationPtrInput interface {
	pulumi.Input

	ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput
	ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Context) WebTestPropertiesConfigurationPtrOutput
}

type webTestPropertiesConfigurationPtrType WebTestPropertiesConfigurationArgs

func WebTestPropertiesConfigurationPtr(v *WebTestPropertiesConfigurationArgs) WebTestPropertiesConfigurationPtrInput {
	return (*webTestPropertiesConfigurationPtrType)(v)
}

func (*webTestPropertiesConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesConfiguration)(nil)).Elem()
}

func (i *webTestPropertiesConfigurationPtrType) ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput {
	return i.ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Background())
}

func (i *webTestPropertiesConfigurationPtrType) ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesConfigurationPtrOutput)
}

type WebTestPropertiesConfigurationOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesConfiguration)(nil)).Elem()
}

func (o WebTestPropertiesConfigurationOutput) ToWebTestPropertiesConfigurationOutput() WebTestPropertiesConfigurationOutput {
	return o
}

func (o WebTestPropertiesConfigurationOutput) ToWebTestPropertiesConfigurationOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationOutput {
	return o
}

func (o WebTestPropertiesConfigurationOutput) ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput {
	return o.ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Background())
}

func (o WebTestPropertiesConfigurationOutput) ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebTestPropertiesConfiguration) *WebTestPropertiesConfiguration {
		return &v
	}).(WebTestPropertiesConfigurationPtrOutput)
}

func (o WebTestPropertiesConfigurationOutput) WebTest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesConfiguration) *string { return v.WebTest }).(pulumi.StringPtrOutput)
}

type WebTestPropertiesConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesConfiguration)(nil)).Elem()
}

func (o WebTestPropertiesConfigurationPtrOutput) ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput {
	return o
}

func (o WebTestPropertiesConfigurationPtrOutput) ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationPtrOutput {
	return o
}

func (o WebTestPropertiesConfigurationPtrOutput) Elem() WebTestPropertiesConfigurationOutput {
	return o.ApplyT(func(v *WebTestPropertiesConfiguration) WebTestPropertiesConfiguration {
		if v != nil {
			return *v
		}
		var ret WebTestPropertiesConfiguration
		return ret
	}).(WebTestPropertiesConfigurationOutput)
}

func (o WebTestPropertiesConfigurationPtrOutput) WebTest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.WebTest
	}).(pulumi.StringPtrOutput)
}

type WebTestPropertiesResponseConfiguration struct {
	WebTest *string `pulumi:"webTest"`
}





type WebTestPropertiesResponseConfigurationInput interface {
	pulumi.Input

	ToWebTestPropertiesResponseConfigurationOutput() WebTestPropertiesResponseConfigurationOutput
	ToWebTestPropertiesResponseConfigurationOutputWithContext(context.Context) WebTestPropertiesResponseConfigurationOutput
}

type WebTestPropertiesResponseConfigurationArgs struct {
	WebTest pulumi.StringPtrInput `pulumi:"webTest"`
}

func (WebTestPropertiesResponseConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesResponseConfiguration)(nil)).Elem()
}

func (i WebTestPropertiesResponseConfigurationArgs) ToWebTestPropertiesResponseConfigurationOutput() WebTestPropertiesResponseConfigurationOutput {
	return i.ToWebTestPropertiesResponseConfigurationOutputWithContext(context.Background())
}

func (i WebTestPropertiesResponseConfigurationArgs) ToWebTestPropertiesResponseConfigurationOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesResponseConfigurationOutput)
}

func (i WebTestPropertiesResponseConfigurationArgs) ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput {
	return i.ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(context.Background())
}

func (i WebTestPropertiesResponseConfigurationArgs) ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesResponseConfigurationOutput).ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx)
}









type WebTestPropertiesResponseConfigurationPtrInput interface {
	pulumi.Input

	ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput
	ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(context.Context) WebTestPropertiesResponseConfigurationPtrOutput
}

type webTestPropertiesResponseConfigurationPtrType WebTestPropertiesResponseConfigurationArgs

func WebTestPropertiesResponseConfigurationPtr(v *WebTestPropertiesResponseConfigurationArgs) WebTestPropertiesResponseConfigurationPtrInput {
	return (*webTestPropertiesResponseConfigurationPtrType)(v)
}

func (*webTestPropertiesResponseConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesResponseConfiguration)(nil)).Elem()
}

func (i *webTestPropertiesResponseConfigurationPtrType) ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput {
	return i.ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(context.Background())
}

func (i *webTestPropertiesResponseConfigurationPtrType) ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesResponseConfigurationPtrOutput)
}

type WebTestPropertiesResponseConfigurationOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesResponseConfiguration)(nil)).Elem()
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationOutput() WebTestPropertiesResponseConfigurationOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput {
	return o.ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(context.Background())
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebTestPropertiesResponseConfiguration) *WebTestPropertiesResponseConfiguration {
		return &v
	}).(WebTestPropertiesResponseConfigurationPtrOutput)
}

func (o WebTestPropertiesResponseConfigurationOutput) WebTest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseConfiguration) *string { return v.WebTest }).(pulumi.StringPtrOutput)
}

type WebTestPropertiesResponseConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesResponseConfiguration)(nil)).Elem()
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationPtrOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) Elem() WebTestPropertiesResponseConfigurationOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseConfiguration) WebTestPropertiesResponseConfiguration {
		if v != nil {
			return *v
		}
		var ret WebTestPropertiesResponseConfiguration
		return ret
	}).(WebTestPropertiesResponseConfigurationOutput)
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) WebTest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.WebTest
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentDataVolumeCapOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentDataVolumeCapPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentDataVolumeCapResponseOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentDataVolumeCapResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationArrayOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationResponseOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationResponseArrayOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesConfigurationOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseConfigurationOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseConfigurationPtrOutput{})
}
