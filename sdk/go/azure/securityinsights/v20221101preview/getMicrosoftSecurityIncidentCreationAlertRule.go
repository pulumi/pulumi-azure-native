


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMicrosoftSecurityIncidentCreationAlertRule(ctx *pulumi.Context, args *LookupMicrosoftSecurityIncidentCreationAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupMicrosoftSecurityIncidentCreationAlertRuleResult, error) {
	var rv LookupMicrosoftSecurityIncidentCreationAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20221101preview:getMicrosoftSecurityIncidentCreationAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMicrosoftSecurityIncidentCreationAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMicrosoftSecurityIncidentCreationAlertRuleResult struct {
	AlertRuleTemplateName     *string            `pulumi:"alertRuleTemplateName"`
	Description               *string            `pulumi:"description"`
	DisplayName               string             `pulumi:"displayName"`
	DisplayNamesExcludeFilter []string           `pulumi:"displayNamesExcludeFilter"`
	DisplayNamesFilter        []string           `pulumi:"displayNamesFilter"`
	Enabled                   bool               `pulumi:"enabled"`
	Etag                      *string            `pulumi:"etag"`
	Id                        string             `pulumi:"id"`
	Kind                      string             `pulumi:"kind"`
	LastModifiedUtc           string             `pulumi:"lastModifiedUtc"`
	Name                      string             `pulumi:"name"`
	ProductFilter             string             `pulumi:"productFilter"`
	SeveritiesFilter          []string           `pulumi:"severitiesFilter"`
	SystemData                SystemDataResponse `pulumi:"systemData"`
	Type                      string             `pulumi:"type"`
}

func LookupMicrosoftSecurityIncidentCreationAlertRuleOutput(ctx *pulumi.Context, args LookupMicrosoftSecurityIncidentCreationAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMicrosoftSecurityIncidentCreationAlertRuleResult, error) {
			args := v.(LookupMicrosoftSecurityIncidentCreationAlertRuleArgs)
			r, err := LookupMicrosoftSecurityIncidentCreationAlertRule(ctx, &args, opts...)
			var s LookupMicrosoftSecurityIncidentCreationAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput)
}

type LookupMicrosoftSecurityIncidentCreationAlertRuleOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleId            pulumi.StringInput `pulumi:"ruleId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMicrosoftSecurityIncidentCreationAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMicrosoftSecurityIncidentCreationAlertRuleArgs)(nil)).Elem()
}


type LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMicrosoftSecurityIncidentCreationAlertRuleResult)(nil)).Elem()
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) ToLookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput() LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput {
	return o
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) ToLookupMicrosoftSecurityIncidentCreationAlertRuleResultOutputWithContext(ctx context.Context) LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput {
	return o
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) *string { return v.AlertRuleTemplateName }).(pulumi.StringPtrOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) DisplayNamesExcludeFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) []string {
		return v.DisplayNamesExcludeFilter
	}).(pulumi.StringArrayOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) DisplayNamesFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) []string { return v.DisplayNamesFilter }).(pulumi.StringArrayOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) ProductFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.ProductFilter }).(pulumi.StringOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) SeveritiesFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) []string { return v.SeveritiesFilter }).(pulumi.StringArrayOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMicrosoftSecurityIncidentCreationAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMicrosoftSecurityIncidentCreationAlertRuleResultOutput{})
}
