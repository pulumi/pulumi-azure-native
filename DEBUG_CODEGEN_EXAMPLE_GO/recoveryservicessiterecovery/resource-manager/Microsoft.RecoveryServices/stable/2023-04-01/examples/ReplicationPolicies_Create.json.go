package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := recoveryservices.NewReplicationPolicy(ctx, "replicationPolicy", &recoveryservices.ReplicationPolicyArgs{
PolicyName: pulumi.String("protectionprofile1"),
Properties: interface{}{
ProviderSpecificInput: recoveryservices.HyperVReplicaAzurePolicyInput{
InstanceType: "HyperVReplicaAzure",
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
