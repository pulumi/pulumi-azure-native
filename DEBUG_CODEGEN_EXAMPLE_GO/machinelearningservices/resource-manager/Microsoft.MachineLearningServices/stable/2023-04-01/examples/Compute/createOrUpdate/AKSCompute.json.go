package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewCompute(ctx, "compute", &machinelearningservices.ComputeArgs{
			ComputeName: pulumi.String("compute123"),
			Location:    pulumi.String("eastus"),
			Properties: machinelearningservices.AKS{
				ComputeType: "AKS",
				Description: "some compute",
				Properties: machinelearningservices.AKSSchemaProperties{
					AgentCount: 4,
				},
				ResourceId: "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute123-56826-c9b00420020b2",
			},
			ResourceGroupName: pulumi.String("testrg123"),
			WorkspaceName:     pulumi.String("workspaces123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
