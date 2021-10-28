


package v20181001

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

type SignalRCorsSettings struct {
	AllowedOrigins []string `pulumi:"allowedOrigins"`
}





type SignalRCorsSettingsInput interface {
	pulumi.Input

	ToSignalRCorsSettingsOutput() SignalRCorsSettingsOutput
	ToSignalRCorsSettingsOutputWithContext(context.Context) SignalRCorsSettingsOutput
}

type SignalRCorsSettingsArgs struct {
	AllowedOrigins pulumi.StringArrayInput `pulumi:"allowedOrigins"`
}

func (SignalRCorsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRCorsSettings)(nil)).Elem()
}

func (i SignalRCorsSettingsArgs) ToSignalRCorsSettingsOutput() SignalRCorsSettingsOutput {
	return i.ToSignalRCorsSettingsOutputWithContext(context.Background())
}

func (i SignalRCorsSettingsArgs) ToSignalRCorsSettingsOutputWithContext(ctx context.Context) SignalRCorsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCorsSettingsOutput)
}

func (i SignalRCorsSettingsArgs) ToSignalRCorsSettingsPtrOutput() SignalRCorsSettingsPtrOutput {
	return i.ToSignalRCorsSettingsPtrOutputWithContext(context.Background())
}

func (i SignalRCorsSettingsArgs) ToSignalRCorsSettingsPtrOutputWithContext(ctx context.Context) SignalRCorsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCorsSettingsOutput).ToSignalRCorsSettingsPtrOutputWithContext(ctx)
}









type SignalRCorsSettingsPtrInput interface {
	pulumi.Input

	ToSignalRCorsSettingsPtrOutput() SignalRCorsSettingsPtrOutput
	ToSignalRCorsSettingsPtrOutputWithContext(context.Context) SignalRCorsSettingsPtrOutput
}

type signalRCorsSettingsPtrType SignalRCorsSettingsArgs

func SignalRCorsSettingsPtr(v *SignalRCorsSettingsArgs) SignalRCorsSettingsPtrInput {
	return (*signalRCorsSettingsPtrType)(v)
}

func (*signalRCorsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCorsSettings)(nil)).Elem()
}

func (i *signalRCorsSettingsPtrType) ToSignalRCorsSettingsPtrOutput() SignalRCorsSettingsPtrOutput {
	return i.ToSignalRCorsSettingsPtrOutputWithContext(context.Background())
}

func (i *signalRCorsSettingsPtrType) ToSignalRCorsSettingsPtrOutputWithContext(ctx context.Context) SignalRCorsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCorsSettingsPtrOutput)
}

type SignalRCorsSettingsOutput struct{ *pulumi.OutputState }

func (SignalRCorsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRCorsSettings)(nil)).Elem()
}

func (o SignalRCorsSettingsOutput) ToSignalRCorsSettingsOutput() SignalRCorsSettingsOutput {
	return o
}

func (o SignalRCorsSettingsOutput) ToSignalRCorsSettingsOutputWithContext(ctx context.Context) SignalRCorsSettingsOutput {
	return o
}

func (o SignalRCorsSettingsOutput) ToSignalRCorsSettingsPtrOutput() SignalRCorsSettingsPtrOutput {
	return o.ToSignalRCorsSettingsPtrOutputWithContext(context.Background())
}

func (o SignalRCorsSettingsOutput) ToSignalRCorsSettingsPtrOutputWithContext(ctx context.Context) SignalRCorsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignalRCorsSettings) *SignalRCorsSettings {
		return &v
	}).(SignalRCorsSettingsPtrOutput)
}

func (o SignalRCorsSettingsOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SignalRCorsSettings) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

type SignalRCorsSettingsPtrOutput struct{ *pulumi.OutputState }

func (SignalRCorsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCorsSettings)(nil)).Elem()
}

func (o SignalRCorsSettingsPtrOutput) ToSignalRCorsSettingsPtrOutput() SignalRCorsSettingsPtrOutput {
	return o
}

