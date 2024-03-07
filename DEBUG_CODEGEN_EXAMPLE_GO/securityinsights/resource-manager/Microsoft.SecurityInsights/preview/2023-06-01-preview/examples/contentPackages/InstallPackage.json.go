package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewContentPackage(ctx, "contentPackage", &securityinsights.ContentPackageArgs{
			ContentId:         pulumi.String("str.azure-sentinel-solution-str"),
			ContentKind:       pulumi.String("Solution"),
			DisplayName:       pulumi.String("str"),
			PackageId:         pulumi.String("str.azure-sentinel-solution-str"),
			ResourceGroupName: pulumi.String("myRg"),
			Version:           pulumi.String("2.0.0"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
