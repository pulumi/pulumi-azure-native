package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databox/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databox.NewJob(ctx, "job", &databox.JobArgs{
			Details: databox.DataBoxJobDetails{
				ContactDetails: databox.ContactDetails{
					ContactName: "XXXX XXXX",
					EmailList: []string{
						"xxxx@xxxx.xxx",
					},
					Phone:          "0000000000",
					PhoneExtension: "",
				},
				DataImportDetails: []databox.DataImportDetails{
					{
						AccountDetails: {
							DataAccountType:  "StorageAccount",
							SharePassword:    "<sharePassword>",
							StorageAccountId: "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName",
						},
					},
				},
				DevicePassword: "<devicePassword>",
				JobDetailsType: "DataBox",
				ShippingAddress: databox.ShippingAddress{
					AddressType:     "Commercial",
					City:            "XXXX XXXX",
					CompanyName:     "XXXX XXXX",
					Country:         "XX",
					PostalCode:      "00000",
					StateOrProvince: "XX",
					StreetAddress1:  "XXXX XXXX",
					StreetAddress2:  "XXXX XXXX",
				},
			},
			JobName:           pulumi.String("TestJobName1"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("YourResourceGroupName"),
			Sku: &databox.SkuArgs{
				Name: pulumi.String("DataBox"),
			},
			TransferType: pulumi.String("ImportToAzure"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
