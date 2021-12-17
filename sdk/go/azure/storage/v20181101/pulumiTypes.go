


package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CorsRule struct {
	AllowedHeaders  []string `pulumi:"allowedHeaders"`
	AllowedMethods  []string `pulumi:"allowedMethods"`
	AllowedOrigins  []string `pulumi:"allowedOrigins"`
	ExposedHeaders  []string `pulumi:"exposedHeaders"`
	MaxAgeInSeconds int      `pulumi:"maxAgeInSeconds"`
}





type CorsRuleInput interface {
	pulumi.Input

	ToCorsRuleOutput() CorsRuleOutput
	ToCorsRuleOutputWithContext(context.Context) CorsRuleOutput
}

type CorsRuleArgs struct {
	AllowedHeaders  pulumi.StringArrayInput `pulumi:"allowedHeaders"`
	AllowedMethods  pulumi.StringArrayInput `pulumi:"allowedMethods"`
	AllowedOrigins  pulumi.StringArrayInput `pulumi:"allowedOrigins"`
	ExposedHeaders  pulumi.StringArrayInput `pulumi:"exposedHeaders"`
	MaxAgeInSeconds pulumi.IntInput         `pulumi:"maxAgeInSeconds"`
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

func (o CorsRuleOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) ExposedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.ExposedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) MaxAgeInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v CorsRule) int { return v.MaxAgeInSeconds }).(pulumi.IntOutput)
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
	AllowedHeaders  []string `pulumi:"allowedHeaders"`
	AllowedMethods  []string `pulumi:"allowedMethods"`
	AllowedOrigins  []string `pulumi:"allowedOrigins"`
	ExposedHeaders  []string `pulumi:"exposedHeaders"`
	MaxAgeInSeconds int      `pulumi:"maxAgeInSeconds"`
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

func (o CorsRuleResponseOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) ExposedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.ExposedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) MaxAgeInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v CorsRuleResponse) int { return v.MaxAgeInSeconds }).(pulumi.IntOutput)
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

type CustomDomain struct {
	Name             string `pulumi:"name"`
	UseSubDomainName *bool  `pulumi:"useSubDomainName"`
}





type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(context.Context) CustomDomainOutput
}

type CustomDomainArgs struct {
	Name             pulumi.StringInput  `pulumi:"name"`
	UseSubDomainName pulumi.BoolPtrInput `pulumi:"useSubDomainName"`
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (i CustomDomainArgs) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

func (i CustomDomainArgs) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput).ToCustomDomainPtrOutputWithContext(ctx)
}









type CustomDomainPtrInput interface {
	pulumi.Input

	ToCustomDomainPtrOutput() CustomDomainPtrOutput
	ToCustomDomainPtrOutputWithContext(context.Context) CustomDomainPtrOutput
}

type customDomainPtrType CustomDomainArgs

func CustomDomainPtr(v *CustomDomainArgs) CustomDomainPtrInput {
	return (*customDomainPtrType)(v)
}

func (*customDomainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (i *customDomainPtrType) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i *customDomainPtrType) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPtrOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return o.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (o CustomDomainOutput) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomain) *CustomDomain {
		return &v
	}).(CustomDomainPtrOutput)
}

func (o CustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomain) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomain) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type CustomDomainPtrOutput struct{ *pulumi.OutputState }

func (CustomDomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (o CustomDomainPtrOutput) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return o
}

func (o CustomDomainPtrOutput) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return o
}

func (o CustomDomainPtrOutput) Elem() CustomDomainOutput {
	return o.ApplyT(func(v *CustomDomain) CustomDomain {
		if v != nil {
			return *v
		}
		var ret CustomDomain
		return ret
	}).(CustomDomainOutput)
}

func (o CustomDomainPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomain) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainPtrOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomain) *bool {
		if v == nil {
			return nil
		}
		return v.UseSubDomainName
	}).(pulumi.BoolPtrOutput)
}

type CustomDomainResponse struct {
	Name             string `pulumi:"name"`
	UseSubDomainName *bool  `pulumi:"useSubDomainName"`
}

type CustomDomainResponseOutput struct{ *pulumi.OutputState }

func (CustomDomainResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutput() CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutputWithContext(ctx context.Context) CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomainResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainResponseOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type DateAfterCreation struct {
	DaysAfterCreationGreaterThan int `pulumi:"daysAfterCreationGreaterThan"`
}





type DateAfterCreationInput interface {
	pulumi.Input

	ToDateAfterCreationOutput() DateAfterCreationOutput
	ToDateAfterCreationOutputWithContext(context.Context) DateAfterCreationOutput
}

type DateAfterCreationArgs struct {
	DaysAfterCreationGreaterThan pulumi.IntInput `pulumi:"daysAfterCreationGreaterThan"`
}

func (DateAfterCreationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterCreation)(nil)).Elem()
}

func (i DateAfterCreationArgs) ToDateAfterCreationOutput() DateAfterCreationOutput {
	return i.ToDateAfterCreationOutputWithContext(context.Background())
}

func (i DateAfterCreationArgs) ToDateAfterCreationOutputWithContext(ctx context.Context) DateAfterCreationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterCreationOutput)
}

func (i DateAfterCreationArgs) ToDateAfterCreationPtrOutput() DateAfterCreationPtrOutput {
	return i.ToDateAfterCreationPtrOutputWithContext(context.Background())
}

func (i DateAfterCreationArgs) ToDateAfterCreationPtrOutputWithContext(ctx context.Context) DateAfterCreationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterCreationOutput).ToDateAfterCreationPtrOutputWithContext(ctx)
}









type DateAfterCreationPtrInput interface {
	pulumi.Input

	ToDateAfterCreationPtrOutput() DateAfterCreationPtrOutput
	ToDateAfterCreationPtrOutputWithContext(context.Context) DateAfterCreationPtrOutput
}

type dateAfterCreationPtrType DateAfterCreationArgs

func DateAfterCreationPtr(v *DateAfterCreationArgs) DateAfterCreationPtrInput {
	return (*dateAfterCreationPtrType)(v)
}

func (*dateAfterCreationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterCreation)(nil)).Elem()
}

func (i *dateAfterCreationPtrType) ToDateAfterCreationPtrOutput() DateAfterCreationPtrOutput {
	return i.ToDateAfterCreationPtrOutputWithContext(context.Background())
}

func (i *dateAfterCreationPtrType) ToDateAfterCreationPtrOutputWithContext(ctx context.Context) DateAfterCreationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterCreationPtrOutput)
}

type DateAfterCreationOutput struct{ *pulumi.OutputState }

func (DateAfterCreationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterCreation)(nil)).Elem()
}

func (o DateAfterCreationOutput) ToDateAfterCreationOutput() DateAfterCreationOutput {
	return o
}

func (o DateAfterCreationOutput) ToDateAfterCreationOutputWithContext(ctx context.Context) DateAfterCreationOutput {
	return o
}

func (o DateAfterCreationOutput) ToDateAfterCreationPtrOutput() DateAfterCreationPtrOutput {
	return o.ToDateAfterCreationPtrOutputWithContext(context.Background())
}

func (o DateAfterCreationOutput) ToDateAfterCreationPtrOutputWithContext(ctx context.Context) DateAfterCreationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DateAfterCreation) *DateAfterCreation {
		return &v
	}).(DateAfterCreationPtrOutput)
}

func (o DateAfterCreationOutput) DaysAfterCreationGreaterThan() pulumi.IntOutput {
	return o.ApplyT(func(v DateAfterCreation) int { return v.DaysAfterCreationGreaterThan }).(pulumi.IntOutput)
}

type DateAfterCreationPtrOutput struct{ *pulumi.OutputState }

func (DateAfterCreationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterCreation)(nil)).Elem()
}

func (o DateAfterCreationPtrOutput) ToDateAfterCreationPtrOutput() DateAfterCreationPtrOutput {
	return o
}

func (o DateAfterCreationPtrOutput) ToDateAfterCreationPtrOutputWithContext(ctx context.Context) DateAfterCreationPtrOutput {
	return o
}

func (o DateAfterCreationPtrOutput) Elem() DateAfterCreationOutput {
	return o.ApplyT(func(v *DateAfterCreation) DateAfterCreation {
		if v != nil {
			return *v
		}
		var ret DateAfterCreation
		return ret
	}).(DateAfterCreationOutput)
}

func (o DateAfterCreationPtrOutput) DaysAfterCreationGreaterThan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DateAfterCreation) *int {
		if v == nil {
			return nil
		}
		return &v.DaysAfterCreationGreaterThan
	}).(pulumi.IntPtrOutput)
}

type DateAfterCreationResponse struct {
	DaysAfterCreationGreaterThan int `pulumi:"daysAfterCreationGreaterThan"`
}

type DateAfterCreationResponseOutput struct{ *pulumi.OutputState }

func (DateAfterCreationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterCreationResponse)(nil)).Elem()
}

func (o DateAfterCreationResponseOutput) ToDateAfterCreationResponseOutput() DateAfterCreationResponseOutput {
	return o
}

func (o DateAfterCreationResponseOutput) ToDateAfterCreationResponseOutputWithContext(ctx context.Context) DateAfterCreationResponseOutput {
	return o
}

func (o DateAfterCreationResponseOutput) DaysAfterCreationGreaterThan() pulumi.IntOutput {
	return o.ApplyT(func(v DateAfterCreationResponse) int { return v.DaysAfterCreationGreaterThan }).(pulumi.IntOutput)
}

type DateAfterCreationResponsePtrOutput struct{ *pulumi.OutputState }

func (DateAfterCreationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterCreationResponse)(nil)).Elem()
}

func (o DateAfterCreationResponsePtrOutput) ToDateAfterCreationResponsePtrOutput() DateAfterCreationResponsePtrOutput {
	return o
}

func (o DateAfterCreationResponsePtrOutput) ToDateAfterCreationResponsePtrOutputWithContext(ctx context.Context) DateAfterCreationResponsePtrOutput {
	return o
}

func (o DateAfterCreationResponsePtrOutput) Elem() DateAfterCreationResponseOutput {
	return o.ApplyT(func(v *DateAfterCreationResponse) DateAfterCreationResponse {
		if v != nil {
			return *v
		}
		var ret DateAfterCreationResponse
		return ret
	}).(DateAfterCreationResponseOutput)
}

func (o DateAfterCreationResponsePtrOutput) DaysAfterCreationGreaterThan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DateAfterCreationResponse) *int {
		if v == nil {
			return nil
		}
		return &v.DaysAfterCreationGreaterThan
	}).(pulumi.IntPtrOutput)
}

type DateAfterModification struct {
	DaysAfterModificationGreaterThan int `pulumi:"daysAfterModificationGreaterThan"`
}





type DateAfterModificationInput interface {
	pulumi.Input

	ToDateAfterModificationOutput() DateAfterModificationOutput
	ToDateAfterModificationOutputWithContext(context.Context) DateAfterModificationOutput
}

type DateAfterModificationArgs struct {
	DaysAfterModificationGreaterThan pulumi.IntInput `pulumi:"daysAfterModificationGreaterThan"`
}

func (DateAfterModificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterModification)(nil)).Elem()
}

func (i DateAfterModificationArgs) ToDateAfterModificationOutput() DateAfterModificationOutput {
	return i.ToDateAfterModificationOutputWithContext(context.Background())
}

func (i DateAfterModificationArgs) ToDateAfterModificationOutputWithContext(ctx context.Context) DateAfterModificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterModificationOutput)
}

