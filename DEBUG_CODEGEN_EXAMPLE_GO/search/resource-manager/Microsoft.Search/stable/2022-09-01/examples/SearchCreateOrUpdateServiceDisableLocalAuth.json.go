package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/search/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := search.NewService(ctx, "service", &search.ServiceArgs{
			DisableLocalAuth:  pulumi.Bool(true),
			HostingMode:       search.HostingModeDefault,
			Location:          pulumi.String("westus"),
			PartitionCount:    pulumi.Int(1),
			ReplicaCount:      pulumi.Int(3),
			ResourceGroupName: pulumi.String("rg1"),
			SearchServiceName: pulumi.String("mysearchservice"),
			Sku: &search.SkuArgs{
				Name: search.SkuNameStandard,
			},
			Tags: pulumi.StringMap{
				"app-name": pulumi.String("My e-commerce app"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
