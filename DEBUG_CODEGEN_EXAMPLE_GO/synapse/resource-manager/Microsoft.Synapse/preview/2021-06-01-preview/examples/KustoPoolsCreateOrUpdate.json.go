package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewKustoPool(ctx, "kustoPool", &synapse.KustoPoolArgs{
			EnablePurge:           pulumi.Bool(true),
			EnableStreamingIngest: pulumi.Bool(true),
			KustoPoolName:         pulumi.String("kustoclusterrptest4"),
			Location:              pulumi.String("westus"),
			ResourceGroupName:     pulumi.String("kustorptest"),
			Sku: &synapse.AzureSkuArgs{
				Capacity: pulumi.Int(2),
				Name:     pulumi.String("Storage optimized"),
				Size:     pulumi.String("Medium"),
			},
			WorkspaceName: pulumi.String("synapseWorkspaceName"),
			WorkspaceUID:  pulumi.String("11111111-2222-3333-444444444444"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