func (i DateAfterModificationArgs) ToDateAfterModificationPtrOutput() DateAfterModificationPtrOutput {
	return i.ToDateAfterModificationPtrOutputWithContext(context.Background())
}

func (i DateAfterModificationArgs) ToDateAfterModificationPtrOutputWithContext(ctx context.Context) DateAfterModificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterModificationOutput).ToDateAfterModificationPtrOutputWithContext(ctx)
}









type DateAfterModificationPtrInput interface {
	pulumi.Input

	ToDateAfterModificationPtrOutput() DateAfterModificationPtrOutput
	ToDateAfterModificationPtrOutputWithContext(context.Context) DateAfterModificationPtrOutput
}

type dateAfterModificationPtrType DateAfterModificationArgs

func DateAfterModificationPtr(v *DateAfterModificationArgs) DateAfterModificationPtrInput {
	return (*dateAfterModificationPtrType)(v)
}

func (*dateAfterModificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterModification)(nil)).Elem()
}

func (i *dateAfterModificationPtrType) ToDateAfterModificationPtrOutput() DateAfterModificationPtrOutput {
	return i.ToDateAfterModificationPtrOutputWithContext(context.Background())
}

func (i *dateAfterModificationPtrType) ToDateAfterModificationPtrOutputWithContext(ctx context.Context) DateAfterModificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterModificationPtrOutput)
}

type DateAfterModificationOutput struct{ *pulumi.OutputState }

func (DateAfterModificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterModification)(nil)).Elem()
}

func (o DateAfterModificationOutput) ToDateAfterModificationOutput() DateAfterModificationOutput {
	return o
}

func (o DateAfterModificationOutput) ToDateAfterModificationOutputWithContext(ctx context.Context) DateAfterModificationOutput {
	return o
}

func (o DateAfterModificationOutput) ToDateAfterModificationPtrOutput() DateAfterModificationPtrOutput {
	return o.ToDateAfterModificationPtrOutputWithContext(context.Background())
}

func (o DateAfterModificationOutput) ToDateAfterModificationPtrOutputWithContext(ctx context.Context) DateAfterModificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DateAfterModification) *DateAfterModification {
		return &v
	}).(DateAfterModificationPtrOutput)
}

func (o DateAfterModificationOutput) DaysAfterModificationGreaterThan() pulumi.IntOutput {
	return o.ApplyT(func(v DateAfterModification) int { return v.DaysAfterModificationGreaterThan }).(pulumi.IntOutput)
}

type DateAfterModificationPtrOutput struct{ *pulumi.OutputState }

func (DateAfterModificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterModification)(nil)).Elem()
}

func (o DateAfterModificationPtrOutput) ToDateAfterModificationPtrOutput() DateAfterModificationPtrOutput {
	return o
}

func (o DateAfterModificationPtrOutput) ToDateAfterModificationPtrOutputWithContext(ctx context.Context) DateAfterModificationPtrOutput {
	return o
}

func (o DateAfterModificationPtrOutput) Elem() DateAfterModificationOutput {
	return o.ApplyT(func(v *DateAfterModification) DateAfterModification {
		if v != nil {
			return *v
		}
		var ret DateAfterModification
		return ret
	}).(DateAfterModificationOutput)
}

func (o DateAfterModificationPtrOutput) DaysAfterModificationGreaterThan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DateAfterModification) *int {
		if v == nil {
			return nil
		}
		return &v.DaysAfterModificationGreaterThan
	}).(pulumi.IntPtrOutput)
}

type DateAfterModificationResponse struct {
	DaysAfterModificationGreaterThan int `pulumi:"daysAfterModificationGreaterThan"`
}

type DateAfterModificationResponseOutput struct{ *pulumi.OutputState }

func (DateAfterModificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterModificationResponse)(nil)).Elem()
}

func (o DateAfterModificationResponseOutput) ToDateAfterModificationResponseOutput() DateAfterModificationResponseOutput {
	return o
}

func (o DateAfterModificationResponseOutput) ToDateAfterModificationResponseOutputWithContext(ctx context.Context) DateAfterModificationResponseOutput {
	return o
}

func (o DateAfterModificationResponseOutput) DaysAfterModificationGreaterThan() pulumi.IntOutput {
	return o.ApplyT(func(v DateAfterModificationResponse) int { return v.DaysAfterModificationGreaterThan }).(pulumi.IntOutput)
}

type DateAfterModificationResponsePtrOutput struct{ *pulumi.OutputState }

func (DateAfterModificationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterModificationResponse)(nil)).Elem()
}

func (o DateAfterModificationResponsePtrOutput) ToDateAfterModificationResponsePtrOutput() DateAfterModificationResponsePtrOutput {
	return o
}

func (o DateAfterModificationResponsePtrOutput) ToDateAfterModificationResponsePtrOutputWithContext(ctx context.Context) DateAfterModificationResponsePtrOutput {
	return o
}

func (o DateAfterModificationResponsePtrOutput) Elem() DateAfterModificationResponseOutput {
	return o.ApplyT(func(v *DateAfterModificationResponse) DateAfterModificationResponse {
		if v != nil {
			return *v
		}
		var ret DateAfterModificationResponse
		return ret
	}).(DateAfterModificationResponseOutput)
}

func (o DateAfterModificationResponsePtrOutput) DaysAfterModificationGreaterThan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DateAfterModificationResponse) *int {
		if v == nil {
			return nil
		}
		return &v.DaysAfterModificationGreaterThan
	}).(pulumi.IntPtrOutput)
}

type DeleteRetentionPolicy struct {
	Days    *int  `pulumi:"days"`
	Enabled *bool `pulumi:"enabled"`
}





type DeleteRetentionPolicyInput interface {
	pulumi.Input

	ToDeleteRetentionPolicyOutput() DeleteRetentionPolicyOutput
	ToDeleteRetentionPolicyOutputWithContext(context.Context) DeleteRetentionPolicyOutput
}

type DeleteRetentionPolicyArgs struct {
	Days    pulumi.IntPtrInput  `pulumi:"days"`
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (DeleteRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteRetentionPolicy)(nil)).Elem()
}

func (i DeleteRetentionPolicyArgs) ToDeleteRetentionPolicyOutput() DeleteRetentionPolicyOutput {
	return i.ToDeleteRetentionPolicyOutputWithContext(context.Background())
}

func (i DeleteRetentionPolicyArgs) ToDeleteRetentionPolicyOutputWithContext(ctx context.Context) DeleteRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeleteRetentionPolicyOutput)
}

func (i DeleteRetentionPolicyArgs) ToDeleteRetentionPolicyPtrOutput() DeleteRetentionPolicyPtrOutput {
	return i.ToDeleteRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i DeleteRetentionPolicyArgs) ToDeleteRetentionPolicyPtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeleteRetentionPolicyOutput).ToDeleteRetentionPolicyPtrOutputWithContext(ctx)
}









type DeleteRetentionPolicyPtrInput interface {
	pulumi.Input

	ToDeleteRetentionPolicyPtrOutput() DeleteRetentionPolicyPtrOutput
	ToDeleteRetentionPolicyPtrOutputWithContext(context.Context) DeleteRetentionPolicyPtrOutput
}

type deleteRetentionPolicyPtrType DeleteRetentionPolicyArgs

func DeleteRetentionPolicyPtr(v *DeleteRetentionPolicyArgs) DeleteRetentionPolicyPtrInput {
	return (*deleteRetentionPolicyPtrType)(v)
}

func (*deleteRetentionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteRetentionPolicy)(nil)).Elem()
}

func (i *deleteRetentionPolicyPtrType) ToDeleteRetentionPolicyPtrOutput() DeleteRetentionPolicyPtrOutput {
	return i.ToDeleteRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i *deleteRetentionPolicyPtrType) ToDeleteRetentionPolicyPtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeleteRetentionPolicyPtrOutput)
}

type DeleteRetentionPolicyOutput struct{ *pulumi.OutputState }

func (DeleteRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteRetentionPolicy)(nil)).Elem()
}

func (o DeleteRetentionPolicyOutput) ToDeleteRetentionPolicyOutput() DeleteRetentionPolicyOutput {
	return o
}

func (o DeleteRetentionPolicyOutput) ToDeleteRetentionPolicyOutputWithContext(ctx context.Context) DeleteRetentionPolicyOutput {
	return o
}

func (o DeleteRetentionPolicyOutput) ToDeleteRetentionPolicyPtrOutput() DeleteRetentionPolicyPtrOutput {
	return o.ToDeleteRetentionPolicyPtrOutputWithContext(context.Background())
}

func (o DeleteRetentionPolicyOutput) ToDeleteRetentionPolicyPtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeleteRetentionPolicy) *DeleteRetentionPolicy {
		return &v
	}).(DeleteRetentionPolicyPtrOutput)
}

func (o DeleteRetentionPolicyOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeleteRetentionPolicy) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o DeleteRetentionPolicyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DeleteRetentionPolicy) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type DeleteRetentionPolicyPtrOutput struct{ *pulumi.OutputState }

func (DeleteRetentionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteRetentionPolicy)(nil)).Elem()
}

func (o DeleteRetentionPolicyPtrOutput) ToDeleteRetentionPolicyPtrOutput() DeleteRetentionPolicyPtrOutput {
	return o
}

func (o DeleteRetentionPolicyPtrOutput) ToDeleteRetentionPolicyPtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyPtrOutput {
	return o
}

func (o DeleteRetentionPolicyPtrOutput) Elem() DeleteRetentionPolicyOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicy) DeleteRetentionPolicy {
		if v != nil {
			return *v
		}
		var ret DeleteRetentionPolicy
		return ret
	}).(DeleteRetentionPolicyOutput)
}

func (o DeleteRetentionPolicyPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicy) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o DeleteRetentionPolicyPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type DeleteRetentionPolicyResponse struct {
	Days    *int  `pulumi:"days"`
	Enabled *bool `pulumi:"enabled"`
}

type DeleteRetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (DeleteRetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteRetentionPolicyResponse)(nil)).Elem()
}

func (o DeleteRetentionPolicyResponseOutput) ToDeleteRetentionPolicyResponseOutput() DeleteRetentionPolicyResponseOutput {
	return o
}

func (o DeleteRetentionPolicyResponseOutput) ToDeleteRetentionPolicyResponseOutputWithContext(ctx context.Context) DeleteRetentionPolicyResponseOutput {
	return o
}

func (o DeleteRetentionPolicyResponseOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeleteRetentionPolicyResponse) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o DeleteRetentionPolicyResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DeleteRetentionPolicyResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type DeleteRetentionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (DeleteRetentionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteRetentionPolicyResponse)(nil)).Elem()
}

func (o DeleteRetentionPolicyResponsePtrOutput) ToDeleteRetentionPolicyResponsePtrOutput() DeleteRetentionPolicyResponsePtrOutput {
	return o
}

func (o DeleteRetentionPolicyResponsePtrOutput) ToDeleteRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyResponsePtrOutput {
	return o
}

func (o DeleteRetentionPolicyResponsePtrOutput) Elem() DeleteRetentionPolicyResponseOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicyResponse) DeleteRetentionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret DeleteRetentionPolicyResponse
		return ret
	}).(DeleteRetentionPolicyResponseOutput)
}

func (o DeleteRetentionPolicyResponsePtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o DeleteRetentionPolicyResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type Encryption struct {
	KeySource          string              `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	Services           *EncryptionServices `pulumi:"services"`
}


func (val *Encryption) Defaults() *Encryption {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		tmp.KeySource = "Microsoft.Storage"
	}
	return &tmp
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeySource          pulumi.StringInput         `pulumi:"keySource"`
	KeyVaultProperties KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
	Services           EncryptionServicesPtrInput `pulumi:"services"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v Encryption) string { return v.KeySource }).(pulumi.StringOutput)
}

