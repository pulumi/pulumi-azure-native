package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewReplicationNetworkMapping(ctx, "replicationNetworkMapping", &recoveryservices.ReplicationNetworkMappingArgs{
			FabricName:         pulumi.String("b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac"),
			NetworkMappingName: pulumi.String("corpe2amap"),
			NetworkName:        pulumi.String("e2267b5c-2650-49bd-ab3f-d66aae694c06"),
			Properties: &recoveryservices.CreateNetworkMappingInputPropertiesArgs{
				FabricSpecificDetails: recoveryservices.VmmToAzureCreateNetworkMappingInput{
					InstanceType: "VmmToAzure",
				},
				RecoveryFabricName: pulumi.String("Microsoft Azure"),
				RecoveryNetworkId:  pulumi.String("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai"),
			},
			ResourceGroupName: pulumi.String("srcBvte2a14C27"),
			ResourceName:      pulumi.String("srce2avaultbvtaC27"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
