


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDefaultUserRule(ctx *pulumi.Context, args *LookupDefaultUserRuleArgs, opts ...pulumi.InvokeOption) (*LookupDefaultUserRuleResult, error) {
	var rv LookupDefaultUserRuleResult
	err := ctx.Invoke("azure-native:network/v20210501preview:getDefaultUserRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDefaultUserRuleArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	RuleName           string `pulumi:"ruleName"`
}


type LookupDefaultUserRuleResult struct {
	Description           string                      `pulumi:"description"`
	DestinationPortRanges []string                    `pulumi:"destinationPortRanges"`
	Destinations          []AddressPrefixItemResponse `pulumi:"destinations"`
	Direction             string                      `pulumi:"direction"`
	DisplayName           string                      `pulumi:"displayName"`
	Etag                  string                      `pulumi:"etag"`
	Flag                  *string                     `pulumi:"flag"`
	Id                    string                      `pulumi:"id"`
	Kind                  string                      `pulumi:"kind"`
	Name                  string                      `pulumi:"name"`
	Protocol              string                      `pulumi:"protocol"`
	ProvisioningState     string                      `pulumi:"provisioningState"`
	SourcePortRanges      []string                    `pulumi:"sourcePortRanges"`
	Sources               []AddressPrefixItemResponse `pulumi:"sources"`
	SystemData            SystemDataResponse          `pulumi:"systemData"`
	Type                  string                      `pulumi:"type"`
}

func LookupDefaultUserRuleOutput(ctx *pulumi.Context, args LookupDefaultUserRuleOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultUserRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultUserRuleResult, error) {
			args := v.(LookupDefaultUserRuleArgs)
			r, err := LookupDefaultUserRule(ctx, &args, opts...)
			var s LookupDefaultUserRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultUserRuleResultOutput)
}

type LookupDefaultUserRuleOutputArgs struct {
	ConfigurationName  pulumi.StringInput `pulumi:"configurationName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
	RuleName           pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupDefaultUserRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultUserRuleArgs)(nil)).Elem()
}


type LookupDefaultUserRuleResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultUserRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultUserRuleResult)(nil)).Elem()
}

func (o LookupDefaultUserRuleResultOutput) ToLookupDefaultUserRuleResultOutput() LookupDefaultUserRuleResultOutput {
	return o
}

func (o LookupDefaultUserRuleResultOutput) ToLookupDefaultUserRuleResultOutputWithContext(ctx context.Context) LookupDefaultUserRuleResultOutput {
	return o
}

func (o LookupDefaultUserRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupDefaultUserRuleResultOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o LookupDefaultUserRuleResultOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) []AddressPrefixItemResponse { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

func (o LookupDefaultUserRuleResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) string { return v.Direction }).(pulumi.StringOutput)
}

func (o LookupDefaultUserRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupDefaultUserRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDefaultUserRuleResultOutput) Flag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) *string { return v.Flag }).(pulumi.StringPtrOutput)
}

func (o LookupDefaultUserRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDefaultUserRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDefaultUserRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDefaultUserRuleResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LookupDefaultUserRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDefaultUserRuleResultOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o LookupDefaultUserRuleResultOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) []AddressPrefixItemResponse { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
}

func (o LookupDefaultUserRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDefaultUserRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultUserRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultUserRuleResultOutput{})
}
