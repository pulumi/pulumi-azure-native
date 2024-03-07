package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := network.NewPolicy(ctx, "policy", &network.PolicyArgs{
CustomRules: network.CustomRuleListResponse{
Rules: network.CustomRuleArray{
interface{}{
Action: pulumi.String("Block"),
MatchConditions: network.FrontDoorMatchConditionArray{
&network.FrontDoorMatchConditionArgs{
MatchValue: pulumi.StringArray{
pulumi.String("192.168.1.0/24"),
pulumi.String("10.0.0.0/24"),
},
MatchVariable: pulumi.String("RemoteAddr"),
Operator: pulumi.String("IPMatch"),
},
},
Name: pulumi.String("Rule1"),
Priority: pulumi.Int(1),
RateLimitThreshold: pulumi.Int(1000),
RuleType: pulumi.String("RateLimitRule"),
},
interface{}{
Action: pulumi.String("Block"),
MatchConditions: network.FrontDoorMatchConditionArray{
&network.FrontDoorMatchConditionArgs{
MatchValue: pulumi.StringArray{
pulumi.String("CH"),
},
MatchVariable: pulumi.String("RemoteAddr"),
Operator: pulumi.String("GeoMatch"),
},
&network.FrontDoorMatchConditionArgs{
MatchValue: pulumi.StringArray{
pulumi.String("windows"),
},
MatchVariable: pulumi.String("RequestHeader"),
Operator: pulumi.String("Contains"),
Selector: pulumi.String("UserAgent"),
Transforms: pulumi.StringArray{
pulumi.String("Lowercase"),
},
},
},
Name: pulumi.String("Rule2"),
Priority: pulumi.Int(2),
RuleType: pulumi.String("MatchRule"),
},
},
},
Location: pulumi.String("WestUs"),
ManagedRules: network.ManagedRuleSetListResponse{
ManagedRuleSets: network.FrontDoorManagedRuleSetArray{
interface{}{
Exclusions: network.ManagedRuleExclusionArray{
&network.ManagedRuleExclusionArgs{
MatchVariable: pulumi.String("RequestHeaderNames"),
Selector: pulumi.String("User-Agent"),
SelectorMatchOperator: pulumi.String("Equals"),
},
},
RuleGroupOverrides: network.FrontDoorManagedRuleGroupOverrideArray{
interface{}{
Exclusions: network.ManagedRuleExclusionArray{
&network.ManagedRuleExclusionArgs{
MatchVariable: pulumi.String("RequestCookieNames"),
Selector: pulumi.String("token"),
SelectorMatchOperator: pulumi.String("StartsWith"),
},
},
RuleGroupName: pulumi.String("SQLI"),
Rules: network.FrontDoorManagedRuleOverrideArray{
interface{}{
Action: pulumi.String("Redirect"),
EnabledState: pulumi.String("Enabled"),
Exclusions: network.ManagedRuleExclusionArray{
&network.ManagedRuleExclusionArgs{
MatchVariable: pulumi.String("QueryStringArgNames"),
Selector: pulumi.String("query"),
SelectorMatchOperator: pulumi.String("Equals"),
},
},
RuleId: pulumi.String("942100"),
},
&network.FrontDoorManagedRuleOverrideArgs{
EnabledState: pulumi.String("Disabled"),
RuleId: pulumi.String("942110"),
},
},
},
},
RuleSetAction: pulumi.String("Block"),
RuleSetType: pulumi.String("DefaultRuleSet"),
RuleSetVersion: pulumi.String("1.0"),
},
},
},
PolicyName: pulumi.String("Policy1"),
PolicySettings: &network.FrontDoorPolicySettingsArgs{
CustomBlockResponseBody: pulumi.String("PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg=="),
CustomBlockResponseStatusCode: pulumi.Int(429),
EnabledState: pulumi.String("Enabled"),
Mode: pulumi.String("Prevention"),
RedirectUrl: pulumi.String("http://www.bing.com"),
RequestBodyCheck: pulumi.String("Disabled"),
},
ResourceGroupName: pulumi.String("rg1"),
Sku: &network.SkuArgs{
Name: pulumi.String("Classic_AzureFrontDoor"),
},
})
if err != nil {
return err
}
return nil
})
}
