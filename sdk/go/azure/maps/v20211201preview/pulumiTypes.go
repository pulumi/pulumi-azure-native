


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CorsRule struct {
	AllowedOrigins []string `pulumi:"allowedOrigins"`
}





type CorsRuleInput interface {
	pulumi.Input

	ToCorsRuleOutput() CorsRuleOutput
	ToCorsRuleOutputWithContext(context.Context) CorsRuleOutput
}

type CorsRuleArgs struct {
	AllowedOrigins pulumi.StringArrayInput `pulumi:"allowedOrigins"`
}

func (CorsRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRule)(nil)).Elem()
}

func (i CorsRuleArgs) ToCorsRuleOutput() CorsRuleOutput {
	return i.ToCorsRuleOutputWithContext(context.Background())
}

func (i CorsRuleArgs) ToCorsRuleOutputWithContext(ctx context.Context) CorsRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRuleOutput)
}





type CorsRuleArrayInput interface {
	pulumi.Input

	ToCorsRuleArrayOutput() CorsRuleArrayOutput
	ToCorsRuleArrayOutputWithContext(context.Context) CorsRuleArrayOutput
}

type CorsRuleArray []CorsRuleInput

func (CorsRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRule)(nil)).Elem()
}

func (i CorsRuleArray) ToCorsRuleArrayOutput() CorsRuleArrayOutput {
	return i.ToCorsRuleArrayOutputWithContext(context.Background())
}

func (i CorsRuleArray) ToCorsRuleArrayOutputWithContext(ctx context.Context) CorsRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRuleArrayOutput)
}

type CorsRuleOutput struct{ *pulumi.OutputState }

func (CorsRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRule)(nil)).Elem()
}

func (o CorsRuleOutput) ToCorsRuleOutput() CorsRuleOutput {
	return o
}

func (o CorsRuleOutput) ToCorsRuleOutputWithContext(ctx context.Context) CorsRuleOutput {
	return o
}

func (o CorsRuleOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

type CorsRuleArrayOutput struct{ *pulumi.OutputState }

func (CorsRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRule)(nil)).Elem()
}

func (o CorsRuleArrayOutput) ToCorsRuleArrayOutput() CorsRuleArrayOutput {
	return o
}

func (o CorsRuleArrayOutput) ToCorsRuleArrayOutputWithContext(ctx context.Context) CorsRuleArrayOutput {
	return o
}

func (o CorsRuleArrayOutput) Index(i pulumi.IntInput) CorsRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CorsRule {
		return vs[0].([]CorsRule)[vs[1].(int)]
	}).(CorsRuleOutput)
}

type CorsRuleResponse struct {
	AllowedOrigins []string `pulumi:"allowedOrigins"`
}





type CorsRuleResponseInput interface {
	pulumi.Input

	ToCorsRuleResponseOutput() CorsRuleResponseOutput
	ToCorsRuleResponseOutputWithContext(context.Context) CorsRuleResponseOutput
}

type CorsRuleResponseArgs struct {
	AllowedOrigins pulumi.StringArrayInput `pulumi:"allowedOrigins"`
}

func (CorsRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRuleResponse)(nil)).Elem()
}

func (i CorsRuleResponseArgs) ToCorsRuleResponseOutput() CorsRuleResponseOutput {
	return i.ToCorsRuleResponseOutputWithContext(context.Background())
}

func (i CorsRuleResponseArgs) ToCorsRuleResponseOutputWithContext(ctx context.Context) CorsRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRuleResponseOutput)
}





type CorsRuleResponseArrayInput interface {
	pulumi.Input

	ToCorsRuleResponseArrayOutput() CorsRuleResponseArrayOutput
	ToCorsRuleResponseArrayOutputWithContext(context.Context) CorsRuleResponseArrayOutput
}

type CorsRuleResponseArray []CorsRuleResponseInput

func (CorsRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRuleResponse)(nil)).Elem()
}

func (i CorsRuleResponseArray) ToCorsRuleResponseArrayOutput() CorsRuleResponseArrayOutput {
	return i.ToCorsRuleResponseArrayOutputWithContext(context.Background())
}

