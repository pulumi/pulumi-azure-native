


package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NewNotificationsResponse struct {
	DisplayName          *string                           `pulumi:"displayName"`
	Icon                 *string                           `pulumi:"icon"`
	IsFuturePlansEnabled *bool                             `pulumi:"isFuturePlansEnabled"`
	MessageCode          *float64                          `pulumi:"messageCode"`
	OfferId              *string                           `pulumi:"offerId"`
	Plans                []PlanNotificationDetailsResponse `pulumi:"plans"`
}

type NewNotificationsResponseOutput struct{ *pulumi.OutputState }

func (NewNotificationsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NewNotificationsResponse)(nil)).Elem()
}

func (o NewNotificationsResponseOutput) ToNewNotificationsResponseOutput() NewNotificationsResponseOutput {
	return o
}

func (o NewNotificationsResponseOutput) ToNewNotificationsResponseOutputWithContext(ctx context.Context) NewNotificationsResponseOutput {
	return o
}

func (o NewNotificationsResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NewNotificationsResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o NewNotificationsResponseOutput) Icon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NewNotificationsResponse) *string { return v.Icon }).(pulumi.StringPtrOutput)
}

func (o NewNotificationsResponseOutput) IsFuturePlansEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NewNotificationsResponse) *bool { return v.IsFuturePlansEnabled }).(pulumi.BoolPtrOutput)
}

func (o NewNotificationsResponseOutput) MessageCode() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v NewNotificationsResponse) *float64 { return v.MessageCode }).(pulumi.Float64PtrOutput)
}

func (o NewNotificationsResponseOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NewNotificationsResponse) *string { return v.OfferId }).(pulumi.StringPtrOutput)
}

func (o NewNotificationsResponseOutput) Plans() PlanNotificationDetailsResponseArrayOutput {
	return o.ApplyT(func(v NewNotificationsResponse) []PlanNotificationDetailsResponse { return v.Plans }).(PlanNotificationDetailsResponseArrayOutput)
}

type NewNotificationsResponseArrayOutput struct{ *pulumi.OutputState }

func (NewNotificationsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NewNotificationsResponse)(nil)).Elem()
}

func (o NewNotificationsResponseArrayOutput) ToNewNotificationsResponseArrayOutput() NewNotificationsResponseArrayOutput {
	return o
}

func (o NewNotificationsResponseArrayOutput) ToNewNotificationsResponseArrayOutputWithContext(ctx context.Context) NewNotificationsResponseArrayOutput {
	return o
}

func (o NewNotificationsResponseArrayOutput) Index(i pulumi.IntInput) NewNotificationsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NewNotificationsResponse {
		return vs[0].([]NewNotificationsResponse)[vs[1].(int)]
	}).(NewNotificationsResponseOutput)
}

type Plan struct {
	Accessibility *string `pulumi:"accessibility"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Accessibility pulumi.StringPtrInput `pulumi:"accessibility"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}





type PlanArrayInput interface {
	pulumi.Input

	ToPlanArrayOutput() PlanArrayOutput
	ToPlanArrayOutputWithContext(context.Context) PlanArrayOutput
}

type PlanArray []PlanInput

func (PlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Plan)(nil)).Elem()
}

func (i PlanArray) ToPlanArrayOutput() PlanArrayOutput {
	return i.ToPlanArrayOutputWithContext(context.Background())
}

func (i PlanArray) ToPlanArrayOutputWithContext(ctx context.Context) PlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanArrayOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) Accessibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Accessibility }).(pulumi.StringPtrOutput)
}

type PlanArrayOutput struct{ *pulumi.OutputState }

func (PlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Plan)(nil)).Elem()
}

func (o PlanArrayOutput) ToPlanArrayOutput() PlanArrayOutput {
	return o
}

func (o PlanArrayOutput) ToPlanArrayOutputWithContext(ctx context.Context) PlanArrayOutput {
	return o
}

