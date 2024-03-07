package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/notificationhubs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := notificationhubs.NewNamespace(ctx, "namespace", &notificationhubs.NamespaceArgs{
Location: pulumi.String("South Central US"),
NamespaceName: pulumi.String("nh-sdk-ns"),
Properties: notificationhubs.NamespacePropertiesResponse{
NetworkAcls: interface{}{
IpRules: notificationhubs.IpRuleArray{
&notificationhubs.IpRuleArgs{
IpMask: pulumi.String("185.48.100.00/24"),
Rights: pulumi.StringArray{
pulumi.String("Manage"),
pulumi.String("Send"),
pulumi.String("Listen"),
},
},
},
PublicNetworkRule: &notificationhubs.PublicInternetAuthorizationRuleArgs{
Rights: pulumi.StringArray{
pulumi.String("Listen"),
},
},
},
ZoneRedundancy: pulumi.String("Enabled"),
},
ResourceGroupName: pulumi.String("5ktrial"),
Sku: &notificationhubs.SkuArgs{
Name: pulumi.String("Standard"),
Tier: pulumi.String("Standard"),
},
Tags: pulumi.StringMap{
"tag1": pulumi.String("value1"),
"tag2": pulumi.String("value2"),
},
})
if err != nil {
return err
}
return nil
})
}