func (o EncryptionOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v Encryption) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

func (o EncryptionOutput) Services() EncryptionServicesPtrOutput {
	return o.ApplyT(func(v Encryption) *EncryptionServices { return v.Services }).(EncryptionServicesPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return &v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *Encryption) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

func (o EncryptionPtrOutput) Services() EncryptionServicesPtrOutput {
	return o.ApplyT(func(v *Encryption) *EncryptionServices {
		if v == nil {
			return nil
		}
		return v.Services
	}).(EncryptionServicesPtrOutput)
}

type EncryptionResponse struct {
	KeySource          string                      `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Services           *EncryptionServicesResponse `pulumi:"services"`
}


func (val *EncryptionResponse) Defaults() *EncryptionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		tmp.KeySource = "Microsoft.Storage"
	}
	return &tmp
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionResponse) string { return v.KeySource }).(pulumi.StringOutput)
}

func (o EncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionResponseOutput) Services() EncryptionServicesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *EncryptionServicesResponse { return v.Services }).(EncryptionServicesResponsePtrOutput)
}

type EncryptionService struct {
	Enabled *bool `pulumi:"enabled"`
}





type EncryptionServiceInput interface {
	pulumi.Input

	ToEncryptionServiceOutput() EncryptionServiceOutput
	ToEncryptionServiceOutputWithContext(context.Context) EncryptionServiceOutput
}

type EncryptionServiceArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (EncryptionServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionService)(nil)).Elem()
}

func (i EncryptionServiceArgs) ToEncryptionServiceOutput() EncryptionServiceOutput {
	return i.ToEncryptionServiceOutputWithContext(context.Background())
}

func (i EncryptionServiceArgs) ToEncryptionServiceOutputWithContext(ctx context.Context) EncryptionServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceOutput)
}

func (i EncryptionServiceArgs) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return i.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (i EncryptionServiceArgs) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceOutput).ToEncryptionServicePtrOutputWithContext(ctx)
}









type EncryptionServicePtrInput interface {
	pulumi.Input

	ToEncryptionServicePtrOutput() EncryptionServicePtrOutput
	ToEncryptionServicePtrOutputWithContext(context.Context) EncryptionServicePtrOutput
}

type encryptionServicePtrType EncryptionServiceArgs

func EncryptionServicePtr(v *EncryptionServiceArgs) EncryptionServicePtrInput {
	return (*encryptionServicePtrType)(v)
}

func (*encryptionServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionService)(nil)).Elem()
}

func (i *encryptionServicePtrType) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return i.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (i *encryptionServicePtrType) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicePtrOutput)
}

type EncryptionServiceOutput struct{ *pulumi.OutputState }

func (EncryptionServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionService)(nil)).Elem()
}

func (o EncryptionServiceOutput) ToEncryptionServiceOutput() EncryptionServiceOutput {
	return o
}

func (o EncryptionServiceOutput) ToEncryptionServiceOutputWithContext(ctx context.Context) EncryptionServiceOutput {
	return o
}

func (o EncryptionServiceOutput) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return o.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (o EncryptionServiceOutput) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionService) *EncryptionService {
		return &v
	}).(EncryptionServicePtrOutput)
}

func (o EncryptionServiceOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionService) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type EncryptionServicePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionService)(nil)).Elem()
}

func (o EncryptionServicePtrOutput) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return o
}

func (o EncryptionServicePtrOutput) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return o
}

func (o EncryptionServicePtrOutput) Elem() EncryptionServiceOutput {
	return o.ApplyT(func(v *EncryptionService) EncryptionService {
		if v != nil {
			return *v
		}
		var ret EncryptionService
		return ret
	}).(EncryptionServiceOutput)
}

func (o EncryptionServicePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionService) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type EncryptionServiceResponse struct {
	Enabled         *bool  `pulumi:"enabled"`
	LastEnabledTime string `pulumi:"lastEnabledTime"`
}

type EncryptionServiceResponseOutput struct{ *pulumi.OutputState }

func (EncryptionServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServiceResponse)(nil)).Elem()
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponseOutput() EncryptionServiceResponseOutput {
	return o
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponseOutputWithContext(ctx context.Context) EncryptionServiceResponseOutput {
	return o
}

func (o EncryptionServiceResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o EncryptionServiceResponseOutput) LastEnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) string { return v.LastEnabledTime }).(pulumi.StringOutput)
}

type EncryptionServiceResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServiceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServiceResponse)(nil)).Elem()
}

func (o EncryptionServiceResponsePtrOutput) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return o
}

func (o EncryptionServiceResponsePtrOutput) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return o
}

func (o EncryptionServiceResponsePtrOutput) Elem() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) EncryptionServiceResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionServiceResponse
		return ret
	}).(EncryptionServiceResponseOutput)
}

func (o EncryptionServiceResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionServiceResponsePtrOutput) LastEnabledTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastEnabledTime
	}).(pulumi.StringPtrOutput)
}

type EncryptionServices struct {
	Blob *EncryptionService `pulumi:"blob"`
	File *EncryptionService `pulumi:"file"`
}





type EncryptionServicesInput interface {
	pulumi.Input

	ToEncryptionServicesOutput() EncryptionServicesOutput
	ToEncryptionServicesOutputWithContext(context.Context) EncryptionServicesOutput
}

type EncryptionServicesArgs struct {
	Blob EncryptionServicePtrInput `pulumi:"blob"`
	File EncryptionServicePtrInput `pulumi:"file"`
}

func (EncryptionServicesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServices)(nil)).Elem()
}

func (i EncryptionServicesArgs) ToEncryptionServicesOutput() EncryptionServicesOutput {
	return i.ToEncryptionServicesOutputWithContext(context.Background())
}

func (i EncryptionServicesArgs) ToEncryptionServicesOutputWithContext(ctx context.Context) EncryptionServicesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesOutput)
}

func (i EncryptionServicesArgs) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return i.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (i EncryptionServicesArgs) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesOutput).ToEncryptionServicesPtrOutputWithContext(ctx)
}









type EncryptionServicesPtrInput interface {
	pulumi.Input

	ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput
	ToEncryptionServicesPtrOutputWithContext(context.Context) EncryptionServicesPtrOutput
}

type encryptionServicesPtrType EncryptionServicesArgs

func EncryptionServicesPtr(v *EncryptionServicesArgs) EncryptionServicesPtrInput {
	return (*encryptionServicesPtrType)(v)
}

func (*encryptionServicesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServices)(nil)).Elem()
}

func (i *encryptionServicesPtrType) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return i.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (i *encryptionServicesPtrType) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesPtrOutput)
}

type EncryptionServicesOutput struct{ *pulumi.OutputState }

func (EncryptionServicesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServices)(nil)).Elem()
}

func (o EncryptionServicesOutput) ToEncryptionServicesOutput() EncryptionServicesOutput {
	return o
}

func (o EncryptionServicesOutput) ToEncryptionServicesOutputWithContext(ctx context.Context) EncryptionServicesOutput {
	return o
}

func (o EncryptionServicesOutput) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return o.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (o EncryptionServicesOutput) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionServices) *EncryptionServices {
		return &v
	}).(EncryptionServicesPtrOutput)
}

func (o EncryptionServicesOutput) Blob() EncryptionServicePtrOutput {
	return o.ApplyT(func(v EncryptionServices) *EncryptionService { return v.Blob }).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesOutput) File() EncryptionServicePtrOutput {
	return o.ApplyT(func(v EncryptionServices) *EncryptionService { return v.File }).(EncryptionServicePtrOutput)
}

type EncryptionServicesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServices)(nil)).Elem()
}

func (o EncryptionServicesPtrOutput) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return o
}

func (o EncryptionServicesPtrOutput) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return o
}

func (o EncryptionServicesPtrOutput) Elem() EncryptionServicesOutput {
	return o.ApplyT(func(v *EncryptionServices) EncryptionServices {
		if v != nil {
			return *v
		}
		var ret EncryptionServices
		return ret
	}).(EncryptionServicesOutput)
}

func (o EncryptionServicesPtrOutput) Blob() EncryptionServicePtrOutput {
	return o.ApplyT(func(v *EncryptionServices) *EncryptionService {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesPtrOutput) File() EncryptionServicePtrOutput {
	return o.ApplyT(func(v *EncryptionServices) *EncryptionService {
		if v == nil {
			return nil
		}
		return v.File
	}).(EncryptionServicePtrOutput)
}

type EncryptionServicesResponse struct {
	Blob  *EncryptionServiceResponse `pulumi:"blob"`
	File  *EncryptionServiceResponse `pulumi:"file"`
	Queue EncryptionServiceResponse  `pulumi:"queue"`
	Table EncryptionServiceResponse  `pulumi:"table"`
}

type EncryptionServicesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionServicesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServicesResponse)(nil)).Elem()
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponseOutput() EncryptionServicesResponseOutput {
	return o
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponseOutputWithContext(ctx context.Context) EncryptionServicesResponseOutput {
	return o
}

func (o EncryptionServicesResponseOutput) Blob() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.Blob }).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) File() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.File }).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) Queue() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) EncryptionServiceResponse { return v.Queue }).(EncryptionServiceResponseOutput)
}

func (o EncryptionServicesResponseOutput) Table() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) EncryptionServiceResponse { return v.Table }).(EncryptionServiceResponseOutput)
}

type EncryptionServicesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServicesResponse)(nil)).Elem()
}

func (o EncryptionServicesResponsePtrOutput) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return o
}

func (o EncryptionServicesResponsePtrOutput) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return o
}

func (o EncryptionServicesResponsePtrOutput) Elem() EncryptionServicesResponseOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) EncryptionServicesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionServicesResponse
		return ret
	}).(EncryptionServicesResponseOutput)
}

func (o EncryptionServicesResponsePtrOutput) Blob() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) File() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.File
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) Queue() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return &v.Queue
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) Table() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return &v.Table
	}).(EncryptionServiceResponsePtrOutput)
}

type EndpointsResponse struct {
	Blob  string `pulumi:"blob"`
	Dfs   string `pulumi:"dfs"`
	File  string `pulumi:"file"`
	Queue string `pulumi:"queue"`
	Table string `pulumi:"table"`
	Web   string `pulumi:"web"`
}

type EndpointsResponseOutput struct{ *pulumi.OutputState }

func (EndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) Blob() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Blob }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Dfs() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Dfs }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.File }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Queue() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Queue }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Table() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Table }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Web() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Web }).(pulumi.StringOutput)
}

type GeoReplicationStatsResponse struct {
	CanFailover  bool   `pulumi:"canFailover"`
	LastSyncTime string `pulumi:"lastSyncTime"`
	Status       string `pulumi:"status"`
}

type GeoReplicationStatsResponseOutput struct{ *pulumi.OutputState }

func (GeoReplicationStatsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoReplicationStatsResponse)(nil)).Elem()
}

func (o GeoReplicationStatsResponseOutput) ToGeoReplicationStatsResponseOutput() GeoReplicationStatsResponseOutput {
	return o
}

func (o GeoReplicationStatsResponseOutput) ToGeoReplicationStatsResponseOutputWithContext(ctx context.Context) GeoReplicationStatsResponseOutput {
	return o
}

func (o GeoReplicationStatsResponseOutput) CanFailover() pulumi.BoolOutput {
	return o.ApplyT(func(v GeoReplicationStatsResponse) bool { return v.CanFailover }).(pulumi.BoolOutput)
}

func (o GeoReplicationStatsResponseOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v GeoReplicationStatsResponse) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

func (o GeoReplicationStatsResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GeoReplicationStatsResponse) string { return v.Status }).(pulumi.StringOutput)
}

type IPRule struct {
	Action           *Action `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}


func (val *IPRule) Defaults() *IPRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := Action("Allow")
		tmp.Action = &action_
	}
	return &tmp
}





