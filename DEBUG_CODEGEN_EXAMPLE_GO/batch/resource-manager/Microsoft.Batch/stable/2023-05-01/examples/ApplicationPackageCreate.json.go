package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/batch/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.NewApplicationPackage(ctx, "applicationPackage", &batch.ApplicationPackageArgs{
			AccountName:       pulumi.String("sampleacct"),
			ApplicationName:   pulumi.String("app1"),
			ResourceGroupName: pulumi.String("default-azurebatch-japaneast"),
			VersionName:       pulumi.String("1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
