package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/providerhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := providerhub.NewDefaultRollout(ctx, "defaultRollout", &providerhub.DefaultRolloutArgs{
Properties: providerhub.DefaultRolloutResponseProperties{
Specification: interface{}{
Canary: &providerhub.DefaultRolloutSpecificationCanaryArgs{
SkipRegions: pulumi.StringArray{
pulumi.String("eastus2euap"),
},
},
ExpeditedRollout: &providerhub.DefaultRolloutSpecificationExpeditedRolloutArgs{
Enabled: pulumi.Bool(true),
},
RestOfTheWorldGroupTwo: &providerhub.DefaultRolloutSpecificationRestOfTheWorldGroupTwoArgs{
WaitDuration: pulumi.String("PT4H"),
},
},
},
ProviderNamespace: pulumi.String("Microsoft.Contoso"),
RolloutName: pulumi.String("2020week10"),
})
if err != nil {
return err
}
return nil
})
}
