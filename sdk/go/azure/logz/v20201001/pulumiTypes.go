


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FilteringTag struct {
	Action *string `pulumi:"action"`
	Name   *string `pulumi:"name"`
	Value  *string `pulumi:"value"`
}





type FilteringTagInput interface {
	pulumi.Input

	ToFilteringTagOutput() FilteringTagOutput
	ToFilteringTagOutputWithContext(context.Context) FilteringTagOutput
}

type FilteringTagArgs struct {
	Action pulumi.StringPtrInput `pulumi:"action"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Value  pulumi.StringPtrInput `pulumi:"value"`
}

func (FilteringTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilteringTag)(nil)).Elem()
}

func (i FilteringTagArgs) ToFilteringTagOutput() FilteringTagOutput {
	return i.ToFilteringTagOutputWithContext(context.Background())
}

func (i FilteringTagArgs) ToFilteringTagOutputWithContext(ctx context.Context) FilteringTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilteringTagOutput)
}





type FilteringTagArrayInput interface {
	pulumi.Input

	ToFilteringTagArrayOutput() FilteringTagArrayOutput
	ToFilteringTagArrayOutputWithContext(context.Context) FilteringTagArrayOutput
}

type FilteringTagArray []FilteringTagInput

func (FilteringTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilteringTag)(nil)).Elem()
}

func (i FilteringTagArray) ToFilteringTagArrayOutput() FilteringTagArrayOutput {
	return i.ToFilteringTagArrayOutputWithContext(context.Background())
}

func (i FilteringTagArray) ToFilteringTagArrayOutputWithContext(ctx context.Context) FilteringTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilteringTagArrayOutput)
}

type FilteringTagOutput struct{ *pulumi.OutputState }

func (FilteringTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilteringTag)(nil)).Elem()
}

func (o FilteringTagOutput) ToFilteringTagOutput() FilteringTagOutput {
	return o
}

func (o FilteringTagOutput) ToFilteringTagOutputWithContext(ctx context.Context) FilteringTagOutput {
	return o
}

func (o FilteringTagOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTag) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o FilteringTagOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTag) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FilteringTagOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTag) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type FilteringTagArrayOutput struct{ *pulumi.OutputState }

func (FilteringTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilteringTag)(nil)).Elem()
}

func (o FilteringTagArrayOutput) ToFilteringTagArrayOutput() FilteringTagArrayOutput {
	return o
}

func (o FilteringTagArrayOutput) ToFilteringTagArrayOutputWithContext(ctx context.Context) FilteringTagArrayOutput {
	return o
}

func (o FilteringTagArrayOutput) Index(i pulumi.IntInput) FilteringTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilteringTag {
		return vs[0].([]FilteringTag)[vs[1].(int)]
	}).(FilteringTagOutput)
}

type FilteringTagResponse struct {
	Action *string `pulumi:"action"`
	Name   *string `pulumi:"name"`
	Value  *string `pulumi:"value"`
}

type FilteringTagResponseOutput struct{ *pulumi.OutputState }

func (FilteringTagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilteringTagResponse)(nil)).Elem()
}

func (o FilteringTagResponseOutput) ToFilteringTagResponseOutput() FilteringTagResponseOutput {
	return o
}

func (o FilteringTagResponseOutput) ToFilteringTagResponseOutputWithContext(ctx context.Context) FilteringTagResponseOutput {
	return o
}

func (o FilteringTagResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTagResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o FilteringTagResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTagResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FilteringTagResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilteringTagResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type FilteringTagResponseArrayOutput struct{ *pulumi.OutputState }

func (FilteringTagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilteringTagResponse)(nil)).Elem()
}

func (o FilteringTagResponseArrayOutput) ToFilteringTagResponseArrayOutput() FilteringTagResponseArrayOutput {
	return o
}

func (o FilteringTagResponseArrayOutput) ToFilteringTagResponseArrayOutputWithContext(ctx context.Context) FilteringTagResponseArrayOutput {
	return o
}

func (o FilteringTagResponseArrayOutput) Index(i pulumi.IntInput) FilteringTagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilteringTagResponse {
		return vs[0].([]FilteringTagResponse)[vs[1].(int)]
	}).(FilteringTagResponseOutput)
}

type IdentityProperties struct {
	Type *string `pulumi:"type"`
}





type IdentityPropertiesInput interface {
	pulumi.Input

	ToIdentityPropertiesOutput() IdentityPropertiesOutput
	ToIdentityPropertiesOutputWithContext(context.Context) IdentityPropertiesOutput
}

type IdentityPropertiesArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return i.ToIdentityPropertiesOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput)
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput).ToIdentityPropertiesPtrOutputWithContext(ctx)
}









type IdentityPropertiesPtrInput interface {
	pulumi.Input

	ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput
	ToIdentityPropertiesPtrOutputWithContext(context.Context) IdentityPropertiesPtrOutput
}

type identityPropertiesPtrType IdentityPropertiesArgs

func IdentityPropertiesPtr(v *IdentityPropertiesArgs) IdentityPropertiesPtrInput {
	return (*identityPropertiesPtrType)(v)
}

func (*identityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesPtrOutput)
}

type IdentityPropertiesOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProperties) *IdentityProperties {
		return &v
	}).(IdentityPropertiesPtrOutput)
}

func (o IdentityPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) Elem() IdentityPropertiesOutput {
	return o.ApplyT(func(v *IdentityProperties) IdentityProperties {
		if v != nil {
			return *v
		}
		var ret IdentityProperties
		return ret
	}).(IdentityPropertiesOutput)
}

func (o IdentityPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityPropertiesResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityPropertiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) Elem() IdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) IdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret IdentityPropertiesResponse
		return ret
	}).(IdentityPropertiesResponseOutput)
}

func (o IdentityPropertiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type LogRules struct {
	FilteringTags        []FilteringTag `pulumi:"filteringTags"`
	SendAadLogs          *bool          `pulumi:"sendAadLogs"`
	SendActivityLogs     *bool          `pulumi:"sendActivityLogs"`
	SendSubscriptionLogs *bool          `pulumi:"sendSubscriptionLogs"`
}





type LogRulesInput interface {
	pulumi.Input

	ToLogRulesOutput() LogRulesOutput
	ToLogRulesOutputWithContext(context.Context) LogRulesOutput
}

type LogRulesArgs struct {
	FilteringTags        FilteringTagArrayInput `pulumi:"filteringTags"`
	SendAadLogs          pulumi.BoolPtrInput    `pulumi:"sendAadLogs"`
	SendActivityLogs     pulumi.BoolPtrInput    `pulumi:"sendActivityLogs"`
	SendSubscriptionLogs pulumi.BoolPtrInput    `pulumi:"sendSubscriptionLogs"`
}

func (LogRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogRules)(nil)).Elem()
}

func (i LogRulesArgs) ToLogRulesOutput() LogRulesOutput {
	return i.ToLogRulesOutputWithContext(context.Background())
}

func (i LogRulesArgs) ToLogRulesOutputWithContext(ctx context.Context) LogRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogRulesOutput)
}

func (i LogRulesArgs) ToLogRulesPtrOutput() LogRulesPtrOutput {
	return i.ToLogRulesPtrOutputWithContext(context.Background())
}

func (i LogRulesArgs) ToLogRulesPtrOutputWithContext(ctx context.Context) LogRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogRulesOutput).ToLogRulesPtrOutputWithContext(ctx)
}









type LogRulesPtrInput interface {
	pulumi.Input

	ToLogRulesPtrOutput() LogRulesPtrOutput
	ToLogRulesPtrOutputWithContext(context.Context) LogRulesPtrOutput
}

type logRulesPtrType LogRulesArgs

func LogRulesPtr(v *LogRulesArgs) LogRulesPtrInput {
	return (*logRulesPtrType)(v)
}

func (*logRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogRules)(nil)).Elem()
}

func (i *logRulesPtrType) ToLogRulesPtrOutput() LogRulesPtrOutput {
	return i.ToLogRulesPtrOutputWithContext(context.Background())
}

func (i *logRulesPtrType) ToLogRulesPtrOutputWithContext(ctx context.Context) LogRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogRulesPtrOutput)
}

type LogRulesOutput struct{ *pulumi.OutputState }

func (LogRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogRules)(nil)).Elem()
}

func (o LogRulesOutput) ToLogRulesOutput() LogRulesOutput {
	return o
}

func (o LogRulesOutput) ToLogRulesOutputWithContext(ctx context.Context) LogRulesOutput {
	return o
}

func (o LogRulesOutput) ToLogRulesPtrOutput() LogRulesPtrOutput {
	return o.ToLogRulesPtrOutputWithContext(context.Background())
}

func (o LogRulesOutput) ToLogRulesPtrOutputWithContext(ctx context.Context) LogRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogRules) *LogRules {
		return &v
	}).(LogRulesPtrOutput)
}

func (o LogRulesOutput) FilteringTags() FilteringTagArrayOutput {
	return o.ApplyT(func(v LogRules) []FilteringTag { return v.FilteringTags }).(FilteringTagArrayOutput)
}

func (o LogRulesOutput) SendAadLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRules) *bool { return v.SendAadLogs }).(pulumi.BoolPtrOutput)
}

func (o LogRulesOutput) SendActivityLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRules) *bool { return v.SendActivityLogs }).(pulumi.BoolPtrOutput)
}

func (o LogRulesOutput) SendSubscriptionLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRules) *bool { return v.SendSubscriptionLogs }).(pulumi.BoolPtrOutput)
}

type LogRulesPtrOutput struct{ *pulumi.OutputState }

func (LogRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogRules)(nil)).Elem()
}

func (o LogRulesPtrOutput) ToLogRulesPtrOutput() LogRulesPtrOutput {
	return o
}

func (o LogRulesPtrOutput) ToLogRulesPtrOutputWithContext(ctx context.Context) LogRulesPtrOutput {
	return o
}

func (o LogRulesPtrOutput) Elem() LogRulesOutput {
	return o.ApplyT(func(v *LogRules) LogRules {
		if v != nil {
			return *v
		}
		var ret LogRules
		return ret
	}).(LogRulesOutput)
}

func (o LogRulesPtrOutput) FilteringTags() FilteringTagArrayOutput {
	return o.ApplyT(func(v *LogRules) []FilteringTag {
		if v == nil {
			return nil
		}
		return v.FilteringTags
	}).(FilteringTagArrayOutput)
}

func (o LogRulesPtrOutput) SendAadLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRules) *bool {
		if v == nil {
			return nil
		}
		return v.SendAadLogs
	}).(pulumi.BoolPtrOutput)
}

func (o LogRulesPtrOutput) SendActivityLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRules) *bool {
		if v == nil {
			return nil
		}
		return v.SendActivityLogs
	}).(pulumi.BoolPtrOutput)
}

func (o LogRulesPtrOutput) SendSubscriptionLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRules) *bool {
		if v == nil {
			return nil
		}
		return v.SendSubscriptionLogs
	}).(pulumi.BoolPtrOutput)
}

type LogRulesResponse struct {
	FilteringTags        []FilteringTagResponse `pulumi:"filteringTags"`
	SendAadLogs          *bool                  `pulumi:"sendAadLogs"`
	SendActivityLogs     *bool                  `pulumi:"sendActivityLogs"`
	SendSubscriptionLogs *bool                  `pulumi:"sendSubscriptionLogs"`
}

type LogRulesResponseOutput struct{ *pulumi.OutputState }

func (LogRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogRulesResponse)(nil)).Elem()
}

func (o LogRulesResponseOutput) ToLogRulesResponseOutput() LogRulesResponseOutput {
	return o
}

func (o LogRulesResponseOutput) ToLogRulesResponseOutputWithContext(ctx context.Context) LogRulesResponseOutput {
	return o
}

func (o LogRulesResponseOutput) FilteringTags() FilteringTagResponseArrayOutput {
	return o.ApplyT(func(v LogRulesResponse) []FilteringTagResponse { return v.FilteringTags }).(FilteringTagResponseArrayOutput)
}

func (o LogRulesResponseOutput) SendAadLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRulesResponse) *bool { return v.SendAadLogs }).(pulumi.BoolPtrOutput)
}

func (o LogRulesResponseOutput) SendActivityLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRulesResponse) *bool { return v.SendActivityLogs }).(pulumi.BoolPtrOutput)
}

func (o LogRulesResponseOutput) SendSubscriptionLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogRulesResponse) *bool { return v.SendSubscriptionLogs }).(pulumi.BoolPtrOutput)
}

type LogRulesResponsePtrOutput struct{ *pulumi.OutputState }

func (LogRulesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogRulesResponse)(nil)).Elem()
}

func (o LogRulesResponsePtrOutput) ToLogRulesResponsePtrOutput() LogRulesResponsePtrOutput {
	return o
}

func (o LogRulesResponsePtrOutput) ToLogRulesResponsePtrOutputWithContext(ctx context.Context) LogRulesResponsePtrOutput {
	return o
}

func (o LogRulesResponsePtrOutput) Elem() LogRulesResponseOutput {
	return o.ApplyT(func(v *LogRulesResponse) LogRulesResponse {
		if v != nil {
			return *v
		}
		var ret LogRulesResponse
		return ret
	}).(LogRulesResponseOutput)
}

func (o LogRulesResponsePtrOutput) FilteringTags() FilteringTagResponseArrayOutput {
	return o.ApplyT(func(v *LogRulesResponse) []FilteringTagResponse {
		if v == nil {
			return nil
		}
		return v.FilteringTags
	}).(FilteringTagResponseArrayOutput)
}

func (o LogRulesResponsePtrOutput) SendAadLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRulesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendAadLogs
	}).(pulumi.BoolPtrOutput)
}

func (o LogRulesResponsePtrOutput) SendActivityLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRulesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendActivityLogs
	}).(pulumi.BoolPtrOutput)
}

func (o LogRulesResponsePtrOutput) SendSubscriptionLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogRulesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendSubscriptionLogs
	}).(pulumi.BoolPtrOutput)
}

type LogzOrganizationProperties struct {
	CompanyName     *string `pulumi:"companyName"`
	EnterpriseAppId *string `pulumi:"enterpriseAppId"`
	SingleSignOnUrl *string `pulumi:"singleSignOnUrl"`
}





type LogzOrganizationPropertiesInput interface {
	pulumi.Input

	ToLogzOrganizationPropertiesOutput() LogzOrganizationPropertiesOutput
	ToLogzOrganizationPropertiesOutputWithContext(context.Context) LogzOrganizationPropertiesOutput
}

type LogzOrganizationPropertiesArgs struct {
	CompanyName     pulumi.StringPtrInput `pulumi:"companyName"`
	EnterpriseAppId pulumi.StringPtrInput `pulumi:"enterpriseAppId"`
	SingleSignOnUrl pulumi.StringPtrInput `pulumi:"singleSignOnUrl"`
}

func (LogzOrganizationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogzOrganizationProperties)(nil)).Elem()
}

func (i LogzOrganizationPropertiesArgs) ToLogzOrganizationPropertiesOutput() LogzOrganizationPropertiesOutput {
	return i.ToLogzOrganizationPropertiesOutputWithContext(context.Background())
}

func (i LogzOrganizationPropertiesArgs) ToLogzOrganizationPropertiesOutputWithContext(ctx context.Context) LogzOrganizationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogzOrganizationPropertiesOutput)
}

func (i LogzOrganizationPropertiesArgs) ToLogzOrganizationPropertiesPtrOutput() LogzOrganizationPropertiesPtrOutput {
	return i.ToLogzOrganizationPropertiesPtrOutputWithContext(context.Background())
}

func (i LogzOrganizationPropertiesArgs) ToLogzOrganizationPropertiesPtrOutputWithContext(ctx context.Context) LogzOrganizationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogzOrganizationPropertiesOutput).ToLogzOrganizationPropertiesPtrOutputWithContext(ctx)
}









type LogzOrganizationPropertiesPtrInput interface {
	pulumi.Input

	ToLogzOrganizationPropertiesPtrOutput() LogzOrganizationPropertiesPtrOutput
	ToLogzOrganizationPropertiesPtrOutputWithContext(context.Context) LogzOrganizationPropertiesPtrOutput
}

type logzOrganizationPropertiesPtrType LogzOrganizationPropertiesArgs

func LogzOrganizationPropertiesPtr(v *LogzOrganizationPropertiesArgs) LogzOrganizationPropertiesPtrInput {
	return (*logzOrganizationPropertiesPtrType)(v)
}

func (*logzOrganizationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogzOrganizationProperties)(nil)).Elem()
}

func (i *logzOrganizationPropertiesPtrType) ToLogzOrganizationPropertiesPtrOutput() LogzOrganizationPropertiesPtrOutput {
	return i.ToLogzOrganizationPropertiesPtrOutputWithContext(context.Background())
}

func (i *logzOrganizationPropertiesPtrType) ToLogzOrganizationPropertiesPtrOutputWithContext(ctx context.Context) LogzOrganizationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogzOrganizationPropertiesPtrOutput)
}

type LogzOrganizationPropertiesOutput struct{ *pulumi.OutputState }

func (LogzOrganizationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogzOrganizationProperties)(nil)).Elem()
}

func (o LogzOrganizationPropertiesOutput) ToLogzOrganizationPropertiesOutput() LogzOrganizationPropertiesOutput {
	return o
}

func (o LogzOrganizationPropertiesOutput) ToLogzOrganizationPropertiesOutputWithContext(ctx context.Context) LogzOrganizationPropertiesOutput {
	return o
}

func (o LogzOrganizationPropertiesOutput) ToLogzOrganizationPropertiesPtrOutput() LogzOrganizationPropertiesPtrOutput {
	return o.ToLogzOrganizationPropertiesPtrOutputWithContext(context.Background())
}

func (o LogzOrganizationPropertiesOutput) ToLogzOrganizationPropertiesPtrOutputWithContext(ctx context.Context) LogzOrganizationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogzOrganizationProperties) *LogzOrganizationProperties {
		return &v
	}).(LogzOrganizationPropertiesPtrOutput)
}

func (o LogzOrganizationPropertiesOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogzOrganizationProperties) *string { return v.CompanyName }).(pulumi.StringPtrOutput)
}

func (o LogzOrganizationPropertiesOutput) EnterpriseAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogzOrganizationProperties) *string { return v.EnterpriseAppId }).(pulumi.StringPtrOutput)
}

func (o LogzOrganizationPropertiesOutput) SingleSignOnUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogzOrganizationProperties) *string { return v.SingleSignOnUrl }).(pulumi.StringPtrOutput)
}

type LogzOrganizationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (LogzOrganizationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogzOrganizationProperties)(nil)).Elem()
}

func (o LogzOrganizationPropertiesPtrOutput) ToLogzOrganizationPropertiesPtrOutput() LogzOrganizationPropertiesPtrOutput {
	return o
}

func (o LogzOrganizationPropertiesPtrOutput) ToLogzOrganizationPropertiesPtrOutputWithContext(ctx context.Context) LogzOrganizationPropertiesPtrOutput {
	return o
}

func (o LogzOrganizationPropertiesPtrOutput) Elem() LogzOrganizationPropertiesOutput {
	return o.ApplyT(func(v *LogzOrganizationProperties) LogzOrganizationProperties {
		if v != nil {
			return *v
		}
		var ret LogzOrganizationProperties
		return ret
	}).(LogzOrganizationPropertiesOutput)
}

func (o LogzOrganizationPropertiesPtrOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogzOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.CompanyName
	}).(pulumi.StringPtrOutput)
}

func (o LogzOrganizationPropertiesPtrOutput) EnterpriseAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogzOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.EnterpriseAppId
	}).(pulumi.StringPtrOutput)
}

func (o LogzOrganizationPropertiesPtrOutput) SingleSignOnUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogzOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.SingleSignOnUrl
	}).(pulumi.StringPtrOutput)
}

type LogzOrganizationPropertiesResponse struct {
	CompanyName     *string `pulumi:"companyName"`
	EnterpriseAppId *string `pulumi:"enterpriseAppId"`
	Id              string  `pulumi:"id"`
	SingleSignOnUrl *string `pulumi:"singleSignOnUrl"`
}

type LogzOrganizationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LogzOrganizationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogzOrganizationPropertiesResponse)(nil)).Elem()
}

func (o LogzOrganizationPropertiesResponseOutput) ToLogzOrganizationPropertiesResponseOutput() LogzOrganizationPropertiesResponseOutput {
	return o
}

func (o LogzOrganizationPropertiesResponseOutput) ToLogzOrganizationPropertiesResponseOutputWithContext(ctx context.Context) LogzOrganizationPropertiesResponseOutput {
	return o
}

func (o LogzOrganizationPropertiesResponseOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogzOrganizationPropertiesResponse) *string { return v.CompanyName }).(pulumi.StringPtrOutput)
}

func (o LogzOrganizationPropertiesResponseOutput) EnterpriseAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogzOrganizationPropertiesResponse) *string { return v.EnterpriseAppId }).(pulumi.StringPtrOutput)
}

func (o LogzOrganizationPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LogzOrganizationPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o LogzOrganizationPropertiesResponseOutput) SingleSignOnUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogzOrganizationPropertiesResponse) *string { return v.SingleSignOnUrl }).(pulumi.StringPtrOutput)
}

type LogzOrganizationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LogzOrganizationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogzOrganizationPropertiesResponse)(nil)).Elem()
}

func (o LogzOrganizationPropertiesResponsePtrOutput) ToLogzOrganizationPropertiesResponsePtrOutput() LogzOrganizationPropertiesResponsePtrOutput {
	return o
}

func (o LogzOrganizationPropertiesResponsePtrOutput) ToLogzOrganizationPropertiesResponsePtrOutputWithContext(ctx context.Context) LogzOrganizationPropertiesResponsePtrOutput {
	return o
}

func (o LogzOrganizationPropertiesResponsePtrOutput) Elem() LogzOrganizationPropertiesResponseOutput {
	return o.ApplyT(func(v *LogzOrganizationPropertiesResponse) LogzOrganizationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LogzOrganizationPropertiesResponse
		return ret
	}).(LogzOrganizationPropertiesResponseOutput)
}

func (o LogzOrganizationPropertiesResponsePtrOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogzOrganizationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CompanyName
	}).(pulumi.StringPtrOutput)
}

func (o LogzOrganizationPropertiesResponsePtrOutput) EnterpriseAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogzOrganizationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.EnterpriseAppId
	}).(pulumi.StringPtrOutput)
}

func (o LogzOrganizationPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogzOrganizationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o LogzOrganizationPropertiesResponsePtrOutput) SingleSignOnUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogzOrganizationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SingleSignOnUrl
	}).(pulumi.StringPtrOutput)
}

type MonitorProperties struct {
	LogzOrganizationProperties    *LogzOrganizationProperties `pulumi:"logzOrganizationProperties"`
	MarketplaceSubscriptionStatus *string                     `pulumi:"marketplaceSubscriptionStatus"`
	MonitoringStatus              *string                     `pulumi:"monitoringStatus"`
	PlanData                      *PlanData                   `pulumi:"planData"`
	UserInfo                      *UserInfo                   `pulumi:"userInfo"`
}





type MonitorPropertiesInput interface {
	pulumi.Input

	ToMonitorPropertiesOutput() MonitorPropertiesOutput
	ToMonitorPropertiesOutputWithContext(context.Context) MonitorPropertiesOutput
}

type MonitorPropertiesArgs struct {
	LogzOrganizationProperties    LogzOrganizationPropertiesPtrInput `pulumi:"logzOrganizationProperties"`
	MarketplaceSubscriptionStatus pulumi.StringPtrInput              `pulumi:"marketplaceSubscriptionStatus"`
	MonitoringStatus              pulumi.StringPtrInput              `pulumi:"monitoringStatus"`
	PlanData                      PlanDataPtrInput                   `pulumi:"planData"`
	UserInfo                      UserInfoPtrInput                   `pulumi:"userInfo"`
}

func (MonitorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProperties)(nil)).Elem()
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesOutput() MonitorPropertiesOutput {
	return i.ToMonitorPropertiesOutputWithContext(context.Background())
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesOutputWithContext(ctx context.Context) MonitorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesOutput)
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return i.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesOutput).ToMonitorPropertiesPtrOutputWithContext(ctx)
}









type MonitorPropertiesPtrInput interface {
	pulumi.Input

	ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput
	ToMonitorPropertiesPtrOutputWithContext(context.Context) MonitorPropertiesPtrOutput
}

type monitorPropertiesPtrType MonitorPropertiesArgs

func MonitorPropertiesPtr(v *MonitorPropertiesArgs) MonitorPropertiesPtrInput {
	return (*monitorPropertiesPtrType)(v)
}

func (*monitorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorProperties)(nil)).Elem()
}

func (i *monitorPropertiesPtrType) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return i.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (i *monitorPropertiesPtrType) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesPtrOutput)
}

type MonitorPropertiesOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProperties)(nil)).Elem()
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesOutput() MonitorPropertiesOutput {
	return o
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesOutputWithContext(ctx context.Context) MonitorPropertiesOutput {
	return o
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return o.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorProperties) *MonitorProperties {
		return &v
	}).(MonitorPropertiesPtrOutput)
}

func (o MonitorPropertiesOutput) LogzOrganizationProperties() LogzOrganizationPropertiesPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *LogzOrganizationProperties { return v.LogzOrganizationProperties }).(LogzOrganizationPropertiesPtrOutput)
}

func (o MonitorPropertiesOutput) MarketplaceSubscriptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *string { return v.MarketplaceSubscriptionStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *string { return v.MonitoringStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesOutput) PlanData() PlanDataPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *PlanData { return v.PlanData }).(PlanDataPtrOutput)
}

func (o MonitorPropertiesOutput) UserInfo() UserInfoPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *UserInfo { return v.UserInfo }).(UserInfoPtrOutput)
}

type MonitorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorProperties)(nil)).Elem()
}

func (o MonitorPropertiesPtrOutput) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return o
}

func (o MonitorPropertiesPtrOutput) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return o
}

func (o MonitorPropertiesPtrOutput) Elem() MonitorPropertiesOutput {
	return o.ApplyT(func(v *MonitorProperties) MonitorProperties {
		if v != nil {
			return *v
		}
		var ret MonitorProperties
		return ret
	}).(MonitorPropertiesOutput)
}

func (o MonitorPropertiesPtrOutput) LogzOrganizationProperties() LogzOrganizationPropertiesPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *LogzOrganizationProperties {
		if v == nil {
			return nil
		}
		return v.LogzOrganizationProperties
	}).(LogzOrganizationPropertiesPtrOutput)
}

func (o MonitorPropertiesPtrOutput) MarketplaceSubscriptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *string {
		if v == nil {
			return nil
		}
		return v.MarketplaceSubscriptionStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesPtrOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *string {
		if v == nil {
			return nil
		}
		return v.MonitoringStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesPtrOutput) PlanData() PlanDataPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *PlanData {
		if v == nil {
			return nil
		}
		return v.PlanData
	}).(PlanDataPtrOutput)
}

func (o MonitorPropertiesPtrOutput) UserInfo() UserInfoPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *UserInfo {
		if v == nil {
			return nil
		}
		return v.UserInfo
	}).(UserInfoPtrOutput)
}

type MonitorPropertiesResponse struct {
	LiftrResourceCategory         string                              `pulumi:"liftrResourceCategory"`
	LiftrResourcePreference       int                                 `pulumi:"liftrResourcePreference"`
	LogzOrganizationProperties    *LogzOrganizationPropertiesResponse `pulumi:"logzOrganizationProperties"`
	MarketplaceSubscriptionStatus *string                             `pulumi:"marketplaceSubscriptionStatus"`
	MonitoringStatus              *string                             `pulumi:"monitoringStatus"`
	PlanData                      *PlanDataResponse                   `pulumi:"planData"`
	ProvisioningState             string                              `pulumi:"provisioningState"`
	UserInfo                      *UserInfoResponse                   `pulumi:"userInfo"`
}

type MonitorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorPropertiesResponse)(nil)).Elem()
}

func (o MonitorPropertiesResponseOutput) ToMonitorPropertiesResponseOutput() MonitorPropertiesResponseOutput {
	return o
}

func (o MonitorPropertiesResponseOutput) ToMonitorPropertiesResponseOutputWithContext(ctx context.Context) MonitorPropertiesResponseOutput {
	return o
}

func (o MonitorPropertiesResponseOutput) LiftrResourceCategory() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) string { return v.LiftrResourceCategory }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseOutput) LiftrResourcePreference() pulumi.IntOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) int { return v.LiftrResourcePreference }).(pulumi.IntOutput)
}

func (o MonitorPropertiesResponseOutput) LogzOrganizationProperties() LogzOrganizationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *LogzOrganizationPropertiesResponse {
		return v.LogzOrganizationProperties
	}).(LogzOrganizationPropertiesResponsePtrOutput)
}

func (o MonitorPropertiesResponseOutput) MarketplaceSubscriptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *string { return v.MarketplaceSubscriptionStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesResponseOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *string { return v.MonitoringStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesResponseOutput) PlanData() PlanDataResponsePtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *PlanDataResponse { return v.PlanData }).(PlanDataResponsePtrOutput)
}

func (o MonitorPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseOutput) UserInfo() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *UserInfoResponse { return v.UserInfo }).(UserInfoResponsePtrOutput)
}

type MonitoredResourceResponse struct {
	Id                     *string            `pulumi:"id"`
	ReasonForLogsStatus    *string            `pulumi:"reasonForLogsStatus"`
	ReasonForMetricsStatus *string            `pulumi:"reasonForMetricsStatus"`
	SendingLogs            *bool              `pulumi:"sendingLogs"`
	SendingMetrics         *bool              `pulumi:"sendingMetrics"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
}

