package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewBuildServiceAgentPool(ctx, "buildServiceAgentPool", &appplatform.BuildServiceAgentPoolArgs{
			AgentPoolName:    pulumi.String("default"),
			BuildServiceName: pulumi.String("default"),
			Properties: &appplatform.BuildServiceAgentPoolPropertiesArgs{
				PoolSize: &appplatform.BuildServiceAgentPoolSizePropertiesArgs{
					Name: pulumi.String("S3"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
