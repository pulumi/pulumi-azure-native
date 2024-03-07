package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewService(ctx, "service", &appplatform.ServiceArgs{
			Location: pulumi.String("eastus"),
			Properties: appplatform.ClusterResourcePropertiesResponse{
				MarketplaceResource: &appplatform.MarketplaceResourceArgs{
					Plan:      pulumi.String("tanzu-asc-ent-mtr"),
					Product:   pulumi.String("azure-spring-cloud-vmware-tanzu-2"),
					Publisher: pulumi.String("vmware-inc"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
			Sku: &appplatform.SkuArgs{
				Name: pulumi.String("E0"),
				Tier: pulumi.String("Enterprise"),
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
