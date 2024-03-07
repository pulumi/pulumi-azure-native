package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewImportCollectorsOperation(ctx, "importCollectorsOperation", &migrate.ImportCollectorsOperationArgs{
			DiscoverySiteId:     pulumi.String("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourcegroups/ayagrawRG/providers/microsoft.offazure/importsites/actualSEA37d4importSite"),
			ImportCollectorName: pulumi.String("importCollectore7d5"),
			ProjectName:         pulumi.String("app18700project"),
			ProvisioningState:   pulumi.String("Succeeded"),
			ResourceGroupName:   pulumi.String("ayagrawRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
