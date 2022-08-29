


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDeviceSecurityGroup(ctx *pulumi.Context, args *LookupDeviceSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupDeviceSecurityGroupResult, error) {
	var rv LookupDeviceSecurityGroupResult
	err := ctx.Invoke("azure-native:security:getDeviceSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceSecurityGroupArgs struct {
	DeviceSecurityGroupName string `pulumi:"deviceSecurityGroupName"`
	ResourceId              string `pulumi:"resourceId"`
}


type LookupDeviceSecurityGroupResult struct {
	AllowlistRules  []AllowlistCustomAlertRuleResponse  `pulumi:"allowlistRules"`
	DenylistRules   []DenylistCustomAlertRuleResponse   `pulumi:"denylistRules"`
	Id              string                              `pulumi:"id"`
	Name            string                              `pulumi:"name"`
	ThresholdRules  []ThresholdCustomAlertRuleResponse  `pulumi:"thresholdRules"`
	TimeWindowRules []TimeWindowCustomAlertRuleResponse `pulumi:"timeWindowRules"`
	Type            string                              `pulumi:"type"`
}

func LookupDeviceSecurityGroupOutput(ctx *pulumi.Context, args LookupDeviceSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceSecurityGroupResult, error) {
			args := v.(LookupDeviceSecurityGroupArgs)
			r, err := LookupDeviceSecurityGroup(ctx, &args, opts...)
			var s LookupDeviceSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeviceSecurityGroupResultOutput)
}

type LookupDeviceSecurityGroupOutputArgs struct {
	DeviceSecurityGroupName pulumi.StringInput `pulumi:"deviceSecurityGroupName"`
	ResourceId              pulumi.StringInput `pulumi:"resourceId"`
}

func (LookupDeviceSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceSecurityGroupArgs)(nil)).Elem()
}


type LookupDeviceSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceSecurityGroupResult)(nil)).Elem()
}

func (o LookupDeviceSecurityGroupResultOutput) ToLookupDeviceSecurityGroupResultOutput() LookupDeviceSecurityGroupResultOutput {
	return o
}

func (o LookupDeviceSecurityGroupResultOutput) ToLookupDeviceSecurityGroupResultOutputWithContext(ctx context.Context) LookupDeviceSecurityGroupResultOutput {
	return o
}

func (o LookupDeviceSecurityGroupResultOutput) AllowlistRules() AllowlistCustomAlertRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDeviceSecurityGroupResult) []AllowlistCustomAlertRuleResponse { return v.AllowlistRules }).(AllowlistCustomAlertRuleResponseArrayOutput)
}

func (o LookupDeviceSecurityGroupResultOutput) DenylistRules() DenylistCustomAlertRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDeviceSecurityGroupResult) []DenylistCustomAlertRuleResponse { return v.DenylistRules }).(DenylistCustomAlertRuleResponseArrayOutput)
}

func (o LookupDeviceSecurityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSecurityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeviceSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeviceSecurityGroupResultOutput) ThresholdRules() ThresholdCustomAlertRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDeviceSecurityGroupResult) []ThresholdCustomAlertRuleResponse { return v.ThresholdRules }).(ThresholdCustomAlertRuleResponseArrayOutput)
}

func (o LookupDeviceSecurityGroupResultOutput) TimeWindowRules() TimeWindowCustomAlertRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDeviceSecurityGroupResult) []TimeWindowCustomAlertRuleResponse { return v.TimeWindowRules }).(TimeWindowCustomAlertRuleResponseArrayOutput)
}

func (o LookupDeviceSecurityGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceSecurityGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceSecurityGroupResultOutput{})
}
