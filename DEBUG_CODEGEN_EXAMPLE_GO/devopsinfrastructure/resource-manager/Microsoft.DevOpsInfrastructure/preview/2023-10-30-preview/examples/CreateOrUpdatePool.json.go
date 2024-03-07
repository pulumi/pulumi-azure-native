package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devopsinfrastructure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devopsinfrastructure.NewPool(ctx, "pool", &devopsinfrastructure.PoolArgs{
			AgentProfile: devopsinfrastructure.StatelessAgentProfile{
				Kind: "Stateless",
			},
			DevCenterProjectResourceId: pulumi.String("/subscriptions/222e81d0-cf38-4dab-baa5-289bf16baaa4/resourceGroups/rg-1es-devcenter/providers/Microsoft.DevCenter/projects/1ES"),
			FabricProfile: &devopsinfrastructure.VmssFabricProfileArgs{
				Images: devopsinfrastructure.PoolImageArray{
					&devopsinfrastructure.PoolImageArgs{
						ResourceId: pulumi.String("/MicrosoftWindowsServer/WindowsServer/2019-Datacenter/latest"),
					},
				},
				Kind: pulumi.String("Vmss"),
				Sku: &devopsinfrastructure.DevOpsAzureSkuArgs{
					Name: pulumi.String("Standard_D4ads_v5"),
				},
			},
			Location:           pulumi.String("eastus"),
			MaximumConcurrency: pulumi.Int(10),
			OrganizationProfile: &devopsinfrastructure.AzureDevOpsOrganizationProfileArgs{
				Kind: pulumi.String("AzureDevOps"),
				Organizations: devopsinfrastructure.OrganizationArray{
					&devopsinfrastructure.OrganizationArgs{
						Url: pulumi.String("https://mseng.visualstudio.com"),
					},
				},
			},
			PoolName:          pulumi.String("pool"),
			ProvisioningState: pulumi.String("Succeeded"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
