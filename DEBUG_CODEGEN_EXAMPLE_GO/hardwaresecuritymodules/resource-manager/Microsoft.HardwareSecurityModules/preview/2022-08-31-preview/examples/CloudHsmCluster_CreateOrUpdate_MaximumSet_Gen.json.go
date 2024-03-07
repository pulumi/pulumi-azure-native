package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hardwaresecuritymodules/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hardwaresecuritymodules.NewCloudHsmCluster(ctx, "cloudHsmCluster", &hardwaresecuritymodules.CloudHsmClusterArgs{
			CloudHsmClusterName: pulumi.String("chsm1"),
			Location:            pulumi.String("eastus2"),
			ResourceGroupName:   pulumi.String("rgcloudhsm"),
			SecurityDomain: &hardwaresecuritymodules.CloudHsmClusterSecurityDomainPropertiesArgs{
				FipsState: pulumi.Int(2),
			},
			Sku: &hardwaresecuritymodules.CloudHsmClusterSkuArgs{
				Family: pulumi.String("B"),
				Name:   hardwaresecuritymodules.CloudHsmClusterSkuName_Standard_B1,
			},
			Tags: pulumi.StringMap{
				"Dept":        pulumi.String("hsm"),
				"Environment": pulumi.String("dogfood"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
