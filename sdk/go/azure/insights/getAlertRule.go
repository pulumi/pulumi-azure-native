


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlertRule(ctx *pulumi.Context, args *LookupAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupAlertRuleResult, error) {
	var rv LookupAlertRuleResult
	err := ctx.Invoke("azure-native:insights:getAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupAlertRuleResult struct {
	Action            interface{}       `pulumi:"action"`
	Actions           []interface{}     `pulumi:"actions"`
	Condition         interface{}       `pulumi:"condition"`
	Description       *string           `pulumi:"description"`
	Id                string            `pulumi:"id"`
	IsEnabled         bool              `pulumi:"isEnabled"`
	LastUpdatedTime   string            `pulumi:"lastUpdatedTime"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}

func LookupAlertRuleOutput(ctx *pulumi.Context, args LookupAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlertRuleResult, error) {
			args := v.(LookupAlertRuleArgs)
			r, err := LookupAlertRule(ctx, &args, opts...)
			var s LookupAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlertRuleResultOutput)
}

type LookupAlertRuleOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleName          pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertRuleArgs)(nil)).Elem()
}


type LookupAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertRuleResult)(nil)).Elem()
}

func (o LookupAlertRuleResultOutput) ToLookupAlertRuleResultOutput() LookupAlertRuleResultOutput {
	return o
}

func (o LookupAlertRuleResultOutput) ToLookupAlertRuleResultOutputWithContext(ctx context.Context) LookupAlertRuleResultOutput {
	return o
}

func (o LookupAlertRuleResultOutput) Action() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) interface{} { return v.Action }).(pulumi.AnyOutput)
}

func (o LookupAlertRuleResultOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o LookupAlertRuleResultOutput) Condition() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) interface{} { return v.Condition }).(pulumi.AnyOutput)
}

func (o LookupAlertRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAlertRuleResultOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o LookupAlertRuleResultOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o LookupAlertRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAlertRuleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupAlertRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlertRuleResultOutput{})
}