func (o PlanArrayOutput) Index(i pulumi.IntInput) PlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Plan {
		return vs[0].([]Plan)[vs[1].(int)]
	}).(PlanOutput)
}

type PlanNotificationDetailsResponse struct {
	PlanDisplayName *string `pulumi:"planDisplayName"`
	PlanId          *string `pulumi:"planId"`
}

type PlanNotificationDetailsResponseOutput struct{ *pulumi.OutputState }

func (PlanNotificationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanNotificationDetailsResponse)(nil)).Elem()
}

func (o PlanNotificationDetailsResponseOutput) ToPlanNotificationDetailsResponseOutput() PlanNotificationDetailsResponseOutput {
	return o
}

func (o PlanNotificationDetailsResponseOutput) ToPlanNotificationDetailsResponseOutputWithContext(ctx context.Context) PlanNotificationDetailsResponseOutput {
	return o
}

func (o PlanNotificationDetailsResponseOutput) PlanDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanNotificationDetailsResponse) *string { return v.PlanDisplayName }).(pulumi.StringPtrOutput)
}

func (o PlanNotificationDetailsResponseOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanNotificationDetailsResponse) *string { return v.PlanId }).(pulumi.StringPtrOutput)
}

type PlanNotificationDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (PlanNotificationDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlanNotificationDetailsResponse)(nil)).Elem()
}

func (o PlanNotificationDetailsResponseArrayOutput) ToPlanNotificationDetailsResponseArrayOutput() PlanNotificationDetailsResponseArrayOutput {
	return o
}

func (o PlanNotificationDetailsResponseArrayOutput) ToPlanNotificationDetailsResponseArrayOutputWithContext(ctx context.Context) PlanNotificationDetailsResponseArrayOutput {
	return o
}

func (o PlanNotificationDetailsResponseArrayOutput) Index(i pulumi.IntInput) PlanNotificationDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlanNotificationDetailsResponse {
		return vs[0].([]PlanNotificationDetailsResponse)[vs[1].(int)]
	}).(PlanNotificationDetailsResponseOutput)
}

type PlanResponse struct {
	Accessibility     *string `pulumi:"accessibility"`
	AltStackReference string  `pulumi:"altStackReference"`
	PlanDisplayName   string  `pulumi:"planDisplayName"`
	PlanId            string  `pulumi:"planId"`
	SkuId             string  `pulumi:"skuId"`
	StackType         string  `pulumi:"stackType"`
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) Accessibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Accessibility }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) AltStackReference() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.AltStackReference }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) PlanDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.PlanDisplayName }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.PlanId }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) SkuId() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.SkuId }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) StackType() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.StackType }).(pulumi.StringOutput)
}

type PlanResponseArrayOutput struct{ *pulumi.OutputState }

func (PlanResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlanResponse)(nil)).Elem()
}

func (o PlanResponseArrayOutput) ToPlanResponseArrayOutput() PlanResponseArrayOutput {
	return o
}

func (o PlanResponseArrayOutput) ToPlanResponseArrayOutputWithContext(ctx context.Context) PlanResponseArrayOutput {
	return o
}

func (o PlanResponseArrayOutput) Index(i pulumi.IntInput) PlanResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlanResponse {
		return vs[0].([]PlanResponse)[vs[1].(int)]
	}).(PlanResponseOutput)
}

type RuleResponse struct {
	Type  *string  `pulumi:"type"`
	Value []string `pulumi:"value"`
}

type RuleResponseOutput struct{ *pulumi.OutputState }

func (RuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleResponse)(nil)).Elem()
}

func (o RuleResponseOutput) ToRuleResponseOutput() RuleResponseOutput {
	return o
}

func (o RuleResponseOutput) ToRuleResponseOutputWithContext(ctx context.Context) RuleResponseOutput {
	return o
}

func (o RuleResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o RuleResponseOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuleResponse) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type RuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleResponse)(nil)).Elem()
}

