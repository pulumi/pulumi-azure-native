package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := recoveryservices.NewReplicationFabric(ctx, "replicationFabric", &recoveryservices.ReplicationFabricArgs{
FabricName: pulumi.String("cloud1"),
Properties: &recoveryservices.FabricCreationInputPropertiesArgs{
CustomDetails: interface{}{
InstanceType: pulumi.String("FabricSpecificCreationInput"),
},
},
ResourceGroupName: pulumi.String("resourceGroupPS1"),
ResourceName: pulumi.String("vault1"),
})
if err != nil {
return err
}
return nil
})
}