func (i CorsRuleResponseArray) ToCorsRuleResponseArrayOutputWithContext(ctx context.Context) CorsRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRuleResponseArrayOutput)
}

type CorsRuleResponseOutput struct{ *pulumi.OutputState }

func (CorsRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRuleResponse)(nil)).Elem()
}

func (o CorsRuleResponseOutput) ToCorsRuleResponseOutput() CorsRuleResponseOutput {
	return o
}

func (o CorsRuleResponseOutput) ToCorsRuleResponseOutputWithContext(ctx context.Context) CorsRuleResponseOutput {
	return o
}

func (o CorsRuleResponseOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

type CorsRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (CorsRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRuleResponse)(nil)).Elem()
}

func (o CorsRuleResponseArrayOutput) ToCorsRuleResponseArrayOutput() CorsRuleResponseArrayOutput {
	return o
}

func (o CorsRuleResponseArrayOutput) ToCorsRuleResponseArrayOutputWithContext(ctx context.Context) CorsRuleResponseArrayOutput {
	return o
}

func (o CorsRuleResponseArrayOutput) Index(i pulumi.IntInput) CorsRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CorsRuleResponse {
		return vs[0].([]CorsRuleResponse)[vs[1].(int)]
	}).(CorsRuleResponseOutput)
}

type CorsRules struct {
	CorsRules []CorsRule `pulumi:"corsRules"`
}





type CorsRulesInput interface {
	pulumi.Input

	ToCorsRulesOutput() CorsRulesOutput
	ToCorsRulesOutputWithContext(context.Context) CorsRulesOutput
}

type CorsRulesArgs struct {
	CorsRules CorsRuleArrayInput `pulumi:"corsRules"`
}

func (CorsRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRules)(nil)).Elem()
}

func (i CorsRulesArgs) ToCorsRulesOutput() CorsRulesOutput {
	return i.ToCorsRulesOutputWithContext(context.Background())
}

func (i CorsRulesArgs) ToCorsRulesOutputWithContext(ctx context.Context) CorsRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesOutput)
}

func (i CorsRulesArgs) ToCorsRulesPtrOutput() CorsRulesPtrOutput {
	return i.ToCorsRulesPtrOutputWithContext(context.Background())
}

func (i CorsRulesArgs) ToCorsRulesPtrOutputWithContext(ctx context.Context) CorsRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesOutput).ToCorsRulesPtrOutputWithContext(ctx)
}









type CorsRulesPtrInput interface {
	pulumi.Input

	ToCorsRulesPtrOutput() CorsRulesPtrOutput
	ToCorsRulesPtrOutputWithContext(context.Context) CorsRulesPtrOutput
}

type corsRulesPtrType CorsRulesArgs

func CorsRulesPtr(v *CorsRulesArgs) CorsRulesPtrInput {
	return (*corsRulesPtrType)(v)
}

func (*corsRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsRules)(nil)).Elem()
}

func (i *corsRulesPtrType) ToCorsRulesPtrOutput() CorsRulesPtrOutput {
	return i.ToCorsRulesPtrOutputWithContext(context.Background())
}

func (i *corsRulesPtrType) ToCorsRulesPtrOutputWithContext(ctx context.Context) CorsRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesPtrOutput)
}

type CorsRulesOutput struct{ *pulumi.OutputState }

func (CorsRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRules)(nil)).Elem()
}

func (o CorsRulesOutput) ToCorsRulesOutput() CorsRulesOutput {
	return o
}

func (o CorsRulesOutput) ToCorsRulesOutputWithContext(ctx context.Context) CorsRulesOutput {
	return o
}

func (o CorsRulesOutput) ToCorsRulesPtrOutput() CorsRulesPtrOutput {
	return o.ToCorsRulesPtrOutputWithContext(context.Background())
}

func (o CorsRulesOutput) ToCorsRulesPtrOutputWithContext(ctx context.Context) CorsRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CorsRules) *CorsRules {
		return &v
	}).(CorsRulesPtrOutput)
}

