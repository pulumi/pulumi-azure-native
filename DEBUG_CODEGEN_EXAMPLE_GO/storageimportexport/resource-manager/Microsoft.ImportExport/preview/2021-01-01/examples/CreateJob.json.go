package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/importexport/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := importexport.NewJob(ctx, "job", &importexport.JobArgs{
			JobName:  pulumi.String("myJob"),
			Location: pulumi.String("West US"),
			Properties: importexport.JobDetailsResponse{
				BackupDriveManifest: pulumi.Bool(true),
				DiagnosticsPath:     pulumi.String("waimportexport"),
				DriveList: importexport.DriveStatusArray{
					&importexport.DriveStatusArgs{
						BitLockerKey:    pulumi.String("238810-662376-448998-450120-652806-203390-606320-483076"),
						DriveHeaderHash: pulumi.String("0:1048576:FB6B6ED500D49DA6E0D723C98D42C657F2881CC13357C28DCECA6A524F1292501571A321238540E621AB5BD9C9A32637615919A75593E6CB5C1515DAE341CABF;135266304:143360:C957A189AFC38C4E80731252301EB91427CE55E61448FA3C73C6FDDE70ABBC197947EC8D0249A2C639BB10B95957D5820A4BE8DFBBF76FFFA688AE5CE0D42EC3"),
						DriveId:         pulumi.String("9CA995BB"),
						ManifestFile:    pulumi.String("\\8a0c23f7-14b7-470a-9633-fcd46590a1bc.manifest"),
						ManifestHash:    pulumi.String("4228EC5D8E048CB9B515338C789314BE8D0B2FDBC7C7A0308E1C826242CDE74E"),
					},
				},
				JobType:  pulumi.String("Import"),
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