type MonitoredResourceResponseOutput struct{ *pulumi.OutputState }

func (MonitoredResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoredResourceResponse)(nil)).Elem()
}

func (o MonitoredResourceResponseOutput) ToMonitoredResourceResponseOutput() MonitoredResourceResponseOutput {
	return o
}

func (o MonitoredResourceResponseOutput) ToMonitoredResourceResponseOutputWithContext(ctx context.Context) MonitoredResourceResponseOutput {
	return o
}

func (o MonitoredResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o MonitoredResourceResponseOutput) ReasonForLogsStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) *string { return v.ReasonForLogsStatus }).(pulumi.StringPtrOutput)
}

func (o MonitoredResourceResponseOutput) ReasonForMetricsStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) *string { return v.ReasonForMetricsStatus }).(pulumi.StringPtrOutput)
}

func (o MonitoredResourceResponseOutput) SendingLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) *bool { return v.SendingLogs }).(pulumi.BoolPtrOutput)
}

func (o MonitoredResourceResponseOutput) SendingMetrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) *bool { return v.SendingMetrics }).(pulumi.BoolPtrOutput)
}

func (o MonitoredResourceResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

type MonitoredResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (MonitoredResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitoredResourceResponse)(nil)).Elem()
}

func (o MonitoredResourceResponseArrayOutput) ToMonitoredResourceResponseArrayOutput() MonitoredResourceResponseArrayOutput {
	return o
}