func (o CorsRulesOutput) CorsRules() CorsRuleArrayOutput {
	return o.ApplyT(func(v CorsRules) []CorsRule { return v.CorsRules }).(CorsRuleArrayOutput)
}

type CorsRulesPtrOutput struct{ *pulumi.OutputState }

func (CorsRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsRules)(nil)).Elem()
}

func (o CorsRulesPtrOutput) ToCorsRulesPtrOutput() CorsRulesPtrOutput {
	return o
}

func (o CorsRulesPtrOutput) ToCorsRulesPtrOutputWithContext(ctx context.Context) CorsRulesPtrOutput {
	return o
}

func (o CorsRulesPtrOutput) Elem() CorsRulesOutput {
	return o.ApplyT(func(v *CorsRules) CorsRules {
		if v != nil {
			return *v
		}
		var ret CorsRules
		return ret
	}).(CorsRulesOutput)
}

func (o CorsRulesPtrOutput) CorsRules() CorsRuleArrayOutput {
	return o.ApplyT(func(v *CorsRules) []CorsRule {
		if v == nil {
			return nil
		}
		return v.CorsRules
	}).(CorsRuleArrayOutput)
}

type CorsRulesResponse struct {
	CorsRules []CorsRuleResponse `pulumi:"corsRules"`
}





type CorsRulesResponseInput interface {
	pulumi.Input

	ToCorsRulesResponseOutput() CorsRulesResponseOutput
	ToCorsRulesResponseOutputWithContext(context.Context) CorsRulesResponseOutput
}

type CorsRulesResponseArgs struct {
	CorsRules CorsRuleResponseArrayInput `pulumi:"corsRules"`
}

func (CorsRulesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRulesResponse)(nil)).Elem()
}

func (i CorsRulesResponseArgs) ToCorsRulesResponseOutput() CorsRulesResponseOutput {
	return i.ToCorsRulesResponseOutputWithContext(context.Background())
}

func (i CorsRulesResponseArgs) ToCorsRulesResponseOutputWithContext(ctx context.Context) CorsRulesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesResponseOutput)
}

func (i CorsRulesResponseArgs) ToCorsRulesResponsePtrOutput() CorsRulesResponsePtrOutput {
	return i.ToCorsRulesResponsePtrOutputWithContext(context.Background())
}

func (i CorsRulesResponseArgs) ToCorsRulesResponsePtrOutputWithContext(ctx context.Context) CorsRulesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesResponseOutput).ToCorsRulesResponsePtrOutputWithContext(ctx)
}









type CorsRulesResponsePtrInput interface {
	pulumi.Input

	ToCorsRulesResponsePtrOutput() CorsRulesResponsePtrOutput
	ToCorsRulesResponsePtrOutputWithContext(context.Context) CorsRulesResponsePtrOutput
}

type corsRulesResponsePtrType CorsRulesResponseArgs

func CorsRulesResponsePtr(v *CorsRulesResponseArgs) CorsRulesResponsePtrInput {
	return (*corsRulesResponsePtrType)(v)
}

func (*corsRulesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsRulesResponse)(nil)).Elem()
}

func (i *corsRulesResponsePtrType) ToCorsRulesResponsePtrOutput() CorsRulesResponsePtrOutput {
	return i.ToCorsRulesResponsePtrOutputWithContext(context.Background())
}

func (i *corsRulesResponsePtrType) ToCorsRulesResponsePtrOutputWithContext(ctx context.Context) CorsRulesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesResponsePtrOutput)
}

type CorsRulesResponseOutput struct{ *pulumi.OutputState }

func (CorsRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRulesResponse)(nil)).Elem()
}

func (o CorsRulesResponseOutput) ToCorsRulesResponseOutput() CorsRulesResponseOutput {
	return o
}

func (o CorsRulesResponseOutput) ToCorsRulesResponseOutputWithContext(ctx context.Context) CorsRulesResponseOutput {
	return o
}