func (o RuleResponseArrayOutput) ToRuleResponseArrayOutput() RuleResponseArrayOutput {
	return o
}

func (o RuleResponseArrayOutput) ToRuleResponseArrayOutputWithContext(ctx context.Context) RuleResponseArrayOutput {
	return o
}

func (o RuleResponseArrayOutput) Index(i pulumi.IntInput) RuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleResponse {
		return vs[0].([]RuleResponse)[vs[1].(int)]
	}).(RuleResponseOutput)
}

type StopSellOffersPlansNotificationsListPropertiesResponse struct {
	DisplayName      string                            `pulumi:"displayName"`
	Icon             string                            `pulumi:"icon"`
	IsEntire         bool                              `pulumi:"isEntire"`
	MessageCode      float64                           `pulumi:"messageCode"`
	OfferId          string                            `pulumi:"offerId"`
	Plans            []PlanNotificationDetailsResponse `pulumi:"plans"`
	PublicContext    bool                              `pulumi:"publicContext"`
	SubscriptionsIds []string                          `pulumi:"subscriptionsIds"`
}

type StopSellOffersPlansNotificationsListPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StopSellOffersPlansNotificationsListPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StopSellOffersPlansNotificationsListPropertiesResponse)(nil)).Elem()
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) ToStopSellOffersPlansNotificationsListPropertiesResponseOutput() StopSellOffersPlansNotificationsListPropertiesResponseOutput {
	return o
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) ToStopSellOffersPlansNotificationsListPropertiesResponseOutputWithContext(ctx context.Context) StopSellOffersPlansNotificationsListPropertiesResponseOutput {
	return o
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) Icon() pulumi.StringOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) string { return v.Icon }).(pulumi.StringOutput)
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) IsEntire() pulumi.BoolOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) bool { return v.IsEntire }).(pulumi.BoolOutput)
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) MessageCode() pulumi.Float64Output {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) float64 { return v.MessageCode }).(pulumi.Float64Output)
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) OfferId() pulumi.StringOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) string { return v.OfferId }).(pulumi.StringOutput)
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) Plans() PlanNotificationDetailsResponseArrayOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) []PlanNotificationDetailsResponse {
		return v.Plans
	}).(PlanNotificationDetailsResponseArrayOutput)
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) PublicContext() pulumi.BoolOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) bool { return v.PublicContext }).(pulumi.BoolOutput)
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseOutput) SubscriptionsIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StopSellOffersPlansNotificationsListPropertiesResponse) []string { return v.SubscriptionsIds }).(pulumi.StringArrayOutput)
}

type StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StopSellOffersPlansNotificationsListPropertiesResponse)(nil)).Elem()
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput) ToStopSellOffersPlansNotificationsListPropertiesResponseArrayOutput() StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput {
	return o
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput) ToStopSellOffersPlansNotificationsListPropertiesResponseArrayOutputWithContext(ctx context.Context) StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput {
	return o
}

func (o StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput) Index(i pulumi.IntInput) StopSellOffersPlansNotificationsListPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StopSellOffersPlansNotificationsListPropertiesResponse {
		return vs[0].([]StopSellOffersPlansNotificationsListPropertiesResponse)[vs[1].(int)]
	}).(StopSellOffersPlansNotificationsListPropertiesResponseOutput)
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
	pulumi.RegisterOutputType(NewNotificationsResponseOutput{})
	pulumi.RegisterOutputType(NewNotificationsResponseArrayOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanArrayOutput{})
	pulumi.RegisterOutputType(PlanNotificationDetailsResponseOutput{})
	pulumi.RegisterOutputType(PlanNotificationDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponseArrayOutput{})
	pulumi.RegisterOutputType(RuleResponseOutput{})
	pulumi.RegisterOutputType(RuleResponseArrayOutput{})
	pulumi.RegisterOutputType(StopSellOffersPlansNotificationsListPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StopSellOffersPlansNotificationsListPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