func (o SignalRCorsSettingsPtrOutput) ToSignalRCorsSettingsPtrOutputWithContext(ctx context.Context) SignalRCorsSettingsPtrOutput {
	return o
}

func (o SignalRCorsSettingsPtrOutput) Elem() SignalRCorsSettingsOutput {
	return o.ApplyT(func(v *SignalRCorsSettings) SignalRCorsSettings {
		if v != nil {
			return *v
		}
		var ret SignalRCorsSettings
		return ret
	}).(SignalRCorsSettingsOutput)
}

func (o SignalRCorsSettingsPtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SignalRCorsSettings) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

type SignalRCorsSettingsResponse struct {
	AllowedOrigins []string `pulumi:"allowedOrigins"`
}





type SignalRCorsSettingsResponseInput interface {
	pulumi.Input

	ToSignalRCorsSettingsResponseOutput() SignalRCorsSettingsResponseOutput
	ToSignalRCorsSettingsResponseOutputWithContext(context.Context) SignalRCorsSettingsResponseOutput
}

type SignalRCorsSettingsResponseArgs struct {
	AllowedOrigins pulumi.StringArrayInput `pulumi:"allowedOrigins"`
}

func (SignalRCorsSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRCorsSettingsResponse)(nil)).Elem()
}

func (i SignalRCorsSettingsResponseArgs) ToSignalRCorsSettingsResponseOutput() SignalRCorsSettingsResponseOutput {
	return i.ToSignalRCorsSettingsResponseOutputWithContext(context.Background())
}

func (i SignalRCorsSettingsResponseArgs) ToSignalRCorsSettingsResponseOutputWithContext(ctx context.Context) SignalRCorsSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCorsSettingsResponseOutput)
}

func (i SignalRCorsSettingsResponseArgs) ToSignalRCorsSettingsResponsePtrOutput() SignalRCorsSettingsResponsePtrOutput {
	return i.ToSignalRCorsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i SignalRCorsSettingsResponseArgs) ToSignalRCorsSettingsResponsePtrOutputWithContext(ctx context.Context) SignalRCorsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCorsSettingsResponseOutput).ToSignalRCorsSettingsResponsePtrOutputWithContext(ctx)
}









type SignalRCorsSettingsResponsePtrInput interface {
	pulumi.Input

	ToSignalRCorsSettingsResponsePtrOutput() SignalRCorsSettingsResponsePtrOutput
	ToSignalRCorsSettingsResponsePtrOutputWithContext(context.Context) SignalRCorsSettingsResponsePtrOutput
}

type signalRCorsSettingsResponsePtrType SignalRCorsSettingsResponseArgs

func SignalRCorsSettingsResponsePtr(v *SignalRCorsSettingsResponseArgs) SignalRCorsSettingsResponsePtrInput {
	return (*signalRCorsSettingsResponsePtrType)(v)
}

func (*signalRCorsSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCorsSettingsResponse)(nil)).Elem()
}

func (i *signalRCorsSettingsResponsePtrType) ToSignalRCorsSettingsResponsePtrOutput() SignalRCorsSettingsResponsePtrOutput {
	return i.ToSignalRCorsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *signalRCorsSettingsResponsePtrType) ToSignalRCorsSettingsResponsePtrOutputWithContext(ctx context.Context) SignalRCorsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCorsSettingsResponsePtrOutput)
}

type SignalRCorsSettingsResponseOutput struct{ *pulumi.OutputState }

func (SignalRCorsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRCorsSettingsResponse)(nil)).Elem()
}

func (o SignalRCorsSettingsResponseOutput) ToSignalRCorsSettingsResponseOutput() SignalRCorsSettingsResponseOutput {
	return o
}

func (o SignalRCorsSettingsResponseOutput) ToSignalRCorsSettingsResponseOutputWithContext(ctx context.Context) SignalRCorsSettingsResponseOutput {
	return o
}

