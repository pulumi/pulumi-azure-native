package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/agfoodplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := agfoodplatform.NewExtension(ctx, "extension", &agfoodplatform.ExtensionArgs{
			DataManagerForAgricultureResourceName: pulumi.String("examples-dataManagerForAgricultureResourceName"),
			ExtensionId:                           pulumi.String("provider.extension"),
			ResourceGroupName:                     pulumi.String("examples-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