func (o MonitoredResourceResponseArrayOutput) ToMonitoredResourceResponseArrayOutputWithContext(ctx context.Context) MonitoredResourceResponseArrayOutput {
	return o
}

func (o MonitoredResourceResponseArrayOutput) Index(i pulumi.IntInput) MonitoredResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitoredResourceResponse {
		return vs[0].([]MonitoredResourceResponse)[vs[1].(int)]
	}).(MonitoredResourceResponseOutput)
}

type MonitoringTagRulesProperties struct {
	LogRules *LogRules `pulumi:"logRules"`
}





type MonitoringTagRulesPropertiesInput interface {
	pulumi.Input

	ToMonitoringTagRulesPropertiesOutput() MonitoringTagRulesPropertiesOutput
	ToMonitoringTagRulesPropertiesOutputWithContext(context.Context) MonitoringTagRulesPropertiesOutput
}

type MonitoringTagRulesPropertiesArgs struct {
	LogRules LogRulesPtrInput `pulumi:"logRules"`
}

func (MonitoringTagRulesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringTagRulesProperties)(nil)).Elem()
}

func (i MonitoringTagRulesPropertiesArgs) ToMonitoringTagRulesPropertiesOutput() MonitoringTagRulesPropertiesOutput {
	return i.ToMonitoringTagRulesPropertiesOutputWithContext(context.Background())
}