func (o CorsRulesResponseOutput) ToCorsRulesResponsePtrOutput() CorsRulesResponsePtrOutput {
	return o.ToCorsRulesResponsePtrOutputWithContext(context.Background())
}

func (o CorsRulesResponseOutput) ToCorsRulesResponsePtrOutputWithContext(ctx context.Context) CorsRulesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CorsRulesResponse) *CorsRulesResponse {
		return &v
	}).(CorsRulesResponsePtrOutput)
}

func (o CorsRulesResponseOutput) CorsRules() CorsRuleResponseArrayOutput {
	return o.ApplyT(func(v CorsRulesResponse) []CorsRuleResponse { return v.CorsRules }).(CorsRuleResponseArrayOutput)
}

type CorsRulesResponsePtrOutput struct{ *pulumi.OutputState }

func (CorsRulesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsRulesResponse)(nil)).Elem()
}

func (o CorsRulesResponsePtrOutput) ToCorsRulesResponsePtrOutput() CorsRulesResponsePtrOutput {
	return o
}

func (o CorsRulesResponsePtrOutput) ToCorsRulesResponsePtrOutputWithContext(ctx context.Context) CorsRulesResponsePtrOutput {
	return o
}

func (o CorsRulesResponsePtrOutput) Elem() CorsRulesResponseOutput {
	return o.ApplyT(func(v *CorsRulesResponse) CorsRulesResponse {
		if v != nil {
			return *v
		}
		var ret CorsRulesResponse
		return ret
	}).(CorsRulesResponseOutput)
}

func (o CorsRulesResponsePtrOutput) CorsRules() CorsRuleResponseArrayOutput {
	return o.ApplyT(func(v *CorsRulesResponse) []CorsRuleResponse {
		if v == nil {
			return nil
		}
		return v.CorsRules
	}).(CorsRuleResponseArrayOutput)
}

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

type LinkedResource struct {
	Id         string `pulumi:"id"`
	UniqueName string `pulumi:"uniqueName"`
}





type LinkedResourceInput interface {
	pulumi.Input

	ToLinkedResourceOutput() LinkedResourceOutput
	ToLinkedResourceOutputWithContext(context.Context) LinkedResourceOutput
}

type LinkedResourceArgs struct {
	Id         pulumi.StringInput `pulumi:"id"`
	UniqueName pulumi.StringInput `pulumi:"uniqueName"`
}

func (LinkedResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedResource)(nil)).Elem()
}

func (i LinkedResourceArgs) ToLinkedResourceOutput() LinkedResourceOutput {
	return i.ToLinkedResourceOutputWithContext(context.Background())
}

func (i LinkedResourceArgs) ToLinkedResourceOutputWithContext(ctx context.Context) LinkedResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedResourceOutput)
}





type LinkedResourceArrayInput interface {
	pulumi.Input

	ToLinkedResourceArrayOutput() LinkedResourceArrayOutput
	ToLinkedResourceArrayOutputWithContext(context.Context) LinkedResourceArrayOutput
}

type LinkedResourceArray []LinkedResourceInput

func (LinkedResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedResource)(nil)).Elem()
}

func (i LinkedResourceArray) ToLinkedResourceArrayOutput() LinkedResourceArrayOutput {
	return i.ToLinkedResourceArrayOutputWithContext(context.Background())
}

func (i LinkedResourceArray) ToLinkedResourceArrayOutputWithContext(ctx context.Context) LinkedResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedResourceArrayOutput)
}

type LinkedResourceOutput struct{ *pulumi.OutputState }

func (LinkedResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedResource)(nil)).Elem()
}

func (o LinkedResourceOutput) ToLinkedResourceOutput() LinkedResourceOutput {
	return o
}

func (o LinkedResourceOutput) ToLinkedResourceOutputWithContext(ctx context.Context) LinkedResourceOutput {
	return o
}

func (o LinkedResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o LinkedResourceOutput) UniqueName() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedResource) string { return v.UniqueName }).(pulumi.StringOutput)
}

type LinkedResourceArrayOutput struct{ *pulumi.OutputState }

