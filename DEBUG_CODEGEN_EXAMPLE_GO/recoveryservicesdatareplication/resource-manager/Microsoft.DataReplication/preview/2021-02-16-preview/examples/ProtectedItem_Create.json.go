package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datareplication/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := datareplication.NewProtectedItem(ctx, "protectedItem", &datareplication.ProtectedItemArgs{
Properties: &datareplication.ProtectedItemModelPropertiesArgs{
CustomProperties: interface{}{
InstanceType: pulumi.String("ProtectedItemModelCustomProperties"),
},
PolicyName: pulumi.String("tjoeiynplt"),
ReplicationExtensionName: pulumi.String("jwxdo"),
},
ProtectedItemName: pulumi.String("d"),
ResourceGroupName: pulumi.String("rgrecoveryservicesdatareplication"),
VaultName: pulumi.String("4"),
})
if err != nil {
return err
}
return nil
})
}
