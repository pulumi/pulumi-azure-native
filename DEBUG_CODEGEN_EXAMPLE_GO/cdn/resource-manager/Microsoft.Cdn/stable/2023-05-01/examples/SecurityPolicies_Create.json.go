package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := cdn.NewSecurityPolicy(ctx, "securityPolicy", &cdn.SecurityPolicyArgs{
Parameters: interface{}{
Associations: cdn.SecurityPolicyWebApplicationFirewallAssociationArray{
interface{}{
Domains: cdn.ActivatedResourceReferenceArray{
&cdn.ActivatedResourceReferenceArgs{
Id: pulumi.String("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain1"),
},
&cdn.ActivatedResourceReferenceArgs{
Id: pulumi.String("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain2"),
},
},
PatternsToMatch: pulumi.StringArray{
pulumi.String("/*"),
},
},
},
Type: pulumi.String("WebApplicationFirewall"),
WafPolicy: &cdn.ResourceReferenceArgs{
Id: pulumi.String("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoorwebapplicationfirewallpolicies/wafTest"),
},
},
ProfileName: pulumi.String("profile1"),
ResourceGroupName: pulumi.String("RG"),
SecurityPolicyName: pulumi.String("securityPolicy1"),
})
if err != nil {
return err
}
return nil
})
}
