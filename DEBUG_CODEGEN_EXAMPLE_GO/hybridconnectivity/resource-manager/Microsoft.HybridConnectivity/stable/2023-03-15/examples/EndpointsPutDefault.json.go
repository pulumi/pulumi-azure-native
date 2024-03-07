package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridconnectivity/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridconnectivity.NewEndpoint(ctx, "endpoint", &hybridconnectivity.EndpointArgs{
			EndpointName: pulumi.String("default"),
			Properties: &hybridconnectivity.EndpointPropertiesArgs{
				Type: pulumi.String("default"),
			},
			ResourceUri: pulumi.String("subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