type IPRuleInput interface {
	pulumi.Input

	ToIPRuleOutput() IPRuleOutput
	ToIPRuleOutputWithContext(context.Context) IPRuleOutput
}

type IPRuleArgs struct {
	Action           ActionPtrInput     `pulumi:"action"`
	IPAddressOrRange pulumi.StringInput `pulumi:"iPAddressOrRange"`
}

func (IPRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (i IPRuleArgs) ToIPRuleOutput() IPRuleOutput {
	return i.ToIPRuleOutputWithContext(context.Background())
}

func (i IPRuleArgs) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleOutput)
}





type IPRuleArrayInput interface {
	pulumi.Input

	ToIPRuleArrayOutput() IPRuleArrayOutput
	ToIPRuleArrayOutputWithContext(context.Context) IPRuleArrayOutput
}

type IPRuleArray []IPRuleInput

func (IPRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (i IPRuleArray) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return i.ToIPRuleArrayOutputWithContext(context.Background())
}

func (i IPRuleArray) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleArrayOutput)
}

type IPRuleOutput struct{ *pulumi.OutputState }

func (IPRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (o IPRuleOutput) ToIPRuleOutput() IPRuleOutput {
	return o
}

func (o IPRuleOutput) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return o
}

func (o IPRuleOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v IPRule) *Action { return v.Action }).(ActionPtrOutput)
}

func (o IPRuleOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRule) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleArrayOutput struct{ *pulumi.OutputState }

func (IPRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) Index(i pulumi.IntInput) IPRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRule {
		return vs[0].([]IPRule)[vs[1].(int)]
	}).(IPRuleOutput)
}

