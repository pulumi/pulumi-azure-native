


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSmartDetectorAlertRule(ctx *pulumi.Context, args *LookupSmartDetectorAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupSmartDetectorAlertRuleResult, error) {
	var rv LookupSmartDetectorAlertRuleResult
	err := ctx.Invoke("azure-native:alertsmanagement/v20210401:getSmartDetectorAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSmartDetectorAlertRuleArgs struct {
	AlertRuleName     string `pulumi:"alertRuleName"`
	ExpandDetector    *bool  `pulumi:"expandDetector"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSmartDetectorAlertRuleResult struct {
	ActionGroups ActionGroupsInformationResponse `pulumi:"actionGroups"`
	Description  *string                         `pulumi:"description"`
	Detector     DetectorResponse                `pulumi:"detector"`
	Frequency    string                          `pulumi:"frequency"`
	Id           string                          `pulumi:"id"`
	Location     *string                         `pulumi:"location"`
	Name         string                          `pulumi:"name"`
	Scope        []string                        `pulumi:"scope"`
	Severity     string                          `pulumi:"severity"`
	State        string                          `pulumi:"state"`
	Tags         map[string]string               `pulumi:"tags"`
	Throttling   *ThrottlingInformationResponse  `pulumi:"throttling"`
	Type         string                          `pulumi:"type"`
}


func (val *LookupSmartDetectorAlertRuleResult) Defaults() *LookupSmartDetectorAlertRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

func LookupSmartDetectorAlertRuleOutput(ctx *pulumi.Context, args LookupSmartDetectorAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupSmartDetectorAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSmartDetectorAlertRuleResult, error) {
			args := v.(LookupSmartDetectorAlertRuleArgs)
			r, err := LookupSmartDetectorAlertRule(ctx, &args, opts...)
			var s LookupSmartDetectorAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSmartDetectorAlertRuleResultOutput)
}

type LookupSmartDetectorAlertRuleOutputArgs struct {
	AlertRuleName     pulumi.StringInput  `pulumi:"alertRuleName"`
	ExpandDetector    pulumi.BoolPtrInput `pulumi:"expandDetector"`
	ResourceGroupName pulumi.StringInput  `pulumi:"resourceGroupName"`
}

func (LookupSmartDetectorAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSmartDetectorAlertRuleArgs)(nil)).Elem()
}


type LookupSmartDetectorAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupSmartDetectorAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSmartDetectorAlertRuleResult)(nil)).Elem()
}

func (o LookupSmartDetectorAlertRuleResultOutput) ToLookupSmartDetectorAlertRuleResultOutput() LookupSmartDetectorAlertRuleResultOutput {
	return o
}

func (o LookupSmartDetectorAlertRuleResultOutput) ToLookupSmartDetectorAlertRuleResultOutputWithContext(ctx context.Context) LookupSmartDetectorAlertRuleResultOutput {
	return o
}

func (o LookupSmartDetectorAlertRuleResultOutput) ActionGroups() ActionGroupsInformationResponseOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) ActionGroupsInformationResponse { return v.ActionGroups }).(ActionGroupsInformationResponseOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Detector() DetectorResponseOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) DetectorResponse { return v.Detector }).(DetectorResponseOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Scope() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) []string { return v.Scope }).(pulumi.StringArrayOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Throttling() ThrottlingInformationResponsePtrOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) *ThrottlingInformationResponse { return v.Throttling }).(ThrottlingInformationResponsePtrOutput)
}

func (o LookupSmartDetectorAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartDetectorAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSmartDetectorAlertRuleResultOutput{})
}
