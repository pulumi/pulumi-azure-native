package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewCluster(ctx, "cluster", &kusto.ClusterArgs{
			AllowedIpRangeList: pulumi.StringArray{
				pulumi.String("0.0.0.0/0"),
			},
			ClusterName:            pulumi.String("kustoCluster"),
			EnableAutoStop:         pulumi.Bool(true),
			EnableDoubleEncryption: pulumi.Bool(false),
			EnablePurge:            pulumi.Bool(true),
			EnableStreamingIngest:  pulumi.Bool(true),
			Identity: &kusto.IdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			LanguageExtensions: kusto.LanguageExtensionsListResponse{
				Value: kusto.LanguageExtensionArray{
					&kusto.LanguageExtensionArgs{
						LanguageExtensionImageName: pulumi.String("Python3_10_8"),
						LanguageExtensionName:      pulumi.String("PYTHON"),
					},
					&kusto.LanguageExtensionArgs{
						LanguageExtensionImageName: pulumi.String("R"),
						LanguageExtensionName:      pulumi.String("R"),
					},
				},
			},
			Location:            pulumi.String("westus"),
			PublicIPType:        pulumi.String("DualStack"),
			PublicNetworkAccess: pulumi.String("Enabled"),
			ResourceGroupName:   pulumi.String("kustorptest"),
			Sku: &kusto.AzureSkuArgs{
				Capacity: pulumi.Int(2),
				Name:     pulumi.String("Standard_L16as_v3"),
				Tier:     pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
