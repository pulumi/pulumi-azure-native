


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupForwardingRule(ctx *pulumi.Context, args *LookupForwardingRuleArgs, opts ...pulumi.InvokeOption) (*LookupForwardingRuleResult, error) {
	var rv LookupForwardingRuleResult
	err := ctx.Invoke("azure-native:network/v20200401preview:getForwardingRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupForwardingRuleArgs struct {
	DnsForwardingRulesetName string `pulumi:"dnsForwardingRulesetName"`
	ForwardingRuleName       string `pulumi:"forwardingRuleName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type LookupForwardingRuleResult struct {
	DomainName          string                    `pulumi:"domainName"`
	Etag                string                    `pulumi:"etag"`
	ForwardingRuleState *string                   `pulumi:"forwardingRuleState"`
	Id                  string                    `pulumi:"id"`
	Metadata            map[string]string         `pulumi:"metadata"`
	Name                string                    `pulumi:"name"`
	ProvisioningState   string                    `pulumi:"provisioningState"`
	SystemData          SystemDataResponse        `pulumi:"systemData"`
	TargetDnsServers    []TargetDnsServerResponse `pulumi:"targetDnsServers"`
	Type                string                    `pulumi:"type"`
}

func LookupForwardingRuleOutput(ctx *pulumi.Context, args LookupForwardingRuleOutputArgs, opts ...pulumi.InvokeOption) LookupForwardingRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupForwardingRuleResult, error) {
			args := v.(LookupForwardingRuleArgs)
			r, err := LookupForwardingRule(ctx, &args, opts...)
			var s LookupForwardingRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupForwardingRuleResultOutput)
}

type LookupForwardingRuleOutputArgs struct {
	DnsForwardingRulesetName pulumi.StringInput `pulumi:"dnsForwardingRulesetName"`
	ForwardingRuleName       pulumi.StringInput `pulumi:"forwardingRuleName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupForwardingRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupForwardingRuleArgs)(nil)).Elem()
}


type LookupForwardingRuleResultOutput struct{ *pulumi.OutputState }

func (LookupForwardingRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupForwardingRuleResult)(nil)).Elem()
}

func (o LookupForwardingRuleResultOutput) ToLookupForwardingRuleResultOutput() LookupForwardingRuleResultOutput {
	return o
}

func (o LookupForwardingRuleResultOutput) ToLookupForwardingRuleResultOutputWithContext(ctx context.Context) LookupForwardingRuleResultOutput {
	return o
}

func (o LookupForwardingRuleResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o LookupForwardingRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupForwardingRuleResultOutput) ForwardingRuleState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) *string { return v.ForwardingRuleState }).(pulumi.StringPtrOutput)
}

func (o LookupForwardingRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupForwardingRuleResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupForwardingRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupForwardingRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupForwardingRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupForwardingRuleResultOutput) TargetDnsServers() TargetDnsServerResponseArrayOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) []TargetDnsServerResponse { return v.TargetDnsServers }).(TargetDnsServerResponseArrayOutput)
}

func (o LookupForwardingRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupForwardingRuleResultOutput{})
}