func (i MonitoringTagRulesPropertiesArgs) ToMonitoringTagRulesPropertiesOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringTagRulesPropertiesOutput)
}

func (i MonitoringTagRulesPropertiesArgs) ToMonitoringTagRulesPropertiesPtrOutput() MonitoringTagRulesPropertiesPtrOutput {
	return i.ToMonitoringTagRulesPropertiesPtrOutputWithContext(context.Background())
}

func (i MonitoringTagRulesPropertiesArgs) ToMonitoringTagRulesPropertiesPtrOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringTagRulesPropertiesOutput).ToMonitoringTagRulesPropertiesPtrOutputWithContext(ctx)
}









type MonitoringTagRulesPropertiesPtrInput interface {
	pulumi.Input

	ToMonitoringTagRulesPropertiesPtrOutput() MonitoringTagRulesPropertiesPtrOutput
	ToMonitoringTagRulesPropertiesPtrOutputWithContext(context.Context) MonitoringTagRulesPropertiesPtrOutput
}

type monitoringTagRulesPropertiesPtrType MonitoringTagRulesPropertiesArgs

func MonitoringTagRulesPropertiesPtr(v *MonitoringTagRulesPropertiesArgs) MonitoringTagRulesPropertiesPtrInput {
	return (*monitoringTagRulesPropertiesPtrType)(v)
}

