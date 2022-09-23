


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRuleSet(ctx *pulumi.Context, args *LookupRuleSetArgs, opts ...pulumi.InvokeOption) (*LookupRuleSetResult, error) {
	var rv LookupRuleSetResult
	err := ctx.Invoke("azure-native:cdn/v20200901:getRuleSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleSetArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleSetName       string `pulumi:"ruleSetName"`
}


type LookupRuleSetResult struct {
	DeploymentStatus  string             `pulumi:"deploymentStatus"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupRuleSetOutput(ctx *pulumi.Context, args LookupRuleSetOutputArgs, opts ...pulumi.InvokeOption) LookupRuleSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleSetResult, error) {
			args := v.(LookupRuleSetArgs)
			r, err := LookupRuleSet(ctx, &args, opts...)
			var s LookupRuleSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRuleSetResultOutput)
}

type LookupRuleSetOutputArgs struct {
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleSetName       pulumi.StringInput `pulumi:"ruleSetName"`
}

func (LookupRuleSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleSetArgs)(nil)).Elem()
}


type LookupRuleSetResultOutput struct{ *pulumi.OutputState }

func (LookupRuleSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleSetResult)(nil)).Elem()
}

func (o LookupRuleSetResultOutput) ToLookupRuleSetResultOutput() LookupRuleSetResultOutput {
	return o
}

func (o LookupRuleSetResultOutput) ToLookupRuleSetResultOutputWithContext(ctx context.Context) LookupRuleSetResultOutput {
	return o
}

func (o LookupRuleSetResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleSetResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o LookupRuleSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRuleSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRuleSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRuleSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRuleSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRuleSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleSetResultOutput{})
}
