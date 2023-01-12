


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupUserRule(ctx *pulumi.Context, args *LookupUserRuleArgs, opts ...pulumi.InvokeOption) (*LookupUserRuleResult, error) {
	var rv LookupUserRuleResult
	err := ctx.Invoke("azure-native:network/v20220401preview:getUserRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserRuleArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	RuleName           string `pulumi:"ruleName"`
}


type LookupUserRuleResult struct {
	Etag       string             `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupUserRuleOutput(ctx *pulumi.Context, args LookupUserRuleOutputArgs, opts ...pulumi.InvokeOption) LookupUserRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserRuleResult, error) {
			args := v.(LookupUserRuleArgs)
			r, err := LookupUserRule(ctx, &args, opts...)
			var s LookupUserRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserRuleResultOutput)
}

type LookupUserRuleOutputArgs struct {
	ConfigurationName  pulumi.StringInput `pulumi:"configurationName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
	RuleName           pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupUserRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserRuleArgs)(nil)).Elem()
}


type LookupUserRuleResultOutput struct{ *pulumi.OutputState }

func (LookupUserRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserRuleResult)(nil)).Elem()
}

func (o LookupUserRuleResultOutput) ToLookupUserRuleResultOutput() LookupUserRuleResultOutput {
	return o
}

func (o LookupUserRuleResultOutput) ToLookupUserRuleResultOutputWithContext(ctx context.Context) LookupUserRuleResultOutput {
	return o
}

func (o LookupUserRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupUserRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUserRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupUserRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUserRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUserRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupUserRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserRuleResultOutput{})
}