type IPRuleResponse struct {
	Action           *string `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}


func (val *IPRuleResponse) Defaults() *IPRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type IPRuleResponseOutput struct{ *pulumi.OutputState }

func (IPRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IPRuleResponseOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRuleResponse) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IPRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) Index(i pulumi.IntInput) IPRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRuleResponse {
		return vs[0].([]IPRuleResponse)[vs[1].(int)]
	}).(IPRuleResponseOutput)
}

type Identity struct {
	Type IdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type IdentityTypeInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() IdentityTypeOutput {
	return o.ApplyT(func(v Identity) IdentityType { return v.Type }).(IdentityTypeOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() IdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *IdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(IdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ImmutabilityPolicyPropertiesResponse struct {
	Etag                                  string                          `pulumi:"etag"`
	ImmutabilityPeriodSinceCreationInDays int                             `pulumi:"immutabilityPeriodSinceCreationInDays"`
	State                                 string                          `pulumi:"state"`
	UpdateHistory                         []UpdateHistoryPropertyResponse `pulumi:"updateHistory"`
}

type ImmutabilityPolicyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ImmutabilityPolicyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilityPolicyPropertiesResponse)(nil)).Elem()
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ToImmutabilityPolicyPropertiesResponseOutput() ImmutabilityPolicyPropertiesResponseOutput {
	return o
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ToImmutabilityPolicyPropertiesResponseOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponseOutput {
	return o
}

func (o ImmutabilityPolicyPropertiesResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) int { return v.ImmutabilityPeriodSinceCreationInDays }).(pulumi.IntOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) UpdateHistory() UpdateHistoryPropertyResponseArrayOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) []UpdateHistoryPropertyResponse { return v.UpdateHistory }).(UpdateHistoryPropertyResponseArrayOutput)
}

type KeyVaultProperties struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type LegalHoldPropertiesResponse struct {
	HasLegalHold bool                  `pulumi:"hasLegalHold"`
	Tags         []TagPropertyResponse `pulumi:"tags"`
}

type LegalHoldPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LegalHoldPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LegalHoldPropertiesResponse)(nil)).Elem()
}

func (o LegalHoldPropertiesResponseOutput) ToLegalHoldPropertiesResponseOutput() LegalHoldPropertiesResponseOutput {
	return o
}

func (o LegalHoldPropertiesResponseOutput) ToLegalHoldPropertiesResponseOutputWithContext(ctx context.Context) LegalHoldPropertiesResponseOutput {
	return o
}

func (o LegalHoldPropertiesResponseOutput) HasLegalHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LegalHoldPropertiesResponse) bool { return v.HasLegalHold }).(pulumi.BoolOutput)
}

func (o LegalHoldPropertiesResponseOutput) Tags() TagPropertyResponseArrayOutput {
	return o.ApplyT(func(v LegalHoldPropertiesResponse) []TagPropertyResponse { return v.Tags }).(TagPropertyResponseArrayOutput)
}

type ManagementPolicyAction struct {
	BaseBlob *ManagementPolicyBaseBlob `pulumi:"baseBlob"`
	Snapshot *ManagementPolicySnapShot `pulumi:"snapshot"`
}





type ManagementPolicyActionInput interface {
	pulumi.Input

	ToManagementPolicyActionOutput() ManagementPolicyActionOutput
	ToManagementPolicyActionOutputWithContext(context.Context) ManagementPolicyActionOutput
}

type ManagementPolicyActionArgs struct {
	BaseBlob ManagementPolicyBaseBlobPtrInput `pulumi:"baseBlob"`
	Snapshot ManagementPolicySnapShotPtrInput `pulumi:"snapshot"`
}

func (ManagementPolicyActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyAction)(nil)).Elem()
}

func (i ManagementPolicyActionArgs) ToManagementPolicyActionOutput() ManagementPolicyActionOutput {
	return i.ToManagementPolicyActionOutputWithContext(context.Background())
}

func (i ManagementPolicyActionArgs) ToManagementPolicyActionOutputWithContext(ctx context.Context) ManagementPolicyActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyActionOutput)
}

type ManagementPolicyActionOutput struct{ *pulumi.OutputState }

func (ManagementPolicyActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyAction)(nil)).Elem()
}

func (o ManagementPolicyActionOutput) ToManagementPolicyActionOutput() ManagementPolicyActionOutput {
	return o
}

func (o ManagementPolicyActionOutput) ToManagementPolicyActionOutputWithContext(ctx context.Context) ManagementPolicyActionOutput {
	return o
}

func (o ManagementPolicyActionOutput) BaseBlob() ManagementPolicyBaseBlobPtrOutput {
	return o.ApplyT(func(v ManagementPolicyAction) *ManagementPolicyBaseBlob { return v.BaseBlob }).(ManagementPolicyBaseBlobPtrOutput)
}

func (o ManagementPolicyActionOutput) Snapshot() ManagementPolicySnapShotPtrOutput {
	return o.ApplyT(func(v ManagementPolicyAction) *ManagementPolicySnapShot { return v.Snapshot }).(ManagementPolicySnapShotPtrOutput)
}

type ManagementPolicyActionResponse struct {
	BaseBlob *ManagementPolicyBaseBlobResponse `pulumi:"baseBlob"`
	Snapshot *ManagementPolicySnapShotResponse `pulumi:"snapshot"`
}

type ManagementPolicyActionResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyActionResponse)(nil)).Elem()
}

func (o ManagementPolicyActionResponseOutput) ToManagementPolicyActionResponseOutput() ManagementPolicyActionResponseOutput {
	return o
}

func (o ManagementPolicyActionResponseOutput) ToManagementPolicyActionResponseOutputWithContext(ctx context.Context) ManagementPolicyActionResponseOutput {
	return o
}

func (o ManagementPolicyActionResponseOutput) BaseBlob() ManagementPolicyBaseBlobResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyActionResponse) *ManagementPolicyBaseBlobResponse { return v.BaseBlob }).(ManagementPolicyBaseBlobResponsePtrOutput)
}

func (o ManagementPolicyActionResponseOutput) Snapshot() ManagementPolicySnapShotResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyActionResponse) *ManagementPolicySnapShotResponse { return v.Snapshot }).(ManagementPolicySnapShotResponsePtrOutput)
}

type ManagementPolicyBaseBlob struct {
	Delete        *DateAfterModification `pulumi:"delete"`
	TierToArchive *DateAfterModification `pulumi:"tierToArchive"`
	TierToCool    *DateAfterModification `pulumi:"tierToCool"`
}





type ManagementPolicyBaseBlobInput interface {
	pulumi.Input

	ToManagementPolicyBaseBlobOutput() ManagementPolicyBaseBlobOutput
	ToManagementPolicyBaseBlobOutputWithContext(context.Context) ManagementPolicyBaseBlobOutput
}

type ManagementPolicyBaseBlobArgs struct {
	Delete        DateAfterModificationPtrInput `pulumi:"delete"`
	TierToArchive DateAfterModificationPtrInput `pulumi:"tierToArchive"`
	TierToCool    DateAfterModificationPtrInput `pulumi:"tierToCool"`
}

func (ManagementPolicyBaseBlobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyBaseBlob)(nil)).Elem()
}

func (i ManagementPolicyBaseBlobArgs) ToManagementPolicyBaseBlobOutput() ManagementPolicyBaseBlobOutput {
	return i.ToManagementPolicyBaseBlobOutputWithContext(context.Background())
}

func (i ManagementPolicyBaseBlobArgs) ToManagementPolicyBaseBlobOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyBaseBlobOutput)
}

func (i ManagementPolicyBaseBlobArgs) ToManagementPolicyBaseBlobPtrOutput() ManagementPolicyBaseBlobPtrOutput {
	return i.ToManagementPolicyBaseBlobPtrOutputWithContext(context.Background())
}

func (i ManagementPolicyBaseBlobArgs) ToManagementPolicyBaseBlobPtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyBaseBlobOutput).ToManagementPolicyBaseBlobPtrOutputWithContext(ctx)
}









type ManagementPolicyBaseBlobPtrInput interface {
	pulumi.Input

	ToManagementPolicyBaseBlobPtrOutput() ManagementPolicyBaseBlobPtrOutput
	ToManagementPolicyBaseBlobPtrOutputWithContext(context.Context) ManagementPolicyBaseBlobPtrOutput
}

type managementPolicyBaseBlobPtrType ManagementPolicyBaseBlobArgs

func ManagementPolicyBaseBlobPtr(v *ManagementPolicyBaseBlobArgs) ManagementPolicyBaseBlobPtrInput {
	return (*managementPolicyBaseBlobPtrType)(v)
}

func (*managementPolicyBaseBlobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyBaseBlob)(nil)).Elem()
}

func (i *managementPolicyBaseBlobPtrType) ToManagementPolicyBaseBlobPtrOutput() ManagementPolicyBaseBlobPtrOutput {
	return i.ToManagementPolicyBaseBlobPtrOutputWithContext(context.Background())
}

func (i *managementPolicyBaseBlobPtrType) ToManagementPolicyBaseBlobPtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyBaseBlobPtrOutput)
}

type ManagementPolicyBaseBlobOutput struct{ *pulumi.OutputState }

func (ManagementPolicyBaseBlobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyBaseBlob)(nil)).Elem()
}

func (o ManagementPolicyBaseBlobOutput) ToManagementPolicyBaseBlobOutput() ManagementPolicyBaseBlobOutput {
	return o
}

func (o ManagementPolicyBaseBlobOutput) ToManagementPolicyBaseBlobOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobOutput {
	return o
}

func (o ManagementPolicyBaseBlobOutput) ToManagementPolicyBaseBlobPtrOutput() ManagementPolicyBaseBlobPtrOutput {
	return o.ToManagementPolicyBaseBlobPtrOutputWithContext(context.Background())
}

func (o ManagementPolicyBaseBlobOutput) ToManagementPolicyBaseBlobPtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicyBaseBlob) *ManagementPolicyBaseBlob {
		return &v
	}).(ManagementPolicyBaseBlobPtrOutput)
}

func (o ManagementPolicyBaseBlobOutput) Delete() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlob) *DateAfterModification { return v.Delete }).(DateAfterModificationPtrOutput)
}

func (o ManagementPolicyBaseBlobOutput) TierToArchive() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlob) *DateAfterModification { return v.TierToArchive }).(DateAfterModificationPtrOutput)
}

func (o ManagementPolicyBaseBlobOutput) TierToCool() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlob) *DateAfterModification { return v.TierToCool }).(DateAfterModificationPtrOutput)
}

type ManagementPolicyBaseBlobPtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicyBaseBlobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyBaseBlob)(nil)).Elem()
}

func (o ManagementPolicyBaseBlobPtrOutput) ToManagementPolicyBaseBlobPtrOutput() ManagementPolicyBaseBlobPtrOutput {
	return o
}

func (o ManagementPolicyBaseBlobPtrOutput) ToManagementPolicyBaseBlobPtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobPtrOutput {
	return o
}

func (o ManagementPolicyBaseBlobPtrOutput) Elem() ManagementPolicyBaseBlobOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlob) ManagementPolicyBaseBlob {
		if v != nil {
			return *v
		}
		var ret ManagementPolicyBaseBlob
		return ret
	}).(ManagementPolicyBaseBlobOutput)
}

func (o ManagementPolicyBaseBlobPtrOutput) Delete() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlob) *DateAfterModification {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(DateAfterModificationPtrOutput)
}

func (o ManagementPolicyBaseBlobPtrOutput) TierToArchive() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlob) *DateAfterModification {
		if v == nil {
			return nil
		}
		return v.TierToArchive
	}).(DateAfterModificationPtrOutput)
}

func (o ManagementPolicyBaseBlobPtrOutput) TierToCool() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlob) *DateAfterModification {
		if v == nil {
			return nil
		}
		return v.TierToCool
	}).(DateAfterModificationPtrOutput)
}

type ManagementPolicyBaseBlobResponse struct {
	Delete        *DateAfterModificationResponse `pulumi:"delete"`
	TierToArchive *DateAfterModificationResponse `pulumi:"tierToArchive"`
	TierToCool    *DateAfterModificationResponse `pulumi:"tierToCool"`
}

type ManagementPolicyBaseBlobResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyBaseBlobResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyBaseBlobResponse)(nil)).Elem()
}

func (o ManagementPolicyBaseBlobResponseOutput) ToManagementPolicyBaseBlobResponseOutput() ManagementPolicyBaseBlobResponseOutput {
	return o
}

func (o ManagementPolicyBaseBlobResponseOutput) ToManagementPolicyBaseBlobResponseOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobResponseOutput {
	return o
}

func (o ManagementPolicyBaseBlobResponseOutput) Delete() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse { return v.Delete }).(DateAfterModificationResponsePtrOutput)
}

func (o ManagementPolicyBaseBlobResponseOutput) TierToArchive() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse { return v.TierToArchive }).(DateAfterModificationResponsePtrOutput)
}

func (o ManagementPolicyBaseBlobResponseOutput) TierToCool() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse { return v.TierToCool }).(DateAfterModificationResponsePtrOutput)
}

type ManagementPolicyBaseBlobResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicyBaseBlobResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyBaseBlobResponse)(nil)).Elem()
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) ToManagementPolicyBaseBlobResponsePtrOutput() ManagementPolicyBaseBlobResponsePtrOutput {
	return o
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) ToManagementPolicyBaseBlobResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobResponsePtrOutput {
	return o
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) Elem() ManagementPolicyBaseBlobResponseOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlobResponse) ManagementPolicyBaseBlobResponse {
		if v != nil {
			return *v
		}
		var ret ManagementPolicyBaseBlobResponse
		return ret
	}).(ManagementPolicyBaseBlobResponseOutput)
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) Delete() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(DateAfterModificationResponsePtrOutput)
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) TierToArchive() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse {
		if v == nil {
			return nil
		}
		return v.TierToArchive
	}).(DateAfterModificationResponsePtrOutput)
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) TierToCool() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse {
		if v == nil {
			return nil
		}
		return v.TierToCool
	}).(DateAfterModificationResponsePtrOutput)
}

type ManagementPolicyDefinition struct {
	Actions ManagementPolicyAction  `pulumi:"actions"`
	Filters *ManagementPolicyFilter `pulumi:"filters"`
}





type ManagementPolicyDefinitionInput interface {
	pulumi.Input

	ToManagementPolicyDefinitionOutput() ManagementPolicyDefinitionOutput
	ToManagementPolicyDefinitionOutputWithContext(context.Context) ManagementPolicyDefinitionOutput
}

type ManagementPolicyDefinitionArgs struct {
	Actions ManagementPolicyActionInput    `pulumi:"actions"`
	Filters ManagementPolicyFilterPtrInput `pulumi:"filters"`
}

func (ManagementPolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyDefinition)(nil)).Elem()
}

func (i ManagementPolicyDefinitionArgs) ToManagementPolicyDefinitionOutput() ManagementPolicyDefinitionOutput {
	return i.ToManagementPolicyDefinitionOutputWithContext(context.Background())
}

func (i ManagementPolicyDefinitionArgs) ToManagementPolicyDefinitionOutputWithContext(ctx context.Context) ManagementPolicyDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyDefinitionOutput)
}

type ManagementPolicyDefinitionOutput struct{ *pulumi.OutputState }

func (ManagementPolicyDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyDefinition)(nil)).Elem()
}

func (o ManagementPolicyDefinitionOutput) ToManagementPolicyDefinitionOutput() ManagementPolicyDefinitionOutput {
	return o
}

func (o ManagementPolicyDefinitionOutput) ToManagementPolicyDefinitionOutputWithContext(ctx context.Context) ManagementPolicyDefinitionOutput {
	return o
}

func (o ManagementPolicyDefinitionOutput) Actions() ManagementPolicyActionOutput {
	return o.ApplyT(func(v ManagementPolicyDefinition) ManagementPolicyAction { return v.Actions }).(ManagementPolicyActionOutput)
}

func (o ManagementPolicyDefinitionOutput) Filters() ManagementPolicyFilterPtrOutput {
	return o.ApplyT(func(v ManagementPolicyDefinition) *ManagementPolicyFilter { return v.Filters }).(ManagementPolicyFilterPtrOutput)
}

type ManagementPolicyDefinitionResponse struct {
	Actions ManagementPolicyActionResponse  `pulumi:"actions"`
	Filters *ManagementPolicyFilterResponse `pulumi:"filters"`
}

type ManagementPolicyDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyDefinitionResponse)(nil)).Elem()
}

func (o ManagementPolicyDefinitionResponseOutput) ToManagementPolicyDefinitionResponseOutput() ManagementPolicyDefinitionResponseOutput {
	return o
}

func (o ManagementPolicyDefinitionResponseOutput) ToManagementPolicyDefinitionResponseOutputWithContext(ctx context.Context) ManagementPolicyDefinitionResponseOutput {
	return o
}

func (o ManagementPolicyDefinitionResponseOutput) Actions() ManagementPolicyActionResponseOutput {
	return o.ApplyT(func(v ManagementPolicyDefinitionResponse) ManagementPolicyActionResponse { return v.Actions }).(ManagementPolicyActionResponseOutput)
}

func (o ManagementPolicyDefinitionResponseOutput) Filters() ManagementPolicyFilterResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyDefinitionResponse) *ManagementPolicyFilterResponse { return v.Filters }).(ManagementPolicyFilterResponsePtrOutput)
}

type ManagementPolicyFilter struct {
	BlobTypes   []string `pulumi:"blobTypes"`
	PrefixMatch []string `pulumi:"prefixMatch"`
}





type ManagementPolicyFilterInput interface {
	pulumi.Input

	ToManagementPolicyFilterOutput() ManagementPolicyFilterOutput
	ToManagementPolicyFilterOutputWithContext(context.Context) ManagementPolicyFilterOutput
}

type ManagementPolicyFilterArgs struct {
	BlobTypes   pulumi.StringArrayInput `pulumi:"blobTypes"`
	PrefixMatch pulumi.StringArrayInput `pulumi:"prefixMatch"`
}

func (ManagementPolicyFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyFilter)(nil)).Elem()
}

func (i ManagementPolicyFilterArgs) ToManagementPolicyFilterOutput() ManagementPolicyFilterOutput {
	return i.ToManagementPolicyFilterOutputWithContext(context.Background())
}

func (i ManagementPolicyFilterArgs) ToManagementPolicyFilterOutputWithContext(ctx context.Context) ManagementPolicyFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyFilterOutput)
}

func (i ManagementPolicyFilterArgs) ToManagementPolicyFilterPtrOutput() ManagementPolicyFilterPtrOutput {
	return i.ToManagementPolicyFilterPtrOutputWithContext(context.Background())
}

func (i ManagementPolicyFilterArgs) ToManagementPolicyFilterPtrOutputWithContext(ctx context.Context) ManagementPolicyFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyFilterOutput).ToManagementPolicyFilterPtrOutputWithContext(ctx)
}









type ManagementPolicyFilterPtrInput interface {
	pulumi.Input

	ToManagementPolicyFilterPtrOutput() ManagementPolicyFilterPtrOutput
	ToManagementPolicyFilterPtrOutputWithContext(context.Context) ManagementPolicyFilterPtrOutput
}

type managementPolicyFilterPtrType ManagementPolicyFilterArgs

func ManagementPolicyFilterPtr(v *ManagementPolicyFilterArgs) ManagementPolicyFilterPtrInput {
	return (*managementPolicyFilterPtrType)(v)
}

func (*managementPolicyFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyFilter)(nil)).Elem()
}

func (i *managementPolicyFilterPtrType) ToManagementPolicyFilterPtrOutput() ManagementPolicyFilterPtrOutput {
	return i.ToManagementPolicyFilterPtrOutputWithContext(context.Background())
}

func (i *managementPolicyFilterPtrType) ToManagementPolicyFilterPtrOutputWithContext(ctx context.Context) ManagementPolicyFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyFilterPtrOutput)
}

type ManagementPolicyFilterOutput struct{ *pulumi.OutputState }

func (ManagementPolicyFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyFilter)(nil)).Elem()
}

func (o ManagementPolicyFilterOutput) ToManagementPolicyFilterOutput() ManagementPolicyFilterOutput {
	return o
}

func (o ManagementPolicyFilterOutput) ToManagementPolicyFilterOutputWithContext(ctx context.Context) ManagementPolicyFilterOutput {
	return o
}

func (o ManagementPolicyFilterOutput) ToManagementPolicyFilterPtrOutput() ManagementPolicyFilterPtrOutput {
	return o.ToManagementPolicyFilterPtrOutputWithContext(context.Background())
}

func (o ManagementPolicyFilterOutput) ToManagementPolicyFilterPtrOutputWithContext(ctx context.Context) ManagementPolicyFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicyFilter) *ManagementPolicyFilter {
		return &v
	}).(ManagementPolicyFilterPtrOutput)
}

func (o ManagementPolicyFilterOutput) BlobTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementPolicyFilter) []string { return v.BlobTypes }).(pulumi.StringArrayOutput)
}

func (o ManagementPolicyFilterOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementPolicyFilter) []string { return v.PrefixMatch }).(pulumi.StringArrayOutput)
}

type ManagementPolicyFilterPtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicyFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyFilter)(nil)).Elem()
}

func (o ManagementPolicyFilterPtrOutput) ToManagementPolicyFilterPtrOutput() ManagementPolicyFilterPtrOutput {
	return o
}

func (o ManagementPolicyFilterPtrOutput) ToManagementPolicyFilterPtrOutputWithContext(ctx context.Context) ManagementPolicyFilterPtrOutput {
	return o
}

func (o ManagementPolicyFilterPtrOutput) Elem() ManagementPolicyFilterOutput {
	return o.ApplyT(func(v *ManagementPolicyFilter) ManagementPolicyFilter {
		if v != nil {
			return *v
		}
		var ret ManagementPolicyFilter
		return ret
	}).(ManagementPolicyFilterOutput)
}

func (o ManagementPolicyFilterPtrOutput) BlobTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementPolicyFilter) []string {
		if v == nil {
			return nil
		}
		return v.BlobTypes
	}).(pulumi.StringArrayOutput)
}

func (o ManagementPolicyFilterPtrOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementPolicyFilter) []string {
		if v == nil {
			return nil
		}
		return v.PrefixMatch
	}).(pulumi.StringArrayOutput)
}

type ManagementPolicyFilterResponse struct {
	BlobTypes   []string `pulumi:"blobTypes"`
	PrefixMatch []string `pulumi:"prefixMatch"`
}

type ManagementPolicyFilterResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyFilterResponse)(nil)).Elem()
}

func (o ManagementPolicyFilterResponseOutput) ToManagementPolicyFilterResponseOutput() ManagementPolicyFilterResponseOutput {
	return o
}

func (o ManagementPolicyFilterResponseOutput) ToManagementPolicyFilterResponseOutputWithContext(ctx context.Context) ManagementPolicyFilterResponseOutput {
	return o
}

func (o ManagementPolicyFilterResponseOutput) BlobTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementPolicyFilterResponse) []string { return v.BlobTypes }).(pulumi.StringArrayOutput)
}

func (o ManagementPolicyFilterResponseOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementPolicyFilterResponse) []string { return v.PrefixMatch }).(pulumi.StringArrayOutput)
}

type ManagementPolicyFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicyFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyFilterResponse)(nil)).Elem()
}

func (o ManagementPolicyFilterResponsePtrOutput) ToManagementPolicyFilterResponsePtrOutput() ManagementPolicyFilterResponsePtrOutput {
	return o
}

func (o ManagementPolicyFilterResponsePtrOutput) ToManagementPolicyFilterResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyFilterResponsePtrOutput {
	return o
}

func (o ManagementPolicyFilterResponsePtrOutput) Elem() ManagementPolicyFilterResponseOutput {
	return o.ApplyT(func(v *ManagementPolicyFilterResponse) ManagementPolicyFilterResponse {
		if v != nil {
			return *v
		}
		var ret ManagementPolicyFilterResponse
		return ret
	}).(ManagementPolicyFilterResponseOutput)
}

func (o ManagementPolicyFilterResponsePtrOutput) BlobTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementPolicyFilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.BlobTypes
	}).(pulumi.StringArrayOutput)
}

func (o ManagementPolicyFilterResponsePtrOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementPolicyFilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.PrefixMatch
	}).(pulumi.StringArrayOutput)
}

type ManagementPolicyRule struct {
	Definition ManagementPolicyDefinition `pulumi:"definition"`
	Enabled    *bool                      `pulumi:"enabled"`
	Name       string                     `pulumi:"name"`
	Type       string                     `pulumi:"type"`
}





type ManagementPolicyRuleInput interface {
	pulumi.Input

	ToManagementPolicyRuleOutput() ManagementPolicyRuleOutput
	ToManagementPolicyRuleOutputWithContext(context.Context) ManagementPolicyRuleOutput
}

type ManagementPolicyRuleArgs struct {
	Definition ManagementPolicyDefinitionInput `pulumi:"definition"`
	Enabled    pulumi.BoolPtrInput             `pulumi:"enabled"`
	Name       pulumi.StringInput              `pulumi:"name"`
	Type       pulumi.StringInput              `pulumi:"type"`
}

func (ManagementPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyRule)(nil)).Elem()
}

func (i ManagementPolicyRuleArgs) ToManagementPolicyRuleOutput() ManagementPolicyRuleOutput {
	return i.ToManagementPolicyRuleOutputWithContext(context.Background())
}

func (i ManagementPolicyRuleArgs) ToManagementPolicyRuleOutputWithContext(ctx context.Context) ManagementPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyRuleOutput)
}





type ManagementPolicyRuleArrayInput interface {
	pulumi.Input

	ToManagementPolicyRuleArrayOutput() ManagementPolicyRuleArrayOutput
	ToManagementPolicyRuleArrayOutputWithContext(context.Context) ManagementPolicyRuleArrayOutput
}

type ManagementPolicyRuleArray []ManagementPolicyRuleInput

func (ManagementPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementPolicyRule)(nil)).Elem()
}

func (i ManagementPolicyRuleArray) ToManagementPolicyRuleArrayOutput() ManagementPolicyRuleArrayOutput {
	return i.ToManagementPolicyRuleArrayOutputWithContext(context.Background())
}

func (i ManagementPolicyRuleArray) ToManagementPolicyRuleArrayOutputWithContext(ctx context.Context) ManagementPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyRuleArrayOutput)
}

type ManagementPolicyRuleOutput struct{ *pulumi.OutputState }

func (ManagementPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyRule)(nil)).Elem()
}

func (o ManagementPolicyRuleOutput) ToManagementPolicyRuleOutput() ManagementPolicyRuleOutput {
	return o
}

func (o ManagementPolicyRuleOutput) ToManagementPolicyRuleOutputWithContext(ctx context.Context) ManagementPolicyRuleOutput {
	return o
}

func (o ManagementPolicyRuleOutput) Definition() ManagementPolicyDefinitionOutput {
	return o.ApplyT(func(v ManagementPolicyRule) ManagementPolicyDefinition { return v.Definition }).(ManagementPolicyDefinitionOutput)
}

func (o ManagementPolicyRuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagementPolicyRule) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ManagementPolicyRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementPolicyRule) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementPolicyRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementPolicyRule) string { return v.Type }).(pulumi.StringOutput)
}

type ManagementPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (ManagementPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementPolicyRule)(nil)).Elem()
}

func (o ManagementPolicyRuleArrayOutput) ToManagementPolicyRuleArrayOutput() ManagementPolicyRuleArrayOutput {
	return o
}

func (o ManagementPolicyRuleArrayOutput) ToManagementPolicyRuleArrayOutputWithContext(ctx context.Context) ManagementPolicyRuleArrayOutput {
	return o
}

func (o ManagementPolicyRuleArrayOutput) Index(i pulumi.IntInput) ManagementPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementPolicyRule {
		return vs[0].([]ManagementPolicyRule)[vs[1].(int)]
	}).(ManagementPolicyRuleOutput)
}

type ManagementPolicyRuleResponse struct {
	Definition ManagementPolicyDefinitionResponse `pulumi:"definition"`
	Enabled    *bool                              `pulumi:"enabled"`
	Name       string                             `pulumi:"name"`
	Type       string                             `pulumi:"type"`
}

type ManagementPolicyRuleResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyRuleResponse)(nil)).Elem()
}

func (o ManagementPolicyRuleResponseOutput) ToManagementPolicyRuleResponseOutput() ManagementPolicyRuleResponseOutput {
	return o
}

func (o ManagementPolicyRuleResponseOutput) ToManagementPolicyRuleResponseOutputWithContext(ctx context.Context) ManagementPolicyRuleResponseOutput {
	return o
}

func (o ManagementPolicyRuleResponseOutput) Definition() ManagementPolicyDefinitionResponseOutput {
	return o.ApplyT(func(v ManagementPolicyRuleResponse) ManagementPolicyDefinitionResponse { return v.Definition }).(ManagementPolicyDefinitionResponseOutput)
}

func (o ManagementPolicyRuleResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagementPolicyRuleResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ManagementPolicyRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementPolicyRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementPolicyRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementPolicyRuleResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ManagementPolicyRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementPolicyRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementPolicyRuleResponse)(nil)).Elem()
}

func (o ManagementPolicyRuleResponseArrayOutput) ToManagementPolicyRuleResponseArrayOutput() ManagementPolicyRuleResponseArrayOutput {
	return o
}

func (o ManagementPolicyRuleResponseArrayOutput) ToManagementPolicyRuleResponseArrayOutputWithContext(ctx context.Context) ManagementPolicyRuleResponseArrayOutput {
	return o
}

func (o ManagementPolicyRuleResponseArrayOutput) Index(i pulumi.IntInput) ManagementPolicyRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementPolicyRuleResponse {
		return vs[0].([]ManagementPolicyRuleResponse)[vs[1].(int)]
	}).(ManagementPolicyRuleResponseOutput)
}

type ManagementPolicySchema struct {
	Rules []ManagementPolicyRule `pulumi:"rules"`
}





type ManagementPolicySchemaInput interface {
	pulumi.Input

	ToManagementPolicySchemaOutput() ManagementPolicySchemaOutput
	ToManagementPolicySchemaOutputWithContext(context.Context) ManagementPolicySchemaOutput
}

type ManagementPolicySchemaArgs struct {
	Rules ManagementPolicyRuleArrayInput `pulumi:"rules"`
}

func (ManagementPolicySchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySchema)(nil)).Elem()
}

func (i ManagementPolicySchemaArgs) ToManagementPolicySchemaOutput() ManagementPolicySchemaOutput {
	return i.ToManagementPolicySchemaOutputWithContext(context.Background())
}

func (i ManagementPolicySchemaArgs) ToManagementPolicySchemaOutputWithContext(ctx context.Context) ManagementPolicySchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySchemaOutput)
}

type ManagementPolicySchemaOutput struct{ *pulumi.OutputState }

func (ManagementPolicySchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySchema)(nil)).Elem()
}

func (o ManagementPolicySchemaOutput) ToManagementPolicySchemaOutput() ManagementPolicySchemaOutput {
	return o
}

func (o ManagementPolicySchemaOutput) ToManagementPolicySchemaOutputWithContext(ctx context.Context) ManagementPolicySchemaOutput {
	return o
}

func (o ManagementPolicySchemaOutput) Rules() ManagementPolicyRuleArrayOutput {
	return o.ApplyT(func(v ManagementPolicySchema) []ManagementPolicyRule { return v.Rules }).(ManagementPolicyRuleArrayOutput)
}

type ManagementPolicySchemaResponse struct {
	Rules []ManagementPolicyRuleResponse `pulumi:"rules"`
}

type ManagementPolicySchemaResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicySchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySchemaResponse)(nil)).Elem()
}

func (o ManagementPolicySchemaResponseOutput) ToManagementPolicySchemaResponseOutput() ManagementPolicySchemaResponseOutput {
	return o
}

func (o ManagementPolicySchemaResponseOutput) ToManagementPolicySchemaResponseOutputWithContext(ctx context.Context) ManagementPolicySchemaResponseOutput {
	return o
}

func (o ManagementPolicySchemaResponseOutput) Rules() ManagementPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v ManagementPolicySchemaResponse) []ManagementPolicyRuleResponse { return v.Rules }).(ManagementPolicyRuleResponseArrayOutput)
}

type ManagementPolicySnapShot struct {
	Delete *DateAfterCreation `pulumi:"delete"`
}





type ManagementPolicySnapShotInput interface {
	pulumi.Input

	ToManagementPolicySnapShotOutput() ManagementPolicySnapShotOutput
	ToManagementPolicySnapShotOutputWithContext(context.Context) ManagementPolicySnapShotOutput
}

type ManagementPolicySnapShotArgs struct {
	Delete DateAfterCreationPtrInput `pulumi:"delete"`
}

func (ManagementPolicySnapShotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySnapShot)(nil)).Elem()
}

func (i ManagementPolicySnapShotArgs) ToManagementPolicySnapShotOutput() ManagementPolicySnapShotOutput {
	return i.ToManagementPolicySnapShotOutputWithContext(context.Background())
}

func (i ManagementPolicySnapShotArgs) ToManagementPolicySnapShotOutputWithContext(ctx context.Context) ManagementPolicySnapShotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySnapShotOutput)
}

func (i ManagementPolicySnapShotArgs) ToManagementPolicySnapShotPtrOutput() ManagementPolicySnapShotPtrOutput {
	return i.ToManagementPolicySnapShotPtrOutputWithContext(context.Background())
}

func (i ManagementPolicySnapShotArgs) ToManagementPolicySnapShotPtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySnapShotOutput).ToManagementPolicySnapShotPtrOutputWithContext(ctx)
}









type ManagementPolicySnapShotPtrInput interface {
	pulumi.Input

	ToManagementPolicySnapShotPtrOutput() ManagementPolicySnapShotPtrOutput
	ToManagementPolicySnapShotPtrOutputWithContext(context.Context) ManagementPolicySnapShotPtrOutput
}

type managementPolicySnapShotPtrType ManagementPolicySnapShotArgs

func ManagementPolicySnapShotPtr(v *ManagementPolicySnapShotArgs) ManagementPolicySnapShotPtrInput {
	return (*managementPolicySnapShotPtrType)(v)
}

func (*managementPolicySnapShotPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySnapShot)(nil)).Elem()
}

func (i *managementPolicySnapShotPtrType) ToManagementPolicySnapShotPtrOutput() ManagementPolicySnapShotPtrOutput {
	return i.ToManagementPolicySnapShotPtrOutputWithContext(context.Background())
}

func (i *managementPolicySnapShotPtrType) ToManagementPolicySnapShotPtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySnapShotPtrOutput)
}

type ManagementPolicySnapShotOutput struct{ *pulumi.OutputState }

func (ManagementPolicySnapShotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySnapShot)(nil)).Elem()
}

func (o ManagementPolicySnapShotOutput) ToManagementPolicySnapShotOutput() ManagementPolicySnapShotOutput {
	return o
}

func (o ManagementPolicySnapShotOutput) ToManagementPolicySnapShotOutputWithContext(ctx context.Context) ManagementPolicySnapShotOutput {
	return o
}

func (o ManagementPolicySnapShotOutput) ToManagementPolicySnapShotPtrOutput() ManagementPolicySnapShotPtrOutput {
	return o.ToManagementPolicySnapShotPtrOutputWithContext(context.Background())
}

func (o ManagementPolicySnapShotOutput) ToManagementPolicySnapShotPtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicySnapShot) *ManagementPolicySnapShot {
		return &v
	}).(ManagementPolicySnapShotPtrOutput)
}

func (o ManagementPolicySnapShotOutput) Delete() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v ManagementPolicySnapShot) *DateAfterCreation { return v.Delete }).(DateAfterCreationPtrOutput)
}

type ManagementPolicySnapShotPtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicySnapShotPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySnapShot)(nil)).Elem()
}

func (o ManagementPolicySnapShotPtrOutput) ToManagementPolicySnapShotPtrOutput() ManagementPolicySnapShotPtrOutput {
	return o
}

func (o ManagementPolicySnapShotPtrOutput) ToManagementPolicySnapShotPtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotPtrOutput {
	return o
}

func (o ManagementPolicySnapShotPtrOutput) Elem() ManagementPolicySnapShotOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShot) ManagementPolicySnapShot {
		if v != nil {
			return *v
		}
		var ret ManagementPolicySnapShot
		return ret
	}).(ManagementPolicySnapShotOutput)
}

func (o ManagementPolicySnapShotPtrOutput) Delete() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShot) *DateAfterCreation {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(DateAfterCreationPtrOutput)
}

type ManagementPolicySnapShotResponse struct {
	Delete *DateAfterCreationResponse `pulumi:"delete"`
}

type ManagementPolicySnapShotResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicySnapShotResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySnapShotResponse)(nil)).Elem()
}

func (o ManagementPolicySnapShotResponseOutput) ToManagementPolicySnapShotResponseOutput() ManagementPolicySnapShotResponseOutput {
	return o
}

func (o ManagementPolicySnapShotResponseOutput) ToManagementPolicySnapShotResponseOutputWithContext(ctx context.Context) ManagementPolicySnapShotResponseOutput {
	return o
}

func (o ManagementPolicySnapShotResponseOutput) Delete() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicySnapShotResponse) *DateAfterCreationResponse { return v.Delete }).(DateAfterCreationResponsePtrOutput)
}

type ManagementPolicySnapShotResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicySnapShotResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySnapShotResponse)(nil)).Elem()
}

func (o ManagementPolicySnapShotResponsePtrOutput) ToManagementPolicySnapShotResponsePtrOutput() ManagementPolicySnapShotResponsePtrOutput {
	return o
}

func (o ManagementPolicySnapShotResponsePtrOutput) ToManagementPolicySnapShotResponsePtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotResponsePtrOutput {
	return o
}

func (o ManagementPolicySnapShotResponsePtrOutput) Elem() ManagementPolicySnapShotResponseOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShotResponse) ManagementPolicySnapShotResponse {
		if v != nil {
			return *v
		}
		var ret ManagementPolicySnapShotResponse
		return ret
	}).(ManagementPolicySnapShotResponseOutput)
}

func (o ManagementPolicySnapShotResponsePtrOutput) Delete() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShotResponse) *DateAfterCreationResponse {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(DateAfterCreationResponsePtrOutput)
}

type NetworkRuleSet struct {
	Bypass              *string              `pulumi:"bypass"`
	DefaultAction       DefaultAction        `pulumi:"defaultAction"`
	IpRules             []IPRule             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}


func (val *NetworkRuleSet) Defaults() *NetworkRuleSet {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Bypass) {
		bypass_ := "AzureServices"
		tmp.Bypass = &bypass_
	}
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = DefaultAction("Allow")
	}
	return &tmp
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	Bypass              pulumi.StringPtrInput        `pulumi:"bypass"`
	DefaultAction       DefaultActionInput           `pulumi:"defaultAction"`
	IpRules             IPRuleArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSet) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetOutput) DefaultAction() DefaultActionOutput {
	return o.ApplyT(func(v NetworkRuleSet) DefaultAction { return v.DefaultAction }).(DefaultActionOutput)
}

func (o NetworkRuleSetOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IPRule { return v.IpRules }).(IPRuleArrayOutput)
}

func (o NetworkRuleSetOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []VirtualNetworkRule { return v.VirtualNetworkRules }).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) DefaultAction() DefaultActionPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *DefaultAction {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(DefaultActionPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IPRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleArrayOutput)
}

func (o NetworkRuleSetPtrOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []VirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	Bypass              *string                      `pulumi:"bypass"`
	DefaultAction       string                       `pulumi:"defaultAction"`
	IpRules             []IPRuleResponse             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}


func (val *NetworkRuleSetResponse) Defaults() *NetworkRuleSetResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Bypass) {
		bypass_ := "AzureServices"
		tmp.Bypass = &bypass_
	}
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = "Allow"
	}
	return &tmp
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponseOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) string { return v.DefaultAction }).(pulumi.StringOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IPRuleResponse { return v.IpRules }).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

type Restriction struct {
	ReasonCode *string `pulumi:"reasonCode"`
}





type RestrictionInput interface {
	pulumi.Input

	ToRestrictionOutput() RestrictionOutput
	ToRestrictionOutputWithContext(context.Context) RestrictionOutput
}

type RestrictionArgs struct {
	ReasonCode pulumi.StringPtrInput `pulumi:"reasonCode"`
}

func (RestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Restriction)(nil)).Elem()
}

func (i RestrictionArgs) ToRestrictionOutput() RestrictionOutput {
	return i.ToRestrictionOutputWithContext(context.Background())
}

func (i RestrictionArgs) ToRestrictionOutputWithContext(ctx context.Context) RestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestrictionOutput)
}





type RestrictionArrayInput interface {
	pulumi.Input

	ToRestrictionArrayOutput() RestrictionArrayOutput
	ToRestrictionArrayOutputWithContext(context.Context) RestrictionArrayOutput
}

type RestrictionArray []RestrictionInput

func (RestrictionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Restriction)(nil)).Elem()
}

func (i RestrictionArray) ToRestrictionArrayOutput() RestrictionArrayOutput {
	return i.ToRestrictionArrayOutputWithContext(context.Background())
}

func (i RestrictionArray) ToRestrictionArrayOutputWithContext(ctx context.Context) RestrictionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestrictionArrayOutput)
}

type RestrictionOutput struct{ *pulumi.OutputState }

func (RestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Restriction)(nil)).Elem()
}

func (o RestrictionOutput) ToRestrictionOutput() RestrictionOutput {
	return o
}

func (o RestrictionOutput) ToRestrictionOutputWithContext(ctx context.Context) RestrictionOutput {
	return o
}

func (o RestrictionOutput) ReasonCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Restriction) *string { return v.ReasonCode }).(pulumi.StringPtrOutput)
}

type RestrictionArrayOutput struct{ *pulumi.OutputState }

func (RestrictionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Restriction)(nil)).Elem()
}

func (o RestrictionArrayOutput) ToRestrictionArrayOutput() RestrictionArrayOutput {
	return o
}

func (o RestrictionArrayOutput) ToRestrictionArrayOutputWithContext(ctx context.Context) RestrictionArrayOutput {
	return o
}

func (o RestrictionArrayOutput) Index(i pulumi.IntInput) RestrictionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Restriction {
		return vs[0].([]Restriction)[vs[1].(int)]
	}).(RestrictionOutput)
}

type RestrictionResponse struct {
	ReasonCode *string  `pulumi:"reasonCode"`
	Type       string   `pulumi:"type"`
	Values     []string `pulumi:"values"`
}

type RestrictionResponseOutput struct{ *pulumi.OutputState }

func (RestrictionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestrictionResponse)(nil)).Elem()
}

func (o RestrictionResponseOutput) ToRestrictionResponseOutput() RestrictionResponseOutput {
	return o
}

func (o RestrictionResponseOutput) ToRestrictionResponseOutputWithContext(ctx context.Context) RestrictionResponseOutput {
	return o
}

func (o RestrictionResponseOutput) ReasonCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestrictionResponse) *string { return v.ReasonCode }).(pulumi.StringPtrOutput)
}

func (o RestrictionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RestrictionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o RestrictionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RestrictionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type RestrictionResponseArrayOutput struct{ *pulumi.OutputState }

func (RestrictionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RestrictionResponse)(nil)).Elem()
}

func (o RestrictionResponseArrayOutput) ToRestrictionResponseArrayOutput() RestrictionResponseArrayOutput {
	return o
}

func (o RestrictionResponseArrayOutput) ToRestrictionResponseArrayOutputWithContext(ctx context.Context) RestrictionResponseArrayOutput {
	return o
}

func (o RestrictionResponseArrayOutput) Index(i pulumi.IntInput) RestrictionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RestrictionResponse {
		return vs[0].([]RestrictionResponse)[vs[1].(int)]
	}).(RestrictionResponseOutput)
}

type SKUCapabilityResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type SKUCapabilityResponseOutput struct{ *pulumi.OutputState }

func (SKUCapabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SKUCapabilityResponse)(nil)).Elem()
}

func (o SKUCapabilityResponseOutput) ToSKUCapabilityResponseOutput() SKUCapabilityResponseOutput {
	return o
}

func (o SKUCapabilityResponseOutput) ToSKUCapabilityResponseOutputWithContext(ctx context.Context) SKUCapabilityResponseOutput {
	return o
}

func (o SKUCapabilityResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SKUCapabilityResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SKUCapabilityResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SKUCapabilityResponse) string { return v.Value }).(pulumi.StringOutput)
}

type SKUCapabilityResponseArrayOutput struct{ *pulumi.OutputState }

func (SKUCapabilityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SKUCapabilityResponse)(nil)).Elem()
}

func (o SKUCapabilityResponseArrayOutput) ToSKUCapabilityResponseArrayOutput() SKUCapabilityResponseArrayOutput {
	return o
}

func (o SKUCapabilityResponseArrayOutput) ToSKUCapabilityResponseArrayOutputWithContext(ctx context.Context) SKUCapabilityResponseArrayOutput {
	return o
}

func (o SKUCapabilityResponseArrayOutput) Index(i pulumi.IntInput) SKUCapabilityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SKUCapabilityResponse {
		return vs[0].([]SKUCapabilityResponse)[vs[1].(int)]
	}).(SKUCapabilityResponseOutput)
}

type Sku struct {
	Name         SkuName       `pulumi:"name"`
	Restrictions []Restriction `pulumi:"restrictions"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name         SkuNameInput          `pulumi:"name"`
	Restrictions RestrictionArrayInput `pulumi:"restrictions"`
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

func (o SkuOutput) Name() SkuNameOutput {
	return o.ApplyT(func(v Sku) SkuName { return v.Name }).(SkuNameOutput)
}

func (o SkuOutput) Restrictions() RestrictionArrayOutput {
	return o.ApplyT(func(v Sku) []Restriction { return v.Restrictions }).(RestrictionArrayOutput)
}

type SkuResponse struct {
	Capabilities []SKUCapabilityResponse `pulumi:"capabilities"`
	Kind         string                  `pulumi:"kind"`
	Locations    []string                `pulumi:"locations"`
	Name         string                  `pulumi:"name"`
	ResourceType string                  `pulumi:"resourceType"`
	Restrictions []RestrictionResponse   `pulumi:"restrictions"`
	Tier         string                  `pulumi:"tier"`
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

func (o SkuResponseOutput) Capabilities() SKUCapabilityResponseArrayOutput {
	return o.ApplyT(func(v SkuResponse) []SKUCapabilityResponse { return v.Capabilities }).(SKUCapabilityResponseArrayOutput)
}

func (o SkuResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SkuResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Restrictions() RestrictionResponseArrayOutput {
	return o.ApplyT(func(v SkuResponse) []RestrictionResponse { return v.Restrictions }).(RestrictionResponseArrayOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type StorageAccountKeyResponse struct {
	KeyName     string `pulumi:"keyName"`
	Permissions string `pulumi:"permissions"`
	Value       string `pulumi:"value"`
}

type TagPropertyResponse struct {
	ObjectIdentifier string `pulumi:"objectIdentifier"`
	Tag              string `pulumi:"tag"`
	TenantId         string `pulumi:"tenantId"`
	Timestamp        string `pulumi:"timestamp"`
	Upn              string `pulumi:"upn"`
}

type TagPropertyResponseOutput struct{ *pulumi.OutputState }

func (TagPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagPropertyResponse)(nil)).Elem()
}

func (o TagPropertyResponseOutput) ToTagPropertyResponseOutput() TagPropertyResponseOutput {
	return o
}

func (o TagPropertyResponseOutput) ToTagPropertyResponseOutputWithContext(ctx context.Context) TagPropertyResponseOutput {
	return o
}

func (o TagPropertyResponseOutput) ObjectIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.ObjectIdentifier }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.Tag }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) Upn() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.Upn }).(pulumi.StringOutput)
}

type TagPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (TagPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagPropertyResponse)(nil)).Elem()
}

func (o TagPropertyResponseArrayOutput) ToTagPropertyResponseArrayOutput() TagPropertyResponseArrayOutput {
	return o
}

func (o TagPropertyResponseArrayOutput) ToTagPropertyResponseArrayOutputWithContext(ctx context.Context) TagPropertyResponseArrayOutput {
	return o
}

func (o TagPropertyResponseArrayOutput) Index(i pulumi.IntInput) TagPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagPropertyResponse {
		return vs[0].([]TagPropertyResponse)[vs[1].(int)]
	}).(TagPropertyResponseOutput)
}

type UpdateHistoryPropertyResponse struct {
	ImmutabilityPeriodSinceCreationInDays int    `pulumi:"immutabilityPeriodSinceCreationInDays"`
	ObjectIdentifier                      string `pulumi:"objectIdentifier"`
	TenantId                              string `pulumi:"tenantId"`
	Timestamp                             string `pulumi:"timestamp"`
	Update                                string `pulumi:"update"`
	Upn                                   string `pulumi:"upn"`
}

type UpdateHistoryPropertyResponseOutput struct{ *pulumi.OutputState }

func (UpdateHistoryPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateHistoryPropertyResponse)(nil)).Elem()
}

