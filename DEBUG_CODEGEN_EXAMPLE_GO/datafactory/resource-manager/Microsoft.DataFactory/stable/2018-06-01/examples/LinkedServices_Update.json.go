package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datafactory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datafactory.NewLinkedService(ctx, "linkedService", &datafactory.LinkedServiceArgs{
			FactoryName:       pulumi.String("exampleFactoryName"),
			LinkedServiceName: pulumi.String("exampleLinkedService"),
			Properties: datafactory.AzureStorageLinkedService{
				ConnectionString: map[string]interface{}{
					"type":  "SecureString",
					"value": "DefaultEndpointsProtocol=https;AccountName=examplestorageaccount;AccountKey=<storage key>",
				},
				Description: "Example description",
				Type:        "AzureStorage",
			},
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
