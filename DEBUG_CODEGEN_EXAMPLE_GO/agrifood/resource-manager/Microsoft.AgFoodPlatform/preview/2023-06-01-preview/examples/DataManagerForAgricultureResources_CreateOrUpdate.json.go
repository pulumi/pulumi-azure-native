package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/agfoodplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := agfoodplatform.NewDataManagerForAgricultureResource(ctx, "dataManagerForAgricultureResource", &agfoodplatform.DataManagerForAgricultureResourceArgs{
			DataManagerForAgricultureResourceName: pulumi.String("examples-farmbeatsResourceName"),
			Location:                              pulumi.String("eastus2"),
			ResourceGroupName:                     pulumi.String("examples-rg"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
				"key2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