func (o SignalRCorsSettingsResponseOutput) ToSignalRCorsSettingsResponsePtrOutput() SignalRCorsSettingsResponsePtrOutput {
	return o.ToSignalRCorsSettingsResponsePtrOutputWithContext(context.Background())
}

func (o SignalRCorsSettingsResponseOutput) ToSignalRCorsSettingsResponsePtrOutputWithContext(ctx context.Context) SignalRCorsSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignalRCorsSettingsResponse) *SignalRCorsSettingsResponse {
		return &v
	}).(SignalRCorsSettingsResponsePtrOutput)
}

func (o SignalRCorsSettingsResponseOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SignalRCorsSettingsResponse) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

type SignalRCorsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SignalRCorsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCorsSettingsResponse)(nil)).Elem()
}

func (o SignalRCorsSettingsResponsePtrOutput) ToSignalRCorsSettingsResponsePtrOutput() SignalRCorsSettingsResponsePtrOutput {
	return o
}

func (o SignalRCorsSettingsResponsePtrOutput) ToSignalRCorsSettingsResponsePtrOutputWithContext(ctx context.Context) SignalRCorsSettingsResponsePtrOutput {
	return o
}

func (o SignalRCorsSettingsResponsePtrOutput) Elem() SignalRCorsSettingsResponseOutput {
	return o.ApplyT(func(v *SignalRCorsSettingsResponse) SignalRCorsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret SignalRCorsSettingsResponse
		return ret
	}).(SignalRCorsSettingsResponseOutput)
}

func (o SignalRCorsSettingsResponsePtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SignalRCorsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

type SignalRCreateOrUpdateProperties struct {
	Cors           *SignalRCorsSettings `pulumi:"cors"`
	Features       []SignalRFeature     `pulumi:"features"`
	HostNamePrefix *string              `pulumi:"hostNamePrefix"`
}





type SignalRCreateOrUpdatePropertiesInput interface {
	pulumi.Input

	ToSignalRCreateOrUpdatePropertiesOutput() SignalRCreateOrUpdatePropertiesOutput
	ToSignalRCreateOrUpdatePropertiesOutputWithContext(context.Context) SignalRCreateOrUpdatePropertiesOutput
}

type SignalRCreateOrUpdatePropertiesArgs struct {
	Cors           SignalRCorsSettingsPtrInput `pulumi:"cors"`
	Features       SignalRFeatureArrayInput    `pulumi:"features"`
	HostNamePrefix pulumi.StringPtrInput       `pulumi:"hostNamePrefix"`
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

func (o SignalRCreateOrUpdatePropertiesOutput) Cors() SignalRCorsSettingsPtrOutput {
	return o.ApplyT(func(v SignalRCreateOrUpdateProperties) *SignalRCorsSettings { return v.Cors }).(SignalRCorsSettingsPtrOutput)
}

func (o SignalRCreateOrUpdatePropertiesOutput) Features() SignalRFeatureArrayOutput {
	return o.ApplyT(func(v SignalRCreateOrUpdateProperties) []SignalRFeature { return v.Features }).(SignalRFeatureArrayOutput)
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

func (o SignalRCreateOrUpdatePropertiesPtrOutput) Cors() SignalRCorsSettingsPtrOutput {
	return o.ApplyT(func(v *SignalRCreateOrUpdateProperties) *SignalRCorsSettings {
		if v == nil {
			return nil
		}
		return v.Cors
	}).(SignalRCorsSettingsPtrOutput)
}

func (o SignalRCreateOrUpdatePropertiesPtrOutput) Features() SignalRFeatureArrayOutput {
	return o.ApplyT(func(v *SignalRCreateOrUpdateProperties) []SignalRFeature {
		if v == nil {
			return nil
		}
		return v.Features
	}).(SignalRFeatureArrayOutput)
}

func (o SignalRCreateOrUpdatePropertiesPtrOutput) HostNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalRCreateOrUpdateProperties) *string {
		if v == nil {
			return nil
		}
		return v.HostNamePrefix
	}).(pulumi.StringPtrOutput)
}

