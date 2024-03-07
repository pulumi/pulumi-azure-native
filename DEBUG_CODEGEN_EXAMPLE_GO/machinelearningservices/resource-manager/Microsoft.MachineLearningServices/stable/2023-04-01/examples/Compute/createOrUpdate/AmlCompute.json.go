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
			Properties: machinelearningservices.AmlCompute{
				ComputeType: "AmlCompute",
				Description: "some compute",
				Properties: machinelearningservices.AmlComputeProperties{
					ScaleSettings: machinelearningservices.ScaleSettings{
						MaxNodeCount:                4,
						MinNodeCount:                4,
						NodeIdleTimeBeforeScaleDown: "PT5M",
					},
				},
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
