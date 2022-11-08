


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRulesEngine(ctx *pulumi.Context, args *LookupRulesEngineArgs, opts ...pulumi.InvokeOption) (*LookupRulesEngineResult, error) {
	var rv LookupRulesEngineResult
	err := ctx.Invoke("azure-native:network/v20200501:getRulesEngine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRulesEngineArgs struct {
	FrontDoorName     string `pulumi:"frontDoorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RulesEngineName   string `pulumi:"rulesEngineName"`
}


type LookupRulesEngineResult struct {
	Id            string                    `pulumi:"id"`
	Name          string                    `pulumi:"name"`
	ResourceState string                    `pulumi:"resourceState"`
	Rules         []RulesEngineRuleResponse `pulumi:"rules"`
	Type          string                    `pulumi:"type"`
}

func LookupRulesEngineOutput(ctx *pulumi.Context, args LookupRulesEngineOutputArgs, opts ...pulumi.InvokeOption) LookupRulesEngineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRulesEngineResult, error) {
			args := v.(LookupRulesEngineArgs)
			r, err := LookupRulesEngine(ctx, &args, opts...)
			var s LookupRulesEngineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRulesEngineResultOutput)
}

type LookupRulesEngineOutputArgs struct {
	FrontDoorName     pulumi.StringInput `pulumi:"frontDoorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RulesEngineName   pulumi.StringInput `pulumi:"rulesEngineName"`
}

func (LookupRulesEngineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRulesEngineArgs)(nil)).Elem()
}


type LookupRulesEngineResultOutput struct{ *pulumi.OutputState }

func (LookupRulesEngineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRulesEngineResult)(nil)).Elem()
}

func (o LookupRulesEngineResultOutput) ToLookupRulesEngineResultOutput() LookupRulesEngineResultOutput {
	return o
}

func (o LookupRulesEngineResultOutput) ToLookupRulesEngineResultOutputWithContext(ctx context.Context) LookupRulesEngineResultOutput {
	return o
}

func (o LookupRulesEngineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRulesEngineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRulesEngineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRulesEngineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRulesEngineResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRulesEngineResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupRulesEngineResultOutput) Rules() RulesEngineRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupRulesEngineResult) []RulesEngineRuleResponse { return v.Rules }).(RulesEngineRuleResponseArrayOutput)
}

func (o LookupRulesEngineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRulesEngineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRulesEngineResultOutput{})
}
