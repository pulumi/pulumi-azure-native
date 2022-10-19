


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAutomationRule(ctx *pulumi.Context, args *LookupAutomationRuleArgs, opts ...pulumi.InvokeOption) (*LookupAutomationRuleResult, error) {
	var rv LookupAutomationRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20220101preview:getAutomationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutomationRuleArgs struct {
	AutomationRuleId  string `pulumi:"automationRuleId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type LookupAutomationRuleResult struct {
	Actions             []interface{}                         `pulumi:"actions"`
	CreatedBy           ClientInfoResponse                    `pulumi:"createdBy"`
	CreatedTimeUtc      string                                `pulumi:"createdTimeUtc"`
	DisplayName         string                                `pulumi:"displayName"`
	Etag                *string                               `pulumi:"etag"`
	Id                  string                                `pulumi:"id"`
	LastModifiedBy      ClientInfoResponse                    `pulumi:"lastModifiedBy"`
	LastModifiedTimeUtc string                                `pulumi:"lastModifiedTimeUtc"`
	Name                string                                `pulumi:"name"`
	Order               int                                   `pulumi:"order"`
	SystemData          SystemDataResponse                    `pulumi:"systemData"`
	TriggeringLogic     AutomationRuleTriggeringLogicResponse `pulumi:"triggeringLogic"`
	Type                string                                `pulumi:"type"`
}

func LookupAutomationRuleOutput(ctx *pulumi.Context, args LookupAutomationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupAutomationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutomationRuleResult, error) {
			args := v.(LookupAutomationRuleArgs)
			r, err := LookupAutomationRule(ctx, &args, opts...)
			var s LookupAutomationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutomationRuleResultOutput)
}

type LookupAutomationRuleOutputArgs struct {
	AutomationRuleId  pulumi.StringInput `pulumi:"automationRuleId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAutomationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutomationRuleArgs)(nil)).Elem()
}

type LookupAutomationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupAutomationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutomationRuleResult)(nil)).Elem()
}

func (o LookupAutomationRuleResultOutput) ToLookupAutomationRuleResultOutput() LookupAutomationRuleResultOutput {
	return o
}

func (o LookupAutomationRuleResultOutput) ToLookupAutomationRuleResultOutputWithContext(ctx context.Context) LookupAutomationRuleResultOutput {
	return o
}

func (o LookupAutomationRuleResultOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o LookupAutomationRuleResultOutput) CreatedBy() ClientInfoResponseOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) ClientInfoResponse { return v.CreatedBy }).(ClientInfoResponseOutput)
}

func (o LookupAutomationRuleResultOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupAutomationRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupAutomationRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAutomationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAutomationRuleResultOutput) LastModifiedBy() ClientInfoResponseOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) ClientInfoResponse { return v.LastModifiedBy }).(ClientInfoResponseOutput)
}

func (o LookupAutomationRuleResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupAutomationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAutomationRuleResultOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) int { return v.Order }).(pulumi.IntOutput)
}

func (o LookupAutomationRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAutomationRuleResultOutput) TriggeringLogic() AutomationRuleTriggeringLogicResponseOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) AutomationRuleTriggeringLogicResponse { return v.TriggeringLogic }).(AutomationRuleTriggeringLogicResponseOutput)
}

func (o LookupAutomationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutomationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutomationRuleResultOutput{})
}
