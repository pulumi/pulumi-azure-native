


package extendedlocation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomLocationPropertiesAuthentication struct {
	Type  *string `pulumi:"type"`
	Value *string `pulumi:"value"`
}





type CustomLocationPropertiesAuthenticationInput interface {
	pulumi.Input

	ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput
	ToCustomLocationPropertiesAuthenticationOutputWithContext(context.Context) CustomLocationPropertiesAuthenticationOutput
}

type CustomLocationPropertiesAuthenticationArgs struct {
	Type  pulumi.StringPtrInput `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (CustomLocationPropertiesAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput {
	return i.ToCustomLocationPropertiesAuthenticationOutputWithContext(context.Background())
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationOutput)
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return i.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationOutput).ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx)
}









type CustomLocationPropertiesAuthenticationPtrInput interface {
	pulumi.Input

	ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput
	ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Context) CustomLocationPropertiesAuthenticationPtrOutput
}

type customLocationPropertiesAuthenticationPtrType CustomLocationPropertiesAuthenticationArgs

func CustomLocationPropertiesAuthenticationPtr(v *CustomLocationPropertiesAuthenticationArgs) CustomLocationPropertiesAuthenticationPtrInput {
	return (*customLocationPropertiesAuthenticationPtrType)(v)
}

func (*customLocationPropertiesAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (i *customLocationPropertiesAuthenticationPtrType) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return i.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (i *customLocationPropertiesAuthenticationPtrType) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationPtrOutput)
}

type CustomLocationPropertiesAuthenticationOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return o.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomLocationPropertiesAuthentication) *CustomLocationPropertiesAuthentication {
		return &v
	}).(CustomLocationPropertiesAuthenticationPtrOutput)
}

func (o CustomLocationPropertiesAuthenticationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesAuthentication) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o CustomLocationPropertiesAuthenticationOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesAuthentication) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) Elem() CustomLocationPropertiesAuthenticationOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) CustomLocationPropertiesAuthentication {
		if v != nil {
			return *v
		}
		var ret CustomLocationPropertiesAuthentication
		return ret
	}).(CustomLocationPropertiesAuthenticationOutput)
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesResponseAuthentication struct {
	Type *string `pulumi:"type"`
}

type CustomLocationPropertiesResponseAuthenticationOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesResponseAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesResponseAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToCustomLocationPropertiesResponseAuthenticationOutput() CustomLocationPropertiesResponseAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToCustomLocationPropertiesResponseAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesResponseAuthentication) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesResponseAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesResponseAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesResponseAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) ToCustomLocationPropertiesResponseAuthenticationPtrOutput() CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) Elem() CustomLocationPropertiesResponseAuthenticationOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesResponseAuthentication) CustomLocationPropertiesResponseAuthentication {
		if v != nil {
			return *v
		}
		var ret CustomLocationPropertiesResponseAuthentication
		return ret
	}).(CustomLocationPropertiesResponseAuthenticationOutput)
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesResponseAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceSyncRulePropertiesResponseSelector struct {
	MatchLabels map[string]string `pulumi:"matchLabels"`
}

type ResourceSyncRulePropertiesResponseSelectorOutput struct{ *pulumi.OutputState }

func (ResourceSyncRulePropertiesResponseSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSyncRulePropertiesResponseSelector)(nil)).Elem()
}

func (o ResourceSyncRulePropertiesResponseSelectorOutput) ToResourceSyncRulePropertiesResponseSelectorOutput() ResourceSyncRulePropertiesResponseSelectorOutput {
	return o
}

func (o ResourceSyncRulePropertiesResponseSelectorOutput) ToResourceSyncRulePropertiesResponseSelectorOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesResponseSelectorOutput {
	return o
}

func (o ResourceSyncRulePropertiesResponseSelectorOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceSyncRulePropertiesResponseSelector) map[string]string { return v.MatchLabels }).(pulumi.StringMapOutput)
}

type ResourceSyncRulePropertiesResponseSelectorPtrOutput struct{ *pulumi.OutputState }

func (ResourceSyncRulePropertiesResponseSelectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSyncRulePropertiesResponseSelector)(nil)).Elem()
}

func (o ResourceSyncRulePropertiesResponseSelectorPtrOutput) ToResourceSyncRulePropertiesResponseSelectorPtrOutput() ResourceSyncRulePropertiesResponseSelectorPtrOutput {
	return o
}

func (o ResourceSyncRulePropertiesResponseSelectorPtrOutput) ToResourceSyncRulePropertiesResponseSelectorPtrOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesResponseSelectorPtrOutput {
	return o
}

func (o ResourceSyncRulePropertiesResponseSelectorPtrOutput) Elem() ResourceSyncRulePropertiesResponseSelectorOutput {
	return o.ApplyT(func(v *ResourceSyncRulePropertiesResponseSelector) ResourceSyncRulePropertiesResponseSelector {
		if v != nil {
			return *v
		}
		var ret ResourceSyncRulePropertiesResponseSelector
		return ret
	}).(ResourceSyncRulePropertiesResponseSelectorOutput)
}

func (o ResourceSyncRulePropertiesResponseSelectorPtrOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceSyncRulePropertiesResponseSelector) map[string]string {
		if v == nil {
			return nil
		}
		return v.MatchLabels
	}).(pulumi.StringMapOutput)
}

type ResourceSyncRulePropertiesSelector struct {
	MatchLabels map[string]string `pulumi:"matchLabels"`
}





type ResourceSyncRulePropertiesSelectorInput interface {
	pulumi.Input

	ToResourceSyncRulePropertiesSelectorOutput() ResourceSyncRulePropertiesSelectorOutput
	ToResourceSyncRulePropertiesSelectorOutputWithContext(context.Context) ResourceSyncRulePropertiesSelectorOutput
}

type ResourceSyncRulePropertiesSelectorArgs struct {
	MatchLabels pulumi.StringMapInput `pulumi:"matchLabels"`
}

func (ResourceSyncRulePropertiesSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSyncRulePropertiesSelector)(nil)).Elem()
}

func (i ResourceSyncRulePropertiesSelectorArgs) ToResourceSyncRulePropertiesSelectorOutput() ResourceSyncRulePropertiesSelectorOutput {
	return i.ToResourceSyncRulePropertiesSelectorOutputWithContext(context.Background())
}

func (i ResourceSyncRulePropertiesSelectorArgs) ToResourceSyncRulePropertiesSelectorOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSyncRulePropertiesSelectorOutput)
}

func (i ResourceSyncRulePropertiesSelectorArgs) ToResourceSyncRulePropertiesSelectorPtrOutput() ResourceSyncRulePropertiesSelectorPtrOutput {
	return i.ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(context.Background())
}

func (i ResourceSyncRulePropertiesSelectorArgs) ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSyncRulePropertiesSelectorOutput).ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx)
}









type ResourceSyncRulePropertiesSelectorPtrInput interface {
	pulumi.Input

	ToResourceSyncRulePropertiesSelectorPtrOutput() ResourceSyncRulePropertiesSelectorPtrOutput
	ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(context.Context) ResourceSyncRulePropertiesSelectorPtrOutput
}

type resourceSyncRulePropertiesSelectorPtrType ResourceSyncRulePropertiesSelectorArgs

func ResourceSyncRulePropertiesSelectorPtr(v *ResourceSyncRulePropertiesSelectorArgs) ResourceSyncRulePropertiesSelectorPtrInput {
	return (*resourceSyncRulePropertiesSelectorPtrType)(v)
}

func (*resourceSyncRulePropertiesSelectorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSyncRulePropertiesSelector)(nil)).Elem()
}

func (i *resourceSyncRulePropertiesSelectorPtrType) ToResourceSyncRulePropertiesSelectorPtrOutput() ResourceSyncRulePropertiesSelectorPtrOutput {
	return i.ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(context.Background())
}

func (i *resourceSyncRulePropertiesSelectorPtrType) ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSyncRulePropertiesSelectorPtrOutput)
}

type ResourceSyncRulePropertiesSelectorOutput struct{ *pulumi.OutputState }

func (ResourceSyncRulePropertiesSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSyncRulePropertiesSelector)(nil)).Elem()
}

func (o ResourceSyncRulePropertiesSelectorOutput) ToResourceSyncRulePropertiesSelectorOutput() ResourceSyncRulePropertiesSelectorOutput {
	return o
}

func (o ResourceSyncRulePropertiesSelectorOutput) ToResourceSyncRulePropertiesSelectorOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorOutput {
	return o
}

func (o ResourceSyncRulePropertiesSelectorOutput) ToResourceSyncRulePropertiesSelectorPtrOutput() ResourceSyncRulePropertiesSelectorPtrOutput {
	return o.ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(context.Background())
}

func (o ResourceSyncRulePropertiesSelectorOutput) ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSyncRulePropertiesSelector) *ResourceSyncRulePropertiesSelector {
		return &v
	}).(ResourceSyncRulePropertiesSelectorPtrOutput)
}

func (o ResourceSyncRulePropertiesSelectorOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceSyncRulePropertiesSelector) map[string]string { return v.MatchLabels }).(pulumi.StringMapOutput)
}

type ResourceSyncRulePropertiesSelectorPtrOutput struct{ *pulumi.OutputState }

func (ResourceSyncRulePropertiesSelectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSyncRulePropertiesSelector)(nil)).Elem()
}

func (o ResourceSyncRulePropertiesSelectorPtrOutput) ToResourceSyncRulePropertiesSelectorPtrOutput() ResourceSyncRulePropertiesSelectorPtrOutput {
	return o
}

func (o ResourceSyncRulePropertiesSelectorPtrOutput) ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorPtrOutput {
	return o
}

func (o ResourceSyncRulePropertiesSelectorPtrOutput) Elem() ResourceSyncRulePropertiesSelectorOutput {
	return o.ApplyT(func(v *ResourceSyncRulePropertiesSelector) ResourceSyncRulePropertiesSelector {
		if v != nil {
			return *v
		}
		var ret ResourceSyncRulePropertiesSelector
		return ret
	}).(ResourceSyncRulePropertiesSelectorOutput)
}

func (o ResourceSyncRulePropertiesSelectorPtrOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceSyncRulePropertiesSelector) map[string]string {
		if v == nil {
			return nil
		}
		return v.MatchLabels
	}).(pulumi.StringMapOutput)
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
	pulumi.RegisterOutputType(CustomLocationPropertiesAuthenticationOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesResponseAuthenticationOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesResponseAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(ResourceSyncRulePropertiesResponseSelectorOutput{})
	pulumi.RegisterOutputType(ResourceSyncRulePropertiesResponseSelectorPtrOutput{})
	pulumi.RegisterOutputType(ResourceSyncRulePropertiesSelectorOutput{})
	pulumi.RegisterOutputType(ResourceSyncRulePropertiesSelectorPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
