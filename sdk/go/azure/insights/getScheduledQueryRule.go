


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledQueryRule(ctx *pulumi.Context, args *LookupScheduledQueryRuleArgs, opts ...pulumi.InvokeOption) (*LookupScheduledQueryRuleResult, error) {
	var rv LookupScheduledQueryRuleResult
	err := ctx.Invoke("azure-native:insights:getScheduledQueryRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupScheduledQueryRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupScheduledQueryRuleResult struct {
	Action                   interface{}       `pulumi:"action"`
	AutoMitigate             *bool             `pulumi:"autoMitigate"`
	CreatedWithApiVersion    string            `pulumi:"createdWithApiVersion"`
	Description              *string           `pulumi:"description"`
	DisplayName              *string           `pulumi:"displayName"`
	Enabled                  *string           `pulumi:"enabled"`
	Etag                     string            `pulumi:"etag"`
	Id                       string            `pulumi:"id"`
	IsLegacyLogAnalyticsRule bool              `pulumi:"isLegacyLogAnalyticsRule"`
	Kind                     string            `pulumi:"kind"`
	LastUpdatedTime          string            `pulumi:"lastUpdatedTime"`
	Location                 string            `pulumi:"location"`
	Name                     string            `pulumi:"name"`
	ProvisioningState        string            `pulumi:"provisioningState"`
	Schedule                 *ScheduleResponse `pulumi:"schedule"`
	Source                   SourceResponse    `pulumi:"source"`
	Tags                     map[string]string `pulumi:"tags"`
	Type                     string            `pulumi:"type"`
}


func (val *LookupScheduledQueryRuleResult) Defaults() *LookupScheduledQueryRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AutoMitigate) {
		autoMitigate_ := false
		tmp.AutoMitigate = &autoMitigate_
	}
	return &tmp
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

func (o LookupScheduledQueryRuleResultOutput) Action() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) interface{} { return v.Action }).(pulumi.AnyOutput)
}

func (o LookupScheduledQueryRuleResultOutput) AutoMitigate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *bool { return v.AutoMitigate }).(pulumi.BoolPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) CreatedWithApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.CreatedWithApiVersion }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) IsLegacyLogAnalyticsRule() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) bool { return v.IsLegacyLogAnalyticsRule }).(pulumi.BoolOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Schedule() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) *ScheduleResponse { return v.Schedule }).(ScheduleResponsePtrOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Source() SourceResponseOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) SourceResponse { return v.Source }).(SourceResponseOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupScheduledQueryRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledQueryRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledQueryRuleResultOutput{})
}
