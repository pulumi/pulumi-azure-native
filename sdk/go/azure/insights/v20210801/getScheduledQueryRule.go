


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledQueryRule(ctx *pulumi.Context, args *LookupScheduledQueryRuleArgs, opts ...pulumi.InvokeOption) (*LookupScheduledQueryRuleResult, error) {
	var rv LookupScheduledQueryRuleResult
	err := ctx.Invoke("azure-native:insights/v20210801:getScheduledQueryRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledQueryRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupScheduledQueryRuleResult struct {
	Actions                               *ActionsResponse                   `pulumi:"actions"`
	AutoMitigate                          *bool                              `pulumi:"autoMitigate"`
	CheckWorkspaceAlertsStorageConfigured *bool                              `pulumi:"checkWorkspaceAlertsStorageConfigured"`
	CreatedWithApiVersion                 string                             `pulumi:"createdWithApiVersion"`
	Criteria                              ScheduledQueryRuleCriteriaResponse `pulumi:"criteria"`
	Description                           *string                            `pulumi:"description"`
	DisplayName                           *string                            `pulumi:"displayName"`
	Enabled                               bool                               `pulumi:"enabled"`
	Etag                                  string                             `pulumi:"etag"`
	EvaluationFrequency                   *string                            `pulumi:"evaluationFrequency"`
	Id                                    string                             `pulumi:"id"`
	IsLegacyLogAnalyticsRule              bool                               `pulumi:"isLegacyLogAnalyticsRule"`
	IsWorkspaceAlertsStorageConfigured    bool                               `pulumi:"isWorkspaceAlertsStorageConfigured"`
	Kind                                  *string                            `pulumi:"kind"`
	Location                              string                             `pulumi:"location"`
	MuteActionsDuration                   *string                            `pulumi:"muteActionsDuration"`
	Name                                  string                             `pulumi:"name"`
	OverrideQueryTimeRange                *string                            `pulumi:"overrideQueryTimeRange"`
	Scopes                                []string                           `pulumi:"scopes"`
	Severity                              *float64                           `pulumi:"severity"`
	SkipQueryValidation                   *bool                              `pulumi:"skipQueryValidation"`
	SystemData                            SystemDataResponse                 `pulumi:"systemData"`
	Tags                                  map[string]string                  `pulumi:"tags"`
	TargetResourceTypes                   []string                           `pulumi:"targetResourceTypes"`
	Type                                  string                             `pulumi:"type"`
	WindowSize                            *string                            `pulumi:"windowSize"`
}

func LookupScheduledQueryRuleOutput(ctx *pulumi.Context, args LookupScheduledQueryRuleOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledQueryRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledQueryRuleResult, error) {
			args := v.(LookupScheduledQueryRuleArgs)
			r, err := LookupScheduledQueryRule(ctx, &args, opts...)
			var s LookupScheduledQueryRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledQueryRuleResultOutput)
}

type LookupScheduledQueryRuleOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleName          pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupScheduledQueryRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledQueryRuleArgs)(nil)).Elem()
}


type LookupScheduledQueryRuleResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledQueryRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledQueryRuleResult)(nil)).Elem()
}

func (o LookupScheduledQueryRuleResultOutput) ToLookupScheduledQueryRuleResultOutput() LookupScheduledQueryRuleResultOutput {
	return o
}

func (o LookupScheduledQueryRuleResultOutput) ToLookupScheduledQueryRuleResultOutputWithContext(ctx context.Context) LookupScheduledQueryRuleResultOutput {
	return o
}

func (o LookupScheduledQueryRuleResultOutput) Actions() ActionsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *ActionsResponse { return v.Actions }).(ActionsResponsePtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) AutoMitigate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *bool { return v.AutoMitigate }).(pulumi.BoolPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) CheckWorkspaceAlertsStorageConfigured() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *bool { return v.CheckWorkspaceAlertsStorageConfigured }).(pulumi.BoolPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) CreatedWithApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.CreatedWithApiVersion }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Criteria() ScheduledQueryRuleCriteriaResponseOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) ScheduledQueryRuleCriteriaResponse { return v.Criteria }).(ScheduledQueryRuleCriteriaResponseOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) EvaluationFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.EvaluationFrequency }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) IsLegacyLogAnalyticsRule() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) bool { return v.IsLegacyLogAnalyticsRule }).(pulumi.BoolOutput)
}

func (o LookupScheduledQueryRuleResultOutput) IsWorkspaceAlertsStorageConfigured() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) bool { return v.IsWorkspaceAlertsStorageConfigured }).(pulumi.BoolOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) MuteActionsDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.MuteActionsDuration }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) OverrideQueryTimeRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.OverrideQueryTimeRange }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Severity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *float64 { return v.Severity }).(pulumi.Float64PtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) SkipQueryValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *bool { return v.SkipQueryValidation }).(pulumi.BoolPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupScheduledQueryRuleResultOutput) TargetResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) []string { return v.TargetResourceTypes }).(pulumi.StringArrayOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledQueryRuleResultOutput{})
}
