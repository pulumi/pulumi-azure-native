


package v20181201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityRule(ctx *pulumi.Context, args *LookupSecurityRuleArgs, opts ...pulumi.InvokeOption) (*LookupSecurityRuleResult, error) {
	var rv LookupSecurityRuleResult
	err := ctx.Invoke("azure-native:network/v20181201:getSecurityRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityRuleArgs struct {
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
	SecurityRuleName         string `pulumi:"securityRuleName"`
}


type LookupSecurityRuleResult struct {
	Access                               string                             `pulumi:"access"`
	Description                          *string                            `pulumi:"description"`
	DestinationAddressPrefix             *string                            `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes           []string                           `pulumi:"destinationAddressPrefixes"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupResponse `pulumi:"destinationApplicationSecurityGroups"`
	DestinationPortRange                 *string                            `pulumi:"destinationPortRange"`
	DestinationPortRanges                []string                           `pulumi:"destinationPortRanges"`
	Direction                            string                             `pulumi:"direction"`
	Etag                                 *string                            `pulumi:"etag"`
	Id                                   *string                            `pulumi:"id"`
	Name                                 *string                            `pulumi:"name"`
	Priority                             *int                               `pulumi:"priority"`
	Protocol                             string                             `pulumi:"protocol"`
	ProvisioningState                    *string                            `pulumi:"provisioningState"`
	SourceAddressPrefix                  *string                            `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes                []string                           `pulumi:"sourceAddressPrefixes"`
	SourceApplicationSecurityGroups      []ApplicationSecurityGroupResponse `pulumi:"sourceApplicationSecurityGroups"`
	SourcePortRange                      *string                            `pulumi:"sourcePortRange"`
	SourcePortRanges                     []string                           `pulumi:"sourcePortRanges"`
}

func LookupSecurityRuleOutput(ctx *pulumi.Context, args LookupSecurityRuleOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityRuleResult, error) {
			args := v.(LookupSecurityRuleArgs)
			r, err := LookupSecurityRule(ctx, &args, opts...)
			var s LookupSecurityRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityRuleResultOutput)
}

type LookupSecurityRuleOutputArgs struct {
	NetworkSecurityGroupName pulumi.StringInput `pulumi:"networkSecurityGroupName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
	SecurityRuleName         pulumi.StringInput `pulumi:"securityRuleName"`
}

func (LookupSecurityRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityRuleArgs)(nil)).Elem()
}


type LookupSecurityRuleResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityRuleResult)(nil)).Elem()
}

func (o LookupSecurityRuleResultOutput) ToLookupSecurityRuleResultOutput() LookupSecurityRuleResultOutput {
	return o
}

func (o LookupSecurityRuleResultOutput) ToLookupSecurityRuleResultOutputWithContext(ctx context.Context) LookupSecurityRuleResultOutput {
	return o
}

func (o LookupSecurityRuleResultOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.Access }).(pulumi.StringOutput)
}

func (o LookupSecurityRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityRuleResultOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityRuleResultOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []string { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o LookupSecurityRuleResultOutput) DestinationApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []ApplicationSecurityGroupResponse {
		return v.DestinationApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

func (o LookupSecurityRuleResultOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityRuleResultOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o LookupSecurityRuleResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.Direction }).(pulumi.StringOutput)
}

func (o LookupSecurityRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityRuleResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupSecurityRuleResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LookupSecurityRuleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityRuleResultOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityRuleResultOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []string { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o LookupSecurityRuleResultOutput) SourceApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []ApplicationSecurityGroupResponse {
		return v.SourceApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

func (o LookupSecurityRuleResultOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityRuleResultOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityRuleResultOutput{})
}
