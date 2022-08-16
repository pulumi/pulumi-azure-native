


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAdminRule(ctx *pulumi.Context, args *LookupAdminRuleArgs, opts ...pulumi.InvokeOption) (*LookupAdminRuleResult, error) {
	var rv LookupAdminRuleResult
	err := ctx.Invoke("azure-native:network/v20210201preview:getAdminRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdminRuleArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	RuleName           string `pulumi:"ruleName"`
}


type LookupAdminRuleResult struct {
	Etag       string             `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupAdminRuleOutput(ctx *pulumi.Context, args LookupAdminRuleOutputArgs, opts ...pulumi.InvokeOption) LookupAdminRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAdminRuleResult, error) {
			args := v.(LookupAdminRuleArgs)
			r, err := LookupAdminRule(ctx, &args, opts...)
			var s LookupAdminRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAdminRuleResultOutput)
}

type LookupAdminRuleOutputArgs struct {
	ConfigurationName  pulumi.StringInput `pulumi:"configurationName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
	RuleName           pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupAdminRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdminRuleArgs)(nil)).Elem()
}


type LookupAdminRuleResultOutput struct{ *pulumi.OutputState }

func (LookupAdminRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdminRuleResult)(nil)).Elem()
}

func (o LookupAdminRuleResultOutput) ToLookupAdminRuleResultOutput() LookupAdminRuleResultOutput {
	return o
}

func (o LookupAdminRuleResultOutput) ToLookupAdminRuleResultOutputWithContext(ctx context.Context) LookupAdminRuleResultOutput {
	return o
}

func (o LookupAdminRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupAdminRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAdminRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAdminRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAdminRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAdminRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdminRuleResultOutput{})
}
