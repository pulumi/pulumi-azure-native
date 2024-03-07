package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/importexport/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := importexport.NewJob(ctx, "job", &importexport.JobArgs{
			JobName:  pulumi.String("myExportJob"),
			Location: pulumi.String("West US"),
			Properties: importexport.JobDetailsResponse{
				BackupDriveManifest: pulumi.Bool(true),
				DiagnosticsPath:     pulumi.String("waimportexport"),
				Export: &importexport.ExportArgs{
					BlobPathPrefix: pulumi.StringArray{
						pulumi.String("/"),
					},
				},
				JobType:  pulumi.String("Export"),
				LogLevel: pulumi.String("Verbose"),
				ReturnAddress: &importexport.ReturnAddressArgs{
					City:            pulumi.String("Redmond"),
					CountryOrRegion: pulumi.String("USA"),
					Email:           pulumi.String("Test@contoso.com"),
					Phone:           pulumi.String("4250000000"),
					PostalCode:      pulumi.String("98007"),
					RecipientName:   pulumi.String("Test"),
					StateOrProvince: pulumi.String("wa"),
					StreetAddress1:  pulumi.String("Street1"),
					StreetAddress2:  pulumi.String("street2"),
				},
				ReturnShipping: &importexport.ReturnShippingArgs{
					CarrierAccountNumber: pulumi.String("989ffff"),
					CarrierName:          pulumi.String("FedEx"),
				},
				StorageAccountId: pulumi.String("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.ClassicStorage/storageAccounts/test"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
