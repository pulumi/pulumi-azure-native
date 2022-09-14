


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDefaultAdminRule(ctx *pulumi.Context, args *LookupDefaultAdminRuleArgs, opts ...pulumi.InvokeOption) (*LookupDefaultAdminRuleResult, error) {
	var rv LookupDefaultAdminRuleResult
	err := ctx.Invoke("azure-native:network/v20220101:getDefaultAdminRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDefaultAdminRuleArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	RuleName           string `pulumi:"ruleName"`
}


type LookupDefaultAdminRuleResult struct {
	Access                string                      `pulumi:"access"`
	Description           string                      `pulumi:"description"`
	DestinationPortRanges []string                    `pulumi:"destinationPortRanges"`
	Destinations          []AddressPrefixItemResponse `pulumi:"destinations"`
	Direction             string                      `pulumi:"direction"`
	Etag                  string                      `pulumi:"etag"`
	Flag                  *string                     `pulumi:"flag"`
	Id                    string                      `pulumi:"id"`
	Kind                  string                      `pulumi:"kind"`
	Name                  string                      `pulumi:"name"`
	Priority              int                         `pulumi:"priority"`
	Protocol              string                      `pulumi:"protocol"`
	ProvisioningState     string                      `pulumi:"provisioningState"`
	SourcePortRanges      []string                    `pulumi:"sourcePortRanges"`
	Sources               []AddressPrefixItemResponse `pulumi:"sources"`
	SystemData            SystemDataResponse          `pulumi:"systemData"`
	Type                  string                      `pulumi:"type"`
}

func LookupDefaultAdminRuleOutput(ctx *pulumi.Context, args LookupDefaultAdminRuleOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultAdminRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultAdminRuleResult, error) {
			args := v.(LookupDefaultAdminRuleArgs)
			r, err := LookupDefaultAdminRule(ctx, &args, opts...)
			var s LookupDefaultAdminRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultAdminRuleResultOutput)
}

type LookupDefaultAdminRuleOutputArgs struct {
	ConfigurationName  pulumi.StringInput `pulumi:"configurationName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
	RuleName           pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupDefaultAdminRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultAdminRuleArgs)(nil)).Elem()
}


type LookupDefaultAdminRuleResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultAdminRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultAdminRuleResult)(nil)).Elem()
}

func (o LookupDefaultAdminRuleResultOutput) ToLookupDefaultAdminRuleResultOutput() LookupDefaultAdminRuleResultOutput {
	return o
}

func (o LookupDefaultAdminRuleResultOutput) ToLookupDefaultAdminRuleResultOutputWithContext(ctx context.Context) LookupDefaultAdminRuleResultOutput {
	return o
}

func (o LookupDefaultAdminRuleResultOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) string { return v.Access }).(pulumi.StringOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupDefaultAdminRuleResultOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) []AddressPrefixItemResponse { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) string { return v.Direction }).(pulumi.StringOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Flag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) *string { return v.Flag }).(pulumi.StringPtrOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) int { return v.Priority }).(pulumi.IntOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LookupDefaultAdminRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDefaultAdminRuleResultOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) []AddressPrefixItemResponse { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
}

func (o LookupDefaultAdminRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDefaultAdminRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultAdminRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultAdminRuleResultOutput{})
}