func (LinkedResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedResource)(nil)).Elem()
}

func (o LinkedResourceArrayOutput) ToLinkedResourceArrayOutput() LinkedResourceArrayOutput {
	return o
}

func (o LinkedResourceArrayOutput) ToLinkedResourceArrayOutputWithContext(ctx context.Context) LinkedResourceArrayOutput {
	return o
}

func (o LinkedResourceArrayOutput) Index(i pulumi.IntInput) LinkedResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedResource {
		return vs[0].([]LinkedResource)[vs[1].(int)]
	}).(LinkedResourceOutput)
}

type LinkedResourceResponse struct {
	Id         string `pulumi:"id"`
	UniqueName string `pulumi:"uniqueName"`
}





type LinkedResourceResponseInput interface {
	pulumi.Input

	ToLinkedResourceResponseOutput() LinkedResourceResponseOutput
	ToLinkedResourceResponseOutputWithContext(context.Context) LinkedResourceResponseOutput
}

type LinkedResourceResponseArgs struct {
	Id         pulumi.StringInput `pulumi:"id"`
	UniqueName pulumi.StringInput `pulumi:"uniqueName"`
}

func (LinkedResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedResourceResponse)(nil)).Elem()
}

func (i LinkedResourceResponseArgs) ToLinkedResourceResponseOutput() LinkedResourceResponseOutput {
	return i.ToLinkedResourceResponseOutputWithContext(context.Background())
}

func (i LinkedResourceResponseArgs) ToLinkedResourceResponseOutputWithContext(ctx context.Context) LinkedResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedResourceResponseOutput)
}





type LinkedResourceResponseArrayInput interface {
	pulumi.Input

	ToLinkedResourceResponseArrayOutput() LinkedResourceResponseArrayOutput
	ToLinkedResourceResponseArrayOutputWithContext(context.Context) LinkedResourceResponseArrayOutput
}

type LinkedResourceResponseArray []LinkedResourceResponseInput

func (LinkedResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedResourceResponse)(nil)).Elem()
}

func (i LinkedResourceResponseArray) ToLinkedResourceResponseArrayOutput() LinkedResourceResponseArrayOutput {
	return i.ToLinkedResourceResponseArrayOutputWithContext(context.Background())
}

func (i LinkedResourceResponseArray) ToLinkedResourceResponseArrayOutputWithContext(ctx context.Context) LinkedResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedResourceResponseArrayOutput)
}

type LinkedResourceResponseOutput struct{ *pulumi.OutputState }

func (LinkedResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedResourceResponse)(nil)).Elem()
}

func (o LinkedResourceResponseOutput) ToLinkedResourceResponseOutput() LinkedResourceResponseOutput {
	return o
}

func (o LinkedResourceResponseOutput) ToLinkedResourceResponseOutputWithContext(ctx context.Context) LinkedResourceResponseOutput {
	return o
}

func (o LinkedResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o LinkedResourceResponseOutput) UniqueName() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedResourceResponse) string { return v.UniqueName }).(pulumi.StringOutput)
}

type LinkedResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (LinkedResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedResourceResponse)(nil)).Elem()
}

func (o LinkedResourceResponseArrayOutput) ToLinkedResourceResponseArrayOutput() LinkedResourceResponseArrayOutput {
	return o
}

func (o LinkedResourceResponseArrayOutput) ToLinkedResourceResponseArrayOutputWithContext(ctx context.Context) LinkedResourceResponseArrayOutput {
	return o
}

func (o LinkedResourceResponseArrayOutput) Index(i pulumi.IntInput) LinkedResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedResourceResponse {
		return vs[0].([]LinkedResourceResponse)[vs[1].(int)]
	}).(LinkedResourceResponseOutput)
}

type ManagedServiceIdentity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                                          `pulumi:"principalId"`
	TenantId               string                                                          `pulumi:"tenantId"`
	Type                   *string                                                         `pulumi:"type"`
	UserAssignedIdentities map[string]ManagedServiceIdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityResponseInput interface {
	pulumi.Input

	ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput
	ToManagedServiceIdentityResponseOutputWithContext(context.Context) ManagedServiceIdentityResponseOutput
}

type ManagedServiceIdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                                           `pulumi:"principalId"`
	TenantId               pulumi.StringInput                                           `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                                        `pulumi:"type"`
	UserAssignedIdentities ManagedServiceIdentityResponseUserAssignedIdentitiesMapInput `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return i.ToManagedServiceIdentityResponseOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponseOutput)
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return i.ToManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponseOutput).ToManagedServiceIdentityResponsePtrOutputWithContext(ctx)
}









type ManagedServiceIdentityResponsePtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput
	ToManagedServiceIdentityResponsePtrOutputWithContext(context.Context) ManagedServiceIdentityResponsePtrOutput
}

type managedServiceIdentityResponsePtrType ManagedServiceIdentityResponseArgs

func ManagedServiceIdentityResponsePtr(v *ManagedServiceIdentityResponseArgs) ManagedServiceIdentityResponsePtrInput {
	return (*managedServiceIdentityResponsePtrType)(v)
}

func (*managedServiceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (i *managedServiceIdentityResponsePtrType) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return i.ToManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityResponsePtrType) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponsePtrOutput)
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o.ToManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentityResponse) *ManagedServiceIdentityResponse {
		return &v
	}).(ManagedServiceIdentityResponsePtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]ManagedServiceIdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]ManagedServiceIdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ManagedServiceIdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type ManagedServiceIdentityResponseUserAssignedIdentitiesInput interface {
	pulumi.Input

	ToManagedServiceIdentityResponseUserAssignedIdentitiesOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesOutput
	ToManagedServiceIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesOutput
}

type ManagedServiceIdentityResponseUserAssignedIdentitiesArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ManagedServiceIdentityResponseUserAssignedIdentitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i ManagedServiceIdentityResponseUserAssignedIdentitiesArgs) ToManagedServiceIdentityResponseUserAssignedIdentitiesOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesOutput {
	return i.ToManagedServiceIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityResponseUserAssignedIdentitiesArgs) ToManagedServiceIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponseUserAssignedIdentitiesOutput)
}





type ManagedServiceIdentityResponseUserAssignedIdentitiesMapInput interface {
	pulumi.Input

	ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput
	ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput
}

type ManagedServiceIdentityResponseUserAssignedIdentitiesMap map[string]ManagedServiceIdentityResponseUserAssignedIdentitiesInput

func (ManagedServiceIdentityResponseUserAssignedIdentitiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedServiceIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i ManagedServiceIdentityResponseUserAssignedIdentitiesMap) ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return i.ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityResponseUserAssignedIdentitiesMap) ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ManagedServiceIdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesOutput) ToManagedServiceIdentityResponseUserAssignedIdentitiesOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesOutput) ToManagedServiceIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedServiceIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput) ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput) ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) ManagedServiceIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedServiceIdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]ManagedServiceIdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(ManagedServiceIdentityResponseUserAssignedIdentitiesOutput)
}

