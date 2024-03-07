package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewImportSitesController(ctx, "importSitesController", &offazure.ImportSitesControllerArgs{
			Location:          pulumi.String("woctgvdufvkzvjcirjpf"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("IHPDPK-1-"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
