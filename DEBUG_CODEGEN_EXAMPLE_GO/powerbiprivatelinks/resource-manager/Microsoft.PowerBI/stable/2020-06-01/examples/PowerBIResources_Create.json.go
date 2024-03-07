package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/powerbi/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := powerbi.NewPowerBIResource(ctx, "powerBIResource", &powerbi.PowerBIResourceArgs{
			AzureResourceName: pulumi.String("azureResourceName"),
			Location:          pulumi.String("global"),
			ResourceGroupName: pulumi.String("resourceGroup"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
			TenantId: pulumi.String("ac2bc297-8a3e-46f3-972d-87c2b4ae6e2f"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
