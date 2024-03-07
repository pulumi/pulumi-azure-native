package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridconnectivity/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridconnectivity.NewServiceConfiguration(ctx, "serviceConfiguration", &hybridconnectivity.ServiceConfigurationArgs{
			EndpointName:             pulumi.String("default"),
			Port:                     pulumi.Float64(22),
			ResourceUri:              pulumi.String("subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default"),
			ServiceConfigurationName: pulumi.String("SSH"),
			ServiceName:              pulumi.String("SSH"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
