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
				Properties: machinelearningservices.AmlComputeProperties{
					EnableNodePublicIp:          true,
					IsolatedNetwork:             false,
					OsType:                      "Windows",
					RemoteLoginPortPublicAccess: "NotSpecified",
					ScaleSettings: machinelearningservices.ScaleSettings{
						MaxNodeCount:                1,
						MinNodeCount:                0,
						NodeIdleTimeBeforeScaleDown: "PT5M",
					},
					VirtualMachineImage: machinelearningservices.VirtualMachineImage{
						Id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myImageGallery/images/myImageDefinition/versions/0.0.1",
					},
					VmPriority: "Dedicated",
					VmSize:     "STANDARD_NC6",
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
