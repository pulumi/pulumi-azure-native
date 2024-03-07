package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datareplication/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := datareplication.NewPolicy(ctx, "policy", &datareplication.PolicyArgs{
PolicyName: pulumi.String("fafqwc"),
Properties: &datareplication.PolicyModelPropertiesArgs{
CustomProperties: interface{}{
InstanceType: pulumi.String("PolicyModelCustomProperties"),
},
},
ResourceGroupName: pulumi.String("rgrecoveryservicesdatareplication"),
VaultName: pulumi.String("4"),
})
if err != nil {
return err
}
return nil
})
}
