package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewImportCollector(ctx, "importCollector", &migrate.ImportCollectorArgs{
			ImportCollectorName: pulumi.String("importCollector2952"),
			ProjectName:         pulumi.String("rajoshCCY9671project"),
			ResourceGroupName:   pulumi.String("markusavstestrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
