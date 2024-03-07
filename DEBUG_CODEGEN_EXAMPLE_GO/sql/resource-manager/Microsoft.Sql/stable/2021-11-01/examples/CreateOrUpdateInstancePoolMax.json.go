package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewInstancePool(ctx, "instancePool", &sql.InstancePoolArgs{
			InstancePoolName:  pulumi.String("testIP"),
			LicenseType:       pulumi.String("LicenseIncluded"),
			Location:          pulumi.String("japaneast"),
			ResourceGroupName: pulumi.String("group1"),
			Sku: &sql.SkuArgs{
				Family: pulumi.String("Gen5"),
				Name:   pulumi.String("GP_Gen5"),
				Tier:   pulumi.String("GeneralPurpose"),
			},
			SubnetId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet1"),
			Tags: pulumi.StringMap{
				"a": pulumi.String("b"),
			},
			VCores: pulumi.Int(8),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