func (o UpdateHistoryPropertyResponseOutput) ToUpdateHistoryPropertyResponseOutput() UpdateHistoryPropertyResponseOutput {
	return o
}

func (o UpdateHistoryPropertyResponseOutput) ToUpdateHistoryPropertyResponseOutputWithContext(ctx context.Context) UpdateHistoryPropertyResponseOutput {
	return o
}

func (o UpdateHistoryPropertyResponseOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) int { return v.ImmutabilityPeriodSinceCreationInDays }).(pulumi.IntOutput)
}

func (o UpdateHistoryPropertyResponseOutput) ObjectIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.ObjectIdentifier }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) Update() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.Update }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) Upn() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.Upn }).(pulumi.StringOutput)
}

type UpdateHistoryPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (UpdateHistoryPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpdateHistoryPropertyResponse)(nil)).Elem()
}

func (o UpdateHistoryPropertyResponseArrayOutput) ToUpdateHistoryPropertyResponseArrayOutput() UpdateHistoryPropertyResponseArrayOutput {
	return o
}

func (o UpdateHistoryPropertyResponseArrayOutput) ToUpdateHistoryPropertyResponseArrayOutputWithContext(ctx context.Context) UpdateHistoryPropertyResponseArrayOutput {
	return o
}

func (o UpdateHistoryPropertyResponseArrayOutput) Index(i pulumi.IntInput) UpdateHistoryPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UpdateHistoryPropertyResponse {
		return vs[0].([]UpdateHistoryPropertyResponse)[vs[1].(int)]
	}).(UpdateHistoryPropertyResponseOutput)
}

