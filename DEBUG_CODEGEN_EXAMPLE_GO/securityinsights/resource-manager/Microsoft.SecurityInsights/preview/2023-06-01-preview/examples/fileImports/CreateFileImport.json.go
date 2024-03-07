package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewFileImport(ctx, "fileImport", &securityinsights.FileImportArgs{
			ContentType:  pulumi.String("StixIndicator"),
			FileImportId: pulumi.String("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
			ImportFile: &securityinsights.FileMetadataArgs{
				FileFormat: pulumi.String("JSON"),
				FileName:   pulumi.String("myFile.json"),
				FileSize:   pulumi.Int(4653),
			},
			IngestionMode:     pulumi.String("IngestAnyValidRecords"),
			ResourceGroupName: pulumi.String("myRg"),
			Source:            pulumi.String("mySource"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
