package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datareplication/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := datareplication.NewReplicationExtension(ctx, "replicationExtension", &datareplication.ReplicationExtensionArgs{
Properties: &datareplication.ReplicationExtensionModelPropertiesArgs{
CustomProperties: interface{}{
InstanceType: pulumi.String("ReplicationExtensionModelCustomProperties"),
},
},
ReplicationExtensionName: pulumi.String("g16yjJ"),
ResourceGroupName: pulumi.String("rgrecoveryservicesdatareplication"),
VaultName: pulumi.String("4"),
})
if err != nil {
return err
}
return nil
})
}