type VirtualNetworkRule struct {
	Action                   *Action `pulumi:"action"`
	State                    *State  `pulumi:"state"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRule) Defaults() *VirtualNetworkRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := Action("Allow")
		tmp.Action = &action_
	}
	return &tmp
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Action                   ActionPtrInput     `pulumi:"action"`
	State                    StatePtrInput      `pulumi:"state"`
	VirtualNetworkResourceId pulumi.StringInput `pulumi:"virtualNetworkResourceId"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *Action { return v.Action }).(ActionPtrOutput)
}

func (o VirtualNetworkRuleOutput) State() StatePtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *State { return v.State }).(StatePtrOutput)
}

func (o VirtualNetworkRuleOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRule) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Action                   *string `pulumi:"action"`
	State                    *string `pulumi:"state"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRuleResponse) Defaults() *VirtualNetworkRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
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
	pulumi.RegisterOutputType(CustomDomainOutput{})
	pulumi.RegisterOutputType(CustomDomainPtrOutput{})
	pulumi.RegisterOutputType(CustomDomainResponseOutput{})
	pulumi.RegisterOutputType(DateAfterCreationOutput{})
	pulumi.RegisterOutputType(DateAfterCreationPtrOutput{})
	pulumi.RegisterOutputType(DateAfterCreationResponseOutput{})
	pulumi.RegisterOutputType(DateAfterCreationResponsePtrOutput{})
	pulumi.RegisterOutputType(DateAfterModificationOutput{})
	pulumi.RegisterOutputType(DateAfterModificationPtrOutput{})
	pulumi.RegisterOutputType(DateAfterModificationResponseOutput{})
	pulumi.RegisterOutputType(DateAfterModificationResponsePtrOutput{})
	pulumi.RegisterOutputType(DeleteRetentionPolicyOutput{})
	pulumi.RegisterOutputType(DeleteRetentionPolicyPtrOutput{})
	pulumi.RegisterOutputType(DeleteRetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(DeleteRetentionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServiceOutput{})
	pulumi.RegisterOutputType(EncryptionServicePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServiceResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServiceResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServicesOutput{})
	pulumi.RegisterOutputType(EncryptionServicesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionServicesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServicesResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointsResponseOutput{})
	pulumi.RegisterOutputType(GeoReplicationStatsResponseOutput{})
	pulumi.RegisterOutputType(IPRuleOutput{})
	pulumi.RegisterOutputType(IPRuleArrayOutput{})
	pulumi.RegisterOutputType(IPRuleResponseOutput{})
	pulumi.RegisterOutputType(IPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ImmutabilityPolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LegalHoldPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyActionOutput{})
	pulumi.RegisterOutputType(ManagementPolicyActionResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyBaseBlobOutput{})
	pulumi.RegisterOutputType(ManagementPolicyBaseBlobPtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyBaseBlobResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyBaseBlobResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyDefinitionOutput{})
	pulumi.RegisterOutputType(ManagementPolicyDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyFilterOutput{})
	pulumi.RegisterOutputType(ManagementPolicyFilterPtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyFilterResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyRuleOutput{})
	pulumi.RegisterOutputType(ManagementPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(ManagementPolicyRuleResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementPolicySchemaOutput{})
	pulumi.RegisterOutputType(ManagementPolicySchemaResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotPtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(RestrictionOutput{})
	pulumi.RegisterOutputType(RestrictionArrayOutput{})
	pulumi.RegisterOutputType(RestrictionResponseOutput{})
	pulumi.RegisterOutputType(RestrictionResponseArrayOutput{})
	pulumi.RegisterOutputType(SKUCapabilityResponseOutput{})
	pulumi.RegisterOutputType(SKUCapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(TagPropertyResponseOutput{})
	pulumi.RegisterOutputType(TagPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(UpdateHistoryPropertyResponseOutput{})
	pulumi.RegisterOutputType(UpdateHistoryPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