type MapsAccountProperties struct {
	Cors             *CorsRules       `pulumi:"cors"`
	DisableLocalAuth *bool            `pulumi:"disableLocalAuth"`
	LinkedResources  []LinkedResource `pulumi:"linkedResources"`
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
	Cors             CorsRulesPtrInput        `pulumi:"cors"`
	DisableLocalAuth pulumi.BoolPtrInput      `pulumi:"disableLocalAuth"`
	LinkedResources  LinkedResourceArrayInput `pulumi:"linkedResources"`
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

func (o MapsAccountPropertiesOutput) Cors() CorsRulesPtrOutput {
	return o.ApplyT(func(v MapsAccountProperties) *CorsRules { return v.Cors }).(CorsRulesPtrOutput)
}

func (o MapsAccountPropertiesOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MapsAccountProperties) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o MapsAccountPropertiesOutput) LinkedResources() LinkedResourceArrayOutput {
	return o.ApplyT(func(v MapsAccountProperties) []LinkedResource { return v.LinkedResources }).(LinkedResourceArrayOutput)
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

func (o MapsAccountPropertiesPtrOutput) Cors() CorsRulesPtrOutput {
	return o.ApplyT(func(v *MapsAccountProperties) *CorsRules {
		if v == nil {
			return nil
		}
		return v.Cors
	}).(CorsRulesPtrOutput)
}

func (o MapsAccountPropertiesPtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MapsAccountProperties) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

func (o MapsAccountPropertiesPtrOutput) LinkedResources() LinkedResourceArrayOutput {
	return o.ApplyT(func(v *MapsAccountProperties) []LinkedResource {
		if v == nil {
			return nil
		}
		return v.LinkedResources
	}).(LinkedResourceArrayOutput)
}

type MapsAccountPropertiesResponse struct {
	Cors              *CorsRulesResponse       `pulumi:"cors"`
	DisableLocalAuth  *bool                    `pulumi:"disableLocalAuth"`
	LinkedResources   []LinkedResourceResponse `pulumi:"linkedResources"`
	ProvisioningState string                   `pulumi:"provisioningState"`
	UniqueId          string                   `pulumi:"uniqueId"`
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





type MapsAccountPropertiesResponseInput interface {
	pulumi.Input

	ToMapsAccountPropertiesResponseOutput() MapsAccountPropertiesResponseOutput
	ToMapsAccountPropertiesResponseOutputWithContext(context.Context) MapsAccountPropertiesResponseOutput
}

type MapsAccountPropertiesResponseArgs struct {
	Cors              CorsRulesResponsePtrInput        `pulumi:"cors"`
	DisableLocalAuth  pulumi.BoolPtrInput              `pulumi:"disableLocalAuth"`
	LinkedResources   LinkedResourceResponseArrayInput `pulumi:"linkedResources"`
	ProvisioningState pulumi.StringInput               `pulumi:"provisioningState"`
	UniqueId          pulumi.StringInput               `pulumi:"uniqueId"`
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

func (o MapsAccountPropertiesResponseOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) *CorsRulesResponse { return v.Cors }).(CorsRulesResponsePtrOutput)
}

func (o MapsAccountPropertiesResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o MapsAccountPropertiesResponseOutput) LinkedResources() LinkedResourceResponseArrayOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) []LinkedResourceResponse { return v.LinkedResources }).(LinkedResourceResponseArrayOutput)
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

func (o MapsAccountPropertiesResponsePtrOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v *MapsAccountPropertiesResponse) *CorsRulesResponse {
		if v == nil {
			return nil
		}
		return v.Cors
	}).(CorsRulesResponsePtrOutput)
}

func (o MapsAccountPropertiesResponsePtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MapsAccountPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

func (o MapsAccountPropertiesResponsePtrOutput) LinkedResources() LinkedResourceResponseArrayOutput {
	return o.ApplyT(func(v *MapsAccountPropertiesResponse) []LinkedResourceResponse {
		if v == nil {
			return nil
		}
		return v.LinkedResources
	}).(LinkedResourceResponseArrayOutput)
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
	pulumi.RegisterOutputType(CorsRuleOutput{})
	pulumi.RegisterOutputType(CorsRuleArrayOutput{})
	pulumi.RegisterOutputType(CorsRuleResponseOutput{})
	pulumi.RegisterOutputType(CorsRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(CorsRulesOutput{})
	pulumi.RegisterOutputType(CorsRulesPtrOutput{})
	pulumi.RegisterOutputType(CorsRulesResponseOutput{})
	pulumi.RegisterOutputType(CorsRulesResponsePtrOutput{})
	pulumi.RegisterOutputType(CreatorPropertiesOutput{})
	pulumi.RegisterOutputType(CreatorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreatorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CreatorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LinkedResourceOutput{})
	pulumi.RegisterOutputType(LinkedResourceArrayOutput{})
	pulumi.RegisterOutputType(LinkedResourceResponseOutput{})
	pulumi.RegisterOutputType(LinkedResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput{})
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