func (*monitoringTagRulesPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringTagRulesProperties)(nil)).Elem()
}

func (i *monitoringTagRulesPropertiesPtrType) ToMonitoringTagRulesPropertiesPtrOutput() MonitoringTagRulesPropertiesPtrOutput {
	return i.ToMonitoringTagRulesPropertiesPtrOutputWithContext(context.Background())
}

func (i *monitoringTagRulesPropertiesPtrType) ToMonitoringTagRulesPropertiesPtrOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringTagRulesPropertiesPtrOutput)
}

type MonitoringTagRulesPropertiesOutput struct{ *pulumi.OutputState }

func (MonitoringTagRulesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringTagRulesProperties)(nil)).Elem()
}

func (o MonitoringTagRulesPropertiesOutput) ToMonitoringTagRulesPropertiesOutput() MonitoringTagRulesPropertiesOutput {
	return o
}

func (o MonitoringTagRulesPropertiesOutput) ToMonitoringTagRulesPropertiesOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesOutput {
	return o
}

func (o MonitoringTagRulesPropertiesOutput) ToMonitoringTagRulesPropertiesPtrOutput() MonitoringTagRulesPropertiesPtrOutput {
	return o.ToMonitoringTagRulesPropertiesPtrOutputWithContext(context.Background())
}

