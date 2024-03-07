package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := recoveryservices.NewReplicationProtectionContainerMapping(ctx, "replicationProtectionContainerMapping", &recoveryservices.ReplicationProtectionContainerMappingArgs{
FabricName: pulumi.String("cloud1"),
MappingName: pulumi.String("cloud1protectionprofile1"),
Properties: recoveryservices.ProtectionContainerMappingPropertiesResponse{
PolicyId: pulumi.String("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
ProviderSpecificInput: interface{}{
InstanceType: pulumi.String("ReplicationProviderSpecificContainerMappingInput"),
},
TargetProtectionContainerId: pulumi.String("Microsoft Azure"),
},
ProtectionContainerName: pulumi.String("cloud_6d224fc6-f326-5d35-96de-fbf51efb3179"),
ResourceGroupName: pulumi.String("resourceGroupPS1"),
ResourceName: pulumi.String("vault1"),
})
if err != nil {
return err
}
return nil
})
}
