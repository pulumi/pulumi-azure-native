


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNrtAlertRule(ctx *pulumi.Context, args *LookupNrtAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupNrtAlertRuleResult, error) {
	var rv LookupNrtAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20221001preview:getNrtAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNrtAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupNrtAlertRuleResult struct {
	AlertDetailsOverride     *AlertDetailsOverrideResponse   `pulumi:"alertDetailsOverride"`
	AlertRuleTemplateName    *string                         `pulumi:"alertRuleTemplateName"`
	CustomDetails            map[string]string               `pulumi:"customDetails"`
	Description              *string                         `pulumi:"description"`
	DisplayName              string                          `pulumi:"displayName"`
	Enabled                  bool                            `pulumi:"enabled"`
	EntityMappings           []EntityMappingResponse         `pulumi:"entityMappings"`
	Etag                     *string                         `pulumi:"etag"`
	EventGroupingSettings    *EventGroupingSettingsResponse  `pulumi:"eventGroupingSettings"`
	Id                       string                          `pulumi:"id"`
	IncidentConfiguration    *IncidentConfigurationResponse  `pulumi:"incidentConfiguration"`
	Kind                     string                          `pulumi:"kind"`
	LastModifiedUtc          string                          `pulumi:"lastModifiedUtc"`
	Name                     string                          `pulumi:"name"`
	Query                    string                          `pulumi:"query"`
	SentinelEntitiesMappings []SentinelEntityMappingResponse `pulumi:"sentinelEntitiesMappings"`
	Severity                 string                          `pulumi:"severity"`
	SuppressionDuration      string                          `pulumi:"suppressionDuration"`
	SuppressionEnabled       bool                            `pulumi:"suppressionEnabled"`
	SystemData               SystemDataResponse              `pulumi:"systemData"`
	Tactics                  []string                        `pulumi:"tactics"`
	Techniques               []string                        `pulumi:"techniques"`
	TemplateVersion          *string                         `pulumi:"templateVersion"`
	Type                     string                          `pulumi:"type"`
}

func LookupNrtAlertRuleOutput(ctx *pulumi.Context, args LookupNrtAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNrtAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNrtAlertRuleResult, error) {
			args := v.(LookupNrtAlertRuleArgs)
			r, err := LookupNrtAlertRule(ctx, &args, opts...)
			var s LookupNrtAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNrtAlertRuleResultOutput)
}

type LookupNrtAlertRuleOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleId            pulumi.StringInput `pulumi:"ruleId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupNrtAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNrtAlertRuleArgs)(nil)).Elem()
}


type LookupNrtAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNrtAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNrtAlertRuleResult)(nil)).Elem()
}

func (o LookupNrtAlertRuleResultOutput) ToLookupNrtAlertRuleResultOutput() LookupNrtAlertRuleResultOutput {
	return o
}

func (o LookupNrtAlertRuleResultOutput) ToLookupNrtAlertRuleResultOutputWithContext(ctx context.Context) LookupNrtAlertRuleResultOutput {
	return o
}

func (o LookupNrtAlertRuleResultOutput) AlertDetailsOverride() AlertDetailsOverrideResponsePtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *AlertDetailsOverrideResponse { return v.AlertDetailsOverride }).(AlertDetailsOverrideResponsePtrOutput)
}

func (o LookupNrtAlertRuleResultOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *string { return v.AlertRuleTemplateName }).(pulumi.StringPtrOutput)
}

func (o LookupNrtAlertRuleResultOutput) CustomDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) map[string]string { return v.CustomDetails }).(pulumi.StringMapOutput)
}

func (o LookupNrtAlertRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupNrtAlertRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupNrtAlertRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupNrtAlertRuleResultOutput) EntityMappings() EntityMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) []EntityMappingResponse { return v.EntityMappings }).(EntityMappingResponseArrayOutput)
}

func (o LookupNrtAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupNrtAlertRuleResultOutput) EventGroupingSettings() EventGroupingSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *EventGroupingSettingsResponse { return v.EventGroupingSettings }).(EventGroupingSettingsResponsePtrOutput)
}

func (o LookupNrtAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNrtAlertRuleResultOutput) IncidentConfiguration() IncidentConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *IncidentConfigurationResponse { return v.IncidentConfiguration }).(IncidentConfigurationResponsePtrOutput)
}

func (o LookupNrtAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupNrtAlertRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o LookupNrtAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNrtAlertRuleResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Query }).(pulumi.StringOutput)
}

func (o LookupNrtAlertRuleResultOutput) SentinelEntitiesMappings() SentinelEntityMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) []SentinelEntityMappingResponse { return v.SentinelEntitiesMappings }).(SentinelEntityMappingResponseArrayOutput)
}

func (o LookupNrtAlertRuleResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupNrtAlertRuleResultOutput) SuppressionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.SuppressionDuration }).(pulumi.StringOutput)
}

func (o LookupNrtAlertRuleResultOutput) SuppressionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) bool { return v.SuppressionEnabled }).(pulumi.BoolOutput)
}

func (o LookupNrtAlertRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupNrtAlertRuleResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o LookupNrtAlertRuleResultOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) []string { return v.Techniques }).(pulumi.StringArrayOutput)
}

func (o LookupNrtAlertRuleResultOutput) TemplateVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) *string { return v.TemplateVersion }).(pulumi.StringPtrOutput)
}

func (o LookupNrtAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNrtAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNrtAlertRuleResultOutput{})
}