func (o MonitoringTagRulesPropertiesOutput) ToMonitoringTagRulesPropertiesPtrOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitoringTagRulesProperties) *MonitoringTagRulesProperties {
		return &v
	}).(MonitoringTagRulesPropertiesPtrOutput)
}

func (o MonitoringTagRulesPropertiesOutput) LogRules() LogRulesPtrOutput {
	return o.ApplyT(func(v MonitoringTagRulesProperties) *LogRules { return v.LogRules }).(LogRulesPtrOutput)
}

type MonitoringTagRulesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MonitoringTagRulesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringTagRulesProperties)(nil)).Elem()
}

func (o MonitoringTagRulesPropertiesPtrOutput) ToMonitoringTagRulesPropertiesPtrOutput() MonitoringTagRulesPropertiesPtrOutput {
	return o
}

func (o MonitoringTagRulesPropertiesPtrOutput) ToMonitoringTagRulesPropertiesPtrOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesPtrOutput {
	return o
}

func (o MonitoringTagRulesPropertiesPtrOutput) Elem() MonitoringTagRulesPropertiesOutput {
	return o.ApplyT(func(v *MonitoringTagRulesProperties) MonitoringTagRulesProperties {
		if v != nil {
			return *v
		}
		var ret MonitoringTagRulesProperties
		return ret
	}).(MonitoringTagRulesPropertiesOutput)
}

func (o MonitoringTagRulesPropertiesPtrOutput) LogRules() LogRulesPtrOutput {
	return o.ApplyT(func(v *MonitoringTagRulesProperties) *LogRules {
		if v == nil {
			return nil
		}
		return v.LogRules
	}).(LogRulesPtrOutput)
}

