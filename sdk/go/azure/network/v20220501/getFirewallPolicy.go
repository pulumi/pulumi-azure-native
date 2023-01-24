


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallPolicy(ctx *pulumi.Context, args *LookupFirewallPolicyArgs, opts ...pulumi.InvokeOption) (*LookupFirewallPolicyResult, error) {
	var rv LookupFirewallPolicyResult
	err := ctx.Invoke("azure-native:network/v20220501:getFirewallPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallPolicyArgs struct {
	Expand             *string `pulumi:"expand"`
	FirewallPolicyName string  `pulumi:"firewallPolicyName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupFirewallPolicyResult struct {
	BasePolicy           *SubResourceResponse                        `pulumi:"basePolicy"`
	ChildPolicies        []SubResourceResponse                       `pulumi:"childPolicies"`
	DnsSettings          *DnsSettingsResponse                        `pulumi:"dnsSettings"`
	Etag                 string                                      `pulumi:"etag"`
	ExplicitProxy        *ExplicitProxyResponse                      `pulumi:"explicitProxy"`
	Firewalls            []SubResourceResponse                       `pulumi:"firewalls"`
	Id                   *string                                     `pulumi:"id"`
	Identity             *ManagedServiceIdentityResponse             `pulumi:"identity"`
	Insights             *FirewallPolicyInsightsResponse             `pulumi:"insights"`
	IntrusionDetection   *FirewallPolicyIntrusionDetectionResponse   `pulumi:"intrusionDetection"`
	Location             *string                                     `pulumi:"location"`
	Name                 string                                      `pulumi:"name"`
	ProvisioningState    string                                      `pulumi:"provisioningState"`
	RuleCollectionGroups []SubResourceResponse                       `pulumi:"ruleCollectionGroups"`
	Sku                  *FirewallPolicySkuResponse                  `pulumi:"sku"`
	Snat                 *FirewallPolicySNATResponse                 `pulumi:"snat"`
	Sql                  *FirewallPolicySQLResponse                  `pulumi:"sql"`
	Tags                 map[string]string                           `pulumi:"tags"`
	ThreatIntelMode      *string                                     `pulumi:"threatIntelMode"`
	ThreatIntelWhitelist *FirewallPolicyThreatIntelWhitelistResponse `pulumi:"threatIntelWhitelist"`
	TransportSecurity    *FirewallPolicyTransportSecurityResponse    `pulumi:"transportSecurity"`
	Type                 string                                      `pulumi:"type"`
}

func LookupFirewallPolicyOutput(ctx *pulumi.Context, args LookupFirewallPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallPolicyResult, error) {
			args := v.(LookupFirewallPolicyArgs)
			r, err := LookupFirewallPolicy(ctx, &args, opts...)
			var s LookupFirewallPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallPolicyResultOutput)
}

type LookupFirewallPolicyOutputArgs struct {
	Expand             pulumi.StringPtrInput `pulumi:"expand"`
	FirewallPolicyName pulumi.StringInput    `pulumi:"firewallPolicyName"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupFirewallPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyArgs)(nil)).Elem()
}


type LookupFirewallPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyResult)(nil)).Elem()
}

func (o LookupFirewallPolicyResultOutput) ToLookupFirewallPolicyResultOutput() LookupFirewallPolicyResultOutput {
	return o
}

func (o LookupFirewallPolicyResultOutput) ToLookupFirewallPolicyResultOutputWithContext(ctx context.Context) LookupFirewallPolicyResultOutput {
	return o
}

func (o LookupFirewallPolicyResultOutput) BasePolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *SubResourceResponse { return v.BasePolicy }).(SubResourceResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) ChildPolicies() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) []SubResourceResponse { return v.ChildPolicies }).(SubResourceResponseArrayOutput)
}

func (o LookupFirewallPolicyResultOutput) DnsSettings() DnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *DnsSettingsResponse { return v.DnsSettings }).(DnsSettingsResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicyResultOutput) ExplicitProxy() ExplicitProxyResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *ExplicitProxyResponse { return v.ExplicitProxy }).(ExplicitProxyResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Firewalls() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) []SubResourceResponse { return v.Firewalls }).(SubResourceResponseArrayOutput)
}

func (o LookupFirewallPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Insights() FirewallPolicyInsightsResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *FirewallPolicyInsightsResponse { return v.Insights }).(FirewallPolicyInsightsResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) IntrusionDetection() FirewallPolicyIntrusionDetectionResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *FirewallPolicyIntrusionDetectionResponse {
		return v.IntrusionDetection
	}).(FirewallPolicyIntrusionDetectionResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicyResultOutput) RuleCollectionGroups() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) []SubResourceResponse { return v.RuleCollectionGroups }).(SubResourceResponseArrayOutput)
}

func (o LookupFirewallPolicyResultOutput) Sku() FirewallPolicySkuResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *FirewallPolicySkuResponse { return v.Sku }).(FirewallPolicySkuResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Snat() FirewallPolicySNATResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *FirewallPolicySNATResponse { return v.Snat }).(FirewallPolicySNATResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Sql() FirewallPolicySQLResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *FirewallPolicySQLResponse { return v.Sql }).(FirewallPolicySQLResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFirewallPolicyResultOutput) ThreatIntelMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *string { return v.ThreatIntelMode }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicyResultOutput) ThreatIntelWhitelist() FirewallPolicyThreatIntelWhitelistResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *FirewallPolicyThreatIntelWhitelistResponse {
		return v.ThreatIntelWhitelist
	}).(FirewallPolicyThreatIntelWhitelistResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) TransportSecurity() FirewallPolicyTransportSecurityResponsePtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) *FirewallPolicyTransportSecurityResponse {
		return v.TransportSecurity
	}).(FirewallPolicyTransportSecurityResponsePtrOutput)
}

func (o LookupFirewallPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallPolicyResultOutput{})
}