type SignalRFeature struct {
	Flag       string            `pulumi:"flag"`
	Properties map[string]string `pulumi:"properties"`
	Value      string            `pulumi:"value"`
}





type SignalRFeatureInput interface {
	pulumi.Input

	ToSignalRFeatureOutput() SignalRFeatureOutput
	ToSignalRFeatureOutputWithContext(context.Context) SignalRFeatureOutput
}

type SignalRFeatureArgs struct {
	Flag       pulumi.StringInput    `pulumi:"flag"`
	Properties pulumi.StringMapInput `pulumi:"properties"`
	Value      pulumi.StringInput    `pulumi:"value"`
}

func (SignalRFeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRFeature)(nil)).Elem()
}

func (i SignalRFeatureArgs) ToSignalRFeatureOutput() SignalRFeatureOutput {
	return i.ToSignalRFeatureOutputWithContext(context.Background())
}

func (i SignalRFeatureArgs) ToSignalRFeatureOutputWithContext(ctx context.Context) SignalRFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRFeatureOutput)
}





type SignalRFeatureArrayInput interface {
	pulumi.Input

	ToSignalRFeatureArrayOutput() SignalRFeatureArrayOutput
	ToSignalRFeatureArrayOutputWithContext(context.Context) SignalRFeatureArrayOutput
}

type SignalRFeatureArray []SignalRFeatureInput

func (SignalRFeatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SignalRFeature)(nil)).Elem()
}

func (i SignalRFeatureArray) ToSignalRFeatureArrayOutput() SignalRFeatureArrayOutput {
	return i.ToSignalRFeatureArrayOutputWithContext(context.Background())
}

func (i SignalRFeatureArray) ToSignalRFeatureArrayOutputWithContext(ctx context.Context) SignalRFeatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRFeatureArrayOutput)
}

type SignalRFeatureOutput struct{ *pulumi.OutputState }

func (SignalRFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRFeature)(nil)).Elem()
}

func (o SignalRFeatureOutput) ToSignalRFeatureOutput() SignalRFeatureOutput {
	return o
}

func (o SignalRFeatureOutput) ToSignalRFeatureOutputWithContext(ctx context.Context) SignalRFeatureOutput {
	return o
}

func (o SignalRFeatureOutput) Flag() pulumi.StringOutput {
	return o.ApplyT(func(v SignalRFeature) string { return v.Flag }).(pulumi.StringOutput)
}

func (o SignalRFeatureOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v SignalRFeature) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o SignalRFeatureOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SignalRFeature) string { return v.Value }).(pulumi.StringOutput)
}

type SignalRFeatureArrayOutput struct{ *pulumi.OutputState }

func (SignalRFeatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SignalRFeature)(nil)).Elem()
}

func (o SignalRFeatureArrayOutput) ToSignalRFeatureArrayOutput() SignalRFeatureArrayOutput {
	return o
}

func (o SignalRFeatureArrayOutput) ToSignalRFeatureArrayOutputWithContext(ctx context.Context) SignalRFeatureArrayOutput {
	return o
}

func (o SignalRFeatureArrayOutput) Index(i pulumi.IntInput) SignalRFeatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SignalRFeature {
		return vs[0].([]SignalRFeature)[vs[1].(int)]
	}).(SignalRFeatureOutput)
}

type SignalRFeatureResponse struct {
	Flag       string            `pulumi:"flag"`
	Properties map[string]string `pulumi:"properties"`
	Value      string            `pulumi:"value"`
}





type SignalRFeatureResponseInput interface {
	pulumi.Input

	ToSignalRFeatureResponseOutput() SignalRFeatureResponseOutput
	ToSignalRFeatureResponseOutputWithContext(context.Context) SignalRFeatureResponseOutput
}

