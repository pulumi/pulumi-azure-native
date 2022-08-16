


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAlertRule(ctx *pulumi.Context, args *LookupAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupAlertRuleResult, error) {
	var rv LookupAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20220101preview:getAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAlertRuleResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
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
	RuleId            pulumi.StringInput `pulumi:"ruleId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
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

func (o LookupAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAlertRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlertRuleResultOutput{})
}
