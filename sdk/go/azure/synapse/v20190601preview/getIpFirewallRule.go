


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupIpFirewallRule(ctx *pulumi.Context, args *LookupIpFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupIpFirewallRuleResult, error) {
	var rv LookupIpFirewallRuleResult
	err := ctx.Invoke("azure-native:synapse/v20190601preview:getIpFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpFirewallRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupIpFirewallRuleResult struct {
	EndIpAddress      *string `pulumi:"endIpAddress"`
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	StartIpAddress    *string `pulumi:"startIpAddress"`
	Type              string  `pulumi:"type"`
}

func LookupIpFirewallRuleOutput(ctx *pulumi.Context, args LookupIpFirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupIpFirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpFirewallRuleResult, error) {
			args := v.(LookupIpFirewallRuleArgs)
			r, err := LookupIpFirewallRule(ctx, &args, opts...)
			var s LookupIpFirewallRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpFirewallRuleResultOutput)
}

type LookupIpFirewallRuleOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleName          pulumi.StringInput `pulumi:"ruleName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIpFirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpFirewallRuleArgs)(nil)).Elem()
}


type LookupIpFirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupIpFirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpFirewallRuleResult)(nil)).Elem()
}

func (o LookupIpFirewallRuleResultOutput) ToLookupIpFirewallRuleResultOutput() LookupIpFirewallRuleResultOutput {
	return o
}

func (o LookupIpFirewallRuleResultOutput) ToLookupIpFirewallRuleResultOutputWithContext(ctx context.Context) LookupIpFirewallRuleResultOutput {
	return o
}

func (o LookupIpFirewallRuleResultOutput) EndIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) *string { return v.EndIpAddress }).(pulumi.StringPtrOutput)
}

func (o LookupIpFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIpFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIpFirewallRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupIpFirewallRuleResultOutput) StartIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) *string { return v.StartIpAddress }).(pulumi.StringPtrOutput)
}

func (o LookupIpFirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpFirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpFirewallRuleResultOutput{})
}