type MonitoringTagRulesPropertiesResponse struct {
	LogRules          *LogRulesResponse  `pulumi:"logRules"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
}

type MonitoringTagRulesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MonitoringTagRulesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringTagRulesPropertiesResponse)(nil)).Elem()
}

func (o MonitoringTagRulesPropertiesResponseOutput) ToMonitoringTagRulesPropertiesResponseOutput() MonitoringTagRulesPropertiesResponseOutput {
	return o
}

func (o MonitoringTagRulesPropertiesResponseOutput) ToMonitoringTagRulesPropertiesResponseOutputWithContext(ctx context.Context) MonitoringTagRulesPropertiesResponseOutput {
	return o
}

func (o MonitoringTagRulesPropertiesResponseOutput) LogRules() LogRulesResponsePtrOutput {
	return o.ApplyT(func(v MonitoringTagRulesPropertiesResponse) *LogRulesResponse { return v.LogRules }).(LogRulesResponsePtrOutput)
}

func (o MonitoringTagRulesPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MonitoringTagRulesPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MonitoringTagRulesPropertiesResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v MonitoringTagRulesPropertiesResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

type PlanData struct {
	BillingCycle  *string `pulumi:"billingCycle"`
	EffectiveDate *string `pulumi:"effectiveDate"`
	PlanDetails   *string `pulumi:"planDetails"`
	UsageType     *string `pulumi:"usageType"`
}





type PlanDataInput interface {
	pulumi.Input

	ToPlanDataOutput() PlanDataOutput
	ToPlanDataOutputWithContext(context.Context) PlanDataOutput
}

type PlanDataArgs struct {
	BillingCycle  pulumi.StringPtrInput `pulumi:"billingCycle"`
	EffectiveDate pulumi.StringPtrInput `pulumi:"effectiveDate"`
	PlanDetails   pulumi.StringPtrInput `pulumi:"planDetails"`
	UsageType     pulumi.StringPtrInput `pulumi:"usageType"`
}

func (PlanDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanData)(nil)).Elem()
}

func (i PlanDataArgs) ToPlanDataOutput() PlanDataOutput {
	return i.ToPlanDataOutputWithContext(context.Background())
}

func (i PlanDataArgs) ToPlanDataOutputWithContext(ctx context.Context) PlanDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanDataOutput)
}

func (i PlanDataArgs) ToPlanDataPtrOutput() PlanDataPtrOutput {
	return i.ToPlanDataPtrOutputWithContext(context.Background())
}

func (i PlanDataArgs) ToPlanDataPtrOutputWithContext(ctx context.Context) PlanDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanDataOutput).ToPlanDataPtrOutputWithContext(ctx)
}









type PlanDataPtrInput interface {
	pulumi.Input

	ToPlanDataPtrOutput() PlanDataPtrOutput
	ToPlanDataPtrOutputWithContext(context.Context) PlanDataPtrOutput
}

type planDataPtrType PlanDataArgs

func PlanDataPtr(v *PlanDataArgs) PlanDataPtrInput {
	return (*planDataPtrType)(v)
}

func (*planDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanData)(nil)).Elem()
}

func (i *planDataPtrType) ToPlanDataPtrOutput() PlanDataPtrOutput {
	return i.ToPlanDataPtrOutputWithContext(context.Background())
}

func (i *planDataPtrType) ToPlanDataPtrOutputWithContext(ctx context.Context) PlanDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanDataPtrOutput)
}

type PlanDataOutput struct{ *pulumi.OutputState }

func (PlanDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanData)(nil)).Elem()
}

func (o PlanDataOutput) ToPlanDataOutput() PlanDataOutput {
	return o
}

func (o PlanDataOutput) ToPlanDataOutputWithContext(ctx context.Context) PlanDataOutput {
	return o
}

func (o PlanDataOutput) ToPlanDataPtrOutput() PlanDataPtrOutput {
	return o.ToPlanDataPtrOutputWithContext(context.Background())
}

func (o PlanDataOutput) ToPlanDataPtrOutputWithContext(ctx context.Context) PlanDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlanData) *PlanData {
		return &v
	}).(PlanDataPtrOutput)
}

func (o PlanDataOutput) BillingCycle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanData) *string { return v.BillingCycle }).(pulumi.StringPtrOutput)
}

func (o PlanDataOutput) EffectiveDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanData) *string { return v.EffectiveDate }).(pulumi.StringPtrOutput)
}

func (o PlanDataOutput) PlanDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanData) *string { return v.PlanDetails }).(pulumi.StringPtrOutput)
}

func (o PlanDataOutput) UsageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanData) *string { return v.UsageType }).(pulumi.StringPtrOutput)
}

type PlanDataPtrOutput struct{ *pulumi.OutputState }

func (PlanDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanData)(nil)).Elem()
}

func (o PlanDataPtrOutput) ToPlanDataPtrOutput() PlanDataPtrOutput {
	return o
}

func (o PlanDataPtrOutput) ToPlanDataPtrOutputWithContext(ctx context.Context) PlanDataPtrOutput {
	return o
}

func (o PlanDataPtrOutput) Elem() PlanDataOutput {
	return o.ApplyT(func(v *PlanData) PlanData {
		if v != nil {
			return *v
		}
		var ret PlanData
		return ret
	}).(PlanDataOutput)
}

func (o PlanDataPtrOutput) BillingCycle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanData) *string {
		if v == nil {
			return nil
		}
		return v.BillingCycle
	}).(pulumi.StringPtrOutput)
}

func (o PlanDataPtrOutput) EffectiveDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanData) *string {
		if v == nil {
			return nil
		}
		return v.EffectiveDate
	}).(pulumi.StringPtrOutput)
}

func (o PlanDataPtrOutput) PlanDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanData) *string {
		if v == nil {
			return nil
		}
		return v.PlanDetails
	}).(pulumi.StringPtrOutput)
}

func (o PlanDataPtrOutput) UsageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanData) *string {
		if v == nil {
			return nil
		}
		return v.UsageType
	}).(pulumi.StringPtrOutput)
}

type PlanDataResponse struct {
	BillingCycle  *string `pulumi:"billingCycle"`
	EffectiveDate *string `pulumi:"effectiveDate"`
	PlanDetails   *string `pulumi:"planDetails"`
	UsageType     *string `pulumi:"usageType"`
}

type PlanDataResponseOutput struct{ *pulumi.OutputState }

func (PlanDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanDataResponse)(nil)).Elem()
}

func (o PlanDataResponseOutput) ToPlanDataResponseOutput() PlanDataResponseOutput {
	return o
}

func (o PlanDataResponseOutput) ToPlanDataResponseOutputWithContext(ctx context.Context) PlanDataResponseOutput {
	return o
}

func (o PlanDataResponseOutput) BillingCycle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanDataResponse) *string { return v.BillingCycle }).(pulumi.StringPtrOutput)
}

func (o PlanDataResponseOutput) EffectiveDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanDataResponse) *string { return v.EffectiveDate }).(pulumi.StringPtrOutput)
}

func (o PlanDataResponseOutput) PlanDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanDataResponse) *string { return v.PlanDetails }).(pulumi.StringPtrOutput)
}

func (o PlanDataResponseOutput) UsageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanDataResponse) *string { return v.UsageType }).(pulumi.StringPtrOutput)
}

type PlanDataResponsePtrOutput struct{ *pulumi.OutputState }

func (PlanDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanDataResponse)(nil)).Elem()
}

func (o PlanDataResponsePtrOutput) ToPlanDataResponsePtrOutput() PlanDataResponsePtrOutput {
	return o
}

func (o PlanDataResponsePtrOutput) ToPlanDataResponsePtrOutputWithContext(ctx context.Context) PlanDataResponsePtrOutput {
	return o
}

func (o PlanDataResponsePtrOutput) Elem() PlanDataResponseOutput {
	return o.ApplyT(func(v *PlanDataResponse) PlanDataResponse {
		if v != nil {
			return *v
		}
		var ret PlanDataResponse
		return ret
	}).(PlanDataResponseOutput)
}

func (o PlanDataResponsePtrOutput) BillingCycle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.BillingCycle
	}).(pulumi.StringPtrOutput)
}

func (o PlanDataResponsePtrOutput) EffectiveDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.EffectiveDate
	}).(pulumi.StringPtrOutput)
}

func (o PlanDataResponsePtrOutput) PlanDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.PlanDetails
	}).(pulumi.StringPtrOutput)
}

func (o PlanDataResponsePtrOutput) UsageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.UsageType
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

type UserInfo struct {
	EmailAddress *string `pulumi:"emailAddress"`
	FirstName    *string `pulumi:"firstName"`
	LastName     *string `pulumi:"lastName"`
	PhoneNumber  *string `pulumi:"phoneNumber"`
}





type UserInfoInput interface {
	pulumi.Input

	ToUserInfoOutput() UserInfoOutput
	ToUserInfoOutputWithContext(context.Context) UserInfoOutput
}

type UserInfoArgs struct {
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	FirstName    pulumi.StringPtrInput `pulumi:"firstName"`
	LastName     pulumi.StringPtrInput `pulumi:"lastName"`
	PhoneNumber  pulumi.StringPtrInput `pulumi:"phoneNumber"`
}

func (UserInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

func (i UserInfoArgs) ToUserInfoOutput() UserInfoOutput {
	return i.ToUserInfoOutputWithContext(context.Background())
}

func (i UserInfoArgs) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoOutput)
}

func (i UserInfoArgs) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return i.ToUserInfoPtrOutputWithContext(context.Background())
}

func (i UserInfoArgs) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoOutput).ToUserInfoPtrOutputWithContext(ctx)
}









type UserInfoPtrInput interface {
	pulumi.Input

	ToUserInfoPtrOutput() UserInfoPtrOutput
	ToUserInfoPtrOutputWithContext(context.Context) UserInfoPtrOutput
}

type userInfoPtrType UserInfoArgs

func UserInfoPtr(v *UserInfoArgs) UserInfoPtrInput {
	return (*userInfoPtrType)(v)
}

func (*userInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfo)(nil)).Elem()
}

func (i *userInfoPtrType) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return i.ToUserInfoPtrOutputWithContext(context.Background())
}

func (i *userInfoPtrType) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoPtrOutput)
}

type UserInfoOutput struct{ *pulumi.OutputState }

func (UserInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

func (o UserInfoOutput) ToUserInfoOutput() UserInfoOutput {
	return o
}

func (o UserInfoOutput) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return o
}

func (o UserInfoOutput) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return o.ToUserInfoPtrOutputWithContext(context.Background())
}

func (o UserInfoOutput) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserInfo) *UserInfo {
		return &v
	}).(UserInfoPtrOutput)
}

func (o UserInfoOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

func (o UserInfoOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

func (o UserInfoOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

func (o UserInfoOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.PhoneNumber }).(pulumi.StringPtrOutput)
}

type UserInfoPtrOutput struct{ *pulumi.OutputState }

func (UserInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfo)(nil)).Elem()
}

func (o UserInfoPtrOutput) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return o
}

func (o UserInfoPtrOutput) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return o
}

func (o UserInfoPtrOutput) Elem() UserInfoOutput {
	return o.ApplyT(func(v *UserInfo) UserInfo {
		if v != nil {
			return *v
		}
		var ret UserInfo
		return ret
	}).(UserInfoOutput)
}

func (o UserInfoPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoPtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoPtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoPtrOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.PhoneNumber
	}).(pulumi.StringPtrOutput)
}

type UserInfoResponse struct {
	EmailAddress *string `pulumi:"emailAddress"`
	FirstName    *string `pulumi:"firstName"`
	LastName     *string `pulumi:"lastName"`
	PhoneNumber  *string `pulumi:"phoneNumber"`
}

type UserInfoResponseOutput struct{ *pulumi.OutputState }

func (UserInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutput() UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutputWithContext(ctx context.Context) UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.PhoneNumber }).(pulumi.StringPtrOutput)
}

type UserInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (UserInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) Elem() UserInfoResponseOutput {
	return o.ApplyT(func(v *UserInfoResponse) UserInfoResponse {
		if v != nil {
			return *v
		}
		var ret UserInfoResponse
		return ret
	}).(UserInfoResponseOutput)
}

func (o UserInfoResponsePtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PhoneNumber
	}).(pulumi.StringPtrOutput)
}

type UserRoleResponseResponse struct {
	Role *string `pulumi:"role"`
}

type UserRoleResponseResponseOutput struct{ *pulumi.OutputState }

func (UserRoleResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRoleResponseResponse)(nil)).Elem()
}

func (o UserRoleResponseResponseOutput) ToUserRoleResponseResponseOutput() UserRoleResponseResponseOutput {
	return o
}

func (o UserRoleResponseResponseOutput) ToUserRoleResponseResponseOutputWithContext(ctx context.Context) UserRoleResponseResponseOutput {
	return o
}

func (o UserRoleResponseResponseOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserRoleResponseResponse) *string { return v.Role }).(pulumi.StringPtrOutput)
}

type UserRoleResponseResponseArrayOutput struct{ *pulumi.OutputState }

func (UserRoleResponseResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserRoleResponseResponse)(nil)).Elem()
}

func (o UserRoleResponseResponseArrayOutput) ToUserRoleResponseResponseArrayOutput() UserRoleResponseResponseArrayOutput {
	return o
}

func (o UserRoleResponseResponseArrayOutput) ToUserRoleResponseResponseArrayOutputWithContext(ctx context.Context) UserRoleResponseResponseArrayOutput {
	return o
}

func (o UserRoleResponseResponseArrayOutput) Index(i pulumi.IntInput) UserRoleResponseResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserRoleResponseResponse {
		return vs[0].([]UserRoleResponseResponse)[vs[1].(int)]
	}).(UserRoleResponseResponseOutput)
}

type VMResourcesResponse struct {
	AgentVersion *string `pulumi:"agentVersion"`
	Id           *string `pulumi:"id"`
}

type VMResourcesResponseOutput struct{ *pulumi.OutputState }

func (VMResourcesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMResourcesResponse)(nil)).Elem()
}

func (o VMResourcesResponseOutput) ToVMResourcesResponseOutput() VMResourcesResponseOutput {
	return o
}

func (o VMResourcesResponseOutput) ToVMResourcesResponseOutputWithContext(ctx context.Context) VMResourcesResponseOutput {
	return o
}

func (o VMResourcesResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMResourcesResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o VMResourcesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMResourcesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VMResourcesResponseArrayOutput struct{ *pulumi.OutputState }

func (VMResourcesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMResourcesResponse)(nil)).Elem()
}

func (o VMResourcesResponseArrayOutput) ToVMResourcesResponseArrayOutput() VMResourcesResponseArrayOutput {
	return o
}

func (o VMResourcesResponseArrayOutput) ToVMResourcesResponseArrayOutputWithContext(ctx context.Context) VMResourcesResponseArrayOutput {
	return o
}

func (o VMResourcesResponseArrayOutput) Index(i pulumi.IntInput) VMResourcesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMResourcesResponse {
		return vs[0].([]VMResourcesResponse)[vs[1].(int)]
	}).(VMResourcesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(FilteringTagOutput{})
	pulumi.RegisterOutputType(FilteringTagArrayOutput{})
	pulumi.RegisterOutputType(FilteringTagResponseOutput{})
	pulumi.RegisterOutputType(FilteringTagResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LogRulesOutput{})
	pulumi.RegisterOutputType(LogRulesPtrOutput{})
	pulumi.RegisterOutputType(LogRulesResponseOutput{})
	pulumi.RegisterOutputType(LogRulesResponsePtrOutput{})
	pulumi.RegisterOutputType(LogzOrganizationPropertiesOutput{})
	pulumi.RegisterOutputType(LogzOrganizationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LogzOrganizationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LogzOrganizationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MonitoredResourceResponseOutput{})
	pulumi.RegisterOutputType(MonitoredResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(MonitoringTagRulesPropertiesOutput{})
	pulumi.RegisterOutputType(MonitoringTagRulesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MonitoringTagRulesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PlanDataOutput{})
	pulumi.RegisterOutputType(PlanDataPtrOutput{})
	pulumi.RegisterOutputType(PlanDataResponseOutput{})
	pulumi.RegisterOutputType(PlanDataResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserInfoOutput{})
	pulumi.RegisterOutputType(UserInfoPtrOutput{})
	pulumi.RegisterOutputType(UserInfoResponseOutput{})
	pulumi.RegisterOutputType(UserInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(UserRoleResponseResponseOutput{})
	pulumi.RegisterOutputType(UserRoleResponseResponseArrayOutput{})
	pulumi.RegisterOutputType(VMResourcesResponseOutput{})
	pulumi.RegisterOutputType(VMResourcesResponseArrayOutput{})
}