type SignalRFeatureResponseArgs struct {
	Flag       pulumi.StringInput    `pulumi:"flag"`
	Properties pulumi.StringMapInput `pulumi:"properties"`
	Value      pulumi.StringInput    `pulumi:"value"`
}

func (SignalRFeatureResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRFeatureResponse)(nil)).Elem()
}

func (i SignalRFeatureResponseArgs) ToSignalRFeatureResponseOutput() SignalRFeatureResponseOutput {
	return i.ToSignalRFeatureResponseOutputWithContext(context.Background())
}

func (i SignalRFeatureResponseArgs) ToSignalRFeatureResponseOutputWithContext(ctx context.Context) SignalRFeatureResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRFeatureResponseOutput)
}





type SignalRFeatureResponseArrayInput interface {
	pulumi.Input

	ToSignalRFeatureResponseArrayOutput() SignalRFeatureResponseArrayOutput
	ToSignalRFeatureResponseArrayOutputWithContext(context.Context) SignalRFeatureResponseArrayOutput
}

type SignalRFeatureResponseArray []SignalRFeatureResponseInput

func (SignalRFeatureResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SignalRFeatureResponse)(nil)).Elem()
}

func (i SignalRFeatureResponseArray) ToSignalRFeatureResponseArrayOutput() SignalRFeatureResponseArrayOutput {
	return i.ToSignalRFeatureResponseArrayOutputWithContext(context.Background())
}

func (i SignalRFeatureResponseArray) ToSignalRFeatureResponseArrayOutputWithContext(ctx context.Context) SignalRFeatureResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRFeatureResponseArrayOutput)
}

type SignalRFeatureResponseOutput struct{ *pulumi.OutputState }

func (SignalRFeatureResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRFeatureResponse)(nil)).Elem()
}

func (o SignalRFeatureResponseOutput) ToSignalRFeatureResponseOutput() SignalRFeatureResponseOutput {
	return o
}

func (o SignalRFeatureResponseOutput) ToSignalRFeatureResponseOutputWithContext(ctx context.Context) SignalRFeatureResponseOutput {
	return o
}

func (o SignalRFeatureResponseOutput) Flag() pulumi.StringOutput {
	return o.ApplyT(func(v SignalRFeatureResponse) string { return v.Flag }).(pulumi.StringOutput)
}

func (o SignalRFeatureResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v SignalRFeatureResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o SignalRFeatureResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SignalRFeatureResponse) string { return v.Value }).(pulumi.StringOutput)
}

type SignalRFeatureResponseArrayOutput struct{ *pulumi.OutputState }

func (SignalRFeatureResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SignalRFeatureResponse)(nil)).Elem()
}

func (o SignalRFeatureResponseArrayOutput) ToSignalRFeatureResponseArrayOutput() SignalRFeatureResponseArrayOutput {
	return o
}

func (o SignalRFeatureResponseArrayOutput) ToSignalRFeatureResponseArrayOutputWithContext(ctx context.Context) SignalRFeatureResponseArrayOutput {
	return o
}

func (o SignalRFeatureResponseArrayOutput) Index(i pulumi.IntInput) SignalRFeatureResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SignalRFeatureResponse {
		return vs[0].([]SignalRFeatureResponse)[vs[1].(int)]
	}).(SignalRFeatureResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SignalRCorsSettingsOutput{})
	pulumi.RegisterOutputType(SignalRCorsSettingsPtrOutput{})
	pulumi.RegisterOutputType(SignalRCorsSettingsResponseOutput{})
	pulumi.RegisterOutputType(SignalRCorsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SignalRCreateOrUpdatePropertiesOutput{})
	pulumi.RegisterOutputType(SignalRCreateOrUpdatePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SignalRFeatureOutput{})
	pulumi.RegisterOutputType(SignalRFeatureArrayOutput{})
	pulumi.RegisterOutputType(SignalRFeatureResponseOutput{})
	pulumi.RegisterOutputType(SignalRFeatureResponseArrayOutput{})
}
