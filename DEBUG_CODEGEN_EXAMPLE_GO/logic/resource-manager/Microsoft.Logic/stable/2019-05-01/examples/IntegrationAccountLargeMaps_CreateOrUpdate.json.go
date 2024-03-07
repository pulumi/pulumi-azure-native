package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewIntegrationAccountMap(ctx, "integrationAccountMap", &logic.IntegrationAccountMapArgs{
			ContentType:            pulumi.String("application/xml"),
			IntegrationAccountName: pulumi.String("testIntegrationAccount"),
			Location:               pulumi.String("westus"),
			MapName:                pulumi.String("testMap"),
			MapType:                pulumi.String("Xslt"),
			Metadata:               nil,
			ResourceGroupName:      pulumi.String("testResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
