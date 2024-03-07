package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/providerhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := providerhub.NewSkusNestedResourceTypeFirst(ctx, "skusNestedResourceTypeFirst", &providerhub.SkusNestedResourceTypeFirstArgs{
NestedResourceTypeFirst: pulumi.String("nestedResourceTypeFirst"),
Properties: providerhub.SkuResourceResponseProperties{
SkuSettings: providerhub.SkuSettingArray{
&providerhub.SkuSettingArgs{
Kind: pulumi.String("Standard"),
Name: pulumi.String("freeSku"),
Tier: pulumi.String("Tier1"),
},
interface{}{
Costs: providerhub.SkuCostArray{
&providerhub.SkuCostArgs{
MeterId: pulumi.String("xxx"),
},
},
Kind: pulumi.String("Premium"),
Name: pulumi.String("premiumSku"),
Tier: pulumi.String("Tier2"),
},
},
},
ProviderNamespace: pulumi.String("Microsoft.Contoso"),
ResourceType: pulumi.String("testResourceType"),
Sku: pulumi.String("testSku"),
})
if err != nil {
return err
}
return nil
})
}
