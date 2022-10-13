


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMetricAlert(ctx *pulumi.Context, args *LookupMetricAlertArgs, opts ...pulumi.InvokeOption) (*LookupMetricAlertResult, error) {
	var rv LookupMetricAlertResult
	err := ctx.Invoke("azure-native:insights:getMetricAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetricAlertArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupMetricAlertResult struct {
	Actions              []MetricAlertActionResponse `pulumi:"actions"`
	AutoMitigate         *bool                       `pulumi:"autoMitigate"`
	Criteria             interface{}                 `pulumi:"criteria"`
	Description          *string                     `pulumi:"description"`
	Enabled              bool                        `pulumi:"enabled"`
	EvaluationFrequency  string                      `pulumi:"evaluationFrequency"`
	Id                   string                      `pulumi:"id"`
	IsMigrated           bool                        `pulumi:"isMigrated"`
	LastUpdatedTime      string                      `pulumi:"lastUpdatedTime"`
	Location             string                      `pulumi:"location"`
	Name                 string                      `pulumi:"name"`
	Scopes               []string                    `pulumi:"scopes"`
	Severity             int                         `pulumi:"severity"`
	Tags                 map[string]string           `pulumi:"tags"`
	TargetResourceRegion *string                     `pulumi:"targetResourceRegion"`
	TargetResourceType   *string                     `pulumi:"targetResourceType"`
	Type                 string                      `pulumi:"type"`
	WindowSize           string                      `pulumi:"windowSize"`
}

func LookupMetricAlertOutput(ctx *pulumi.Context, args LookupMetricAlertOutputArgs, opts ...pulumi.InvokeOption) LookupMetricAlertResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetricAlertResult, error) {
			args := v.(LookupMetricAlertArgs)
			r, err := LookupMetricAlert(ctx, &args, opts...)
			var s LookupMetricAlertResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetricAlertResultOutput)
}

type LookupMetricAlertOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleName          pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupMetricAlertOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricAlertArgs)(nil)).Elem()
}


type LookupMetricAlertResultOutput struct{ *pulumi.OutputState }

func (LookupMetricAlertResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricAlertResult)(nil)).Elem()
}

func (o LookupMetricAlertResultOutput) ToLookupMetricAlertResultOutput() LookupMetricAlertResultOutput {
	return o
}

func (o LookupMetricAlertResultOutput) ToLookupMetricAlertResultOutputWithContext(ctx context.Context) LookupMetricAlertResultOutput {
	return o
}

func (o LookupMetricAlertResultOutput) Actions() MetricAlertActionResponseArrayOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) []MetricAlertActionResponse { return v.Actions }).(MetricAlertActionResponseArrayOutput)
}

func (o LookupMetricAlertResultOutput) AutoMitigate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) *bool { return v.AutoMitigate }).(pulumi.BoolPtrOutput)
}

func (o LookupMetricAlertResultOutput) Criteria() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) interface{} { return v.Criteria }).(pulumi.AnyOutput)
}

func (o LookupMetricAlertResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMetricAlertResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupMetricAlertResultOutput) EvaluationFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) string { return v.EvaluationFrequency }).(pulumi.StringOutput)
}

func (o LookupMetricAlertResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMetricAlertResultOutput) IsMigrated() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) bool { return v.IsMigrated }).(pulumi.BoolOutput)
}

func (o LookupMetricAlertResultOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o LookupMetricAlertResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMetricAlertResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMetricAlertResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o LookupMetricAlertResultOutput) Severity() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) int { return v.Severity }).(pulumi.IntOutput)
}

func (o LookupMetricAlertResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMetricAlertResultOutput) TargetResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) *string { return v.TargetResourceRegion }).(pulumi.StringPtrOutput)
}

func (o LookupMetricAlertResultOutput) TargetResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) *string { return v.TargetResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupMetricAlertResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupMetricAlertResultOutput) WindowSize() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricAlertResult) string { return v.WindowSize }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetricAlertResultOutput{})
}
