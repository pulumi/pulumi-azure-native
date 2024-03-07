package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewManagedCluster(ctx, "managedCluster", &servicefabric.ManagedClusterArgs{
			AdminPassword:         pulumi.String("{vm-password}"),
			AdminUserName:         pulumi.String("vmadmin"),
			ClusterName:           pulumi.String("myCluster"),
			ClusterUpgradeCadence: pulumi.String("Wave1"),
			ClusterUpgradeMode:    pulumi.String("Automatic"),
			DnsName:               pulumi.String("myCluster"),
			FabricSettings: []servicefabric.SettingsSectionDescriptionArgs{
				{
					Name: pulumi.String("ManagedIdentityTokenService"),
					Parameters: servicefabric.SettingsParameterDescriptionArray{
						{
							Name:  pulumi.String("IsEnabled"),
							Value: pulumi.String("true"),
						},
					},
				},
			},
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("resRg"),
			Sku: &servicefabric.SkuArgs{
				Name: pulumi.String("Basic"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
