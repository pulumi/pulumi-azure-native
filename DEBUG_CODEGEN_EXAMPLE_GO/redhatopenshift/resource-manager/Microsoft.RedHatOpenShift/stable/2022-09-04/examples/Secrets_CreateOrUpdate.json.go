package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/redhatopenshift/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redhatopenshift.NewSecret(ctx, "secret", &redhatopenshift.SecretArgs{
			ChildResourceName: pulumi.String("childResourceName"),
			ResourceGroupName: pulumi.String("resourceGroup"),
			ResourceName:      pulumi.String("resourceName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
