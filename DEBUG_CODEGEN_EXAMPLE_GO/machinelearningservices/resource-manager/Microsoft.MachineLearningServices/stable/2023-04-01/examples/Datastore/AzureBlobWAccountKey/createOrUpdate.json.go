package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewDatastore(ctx, "datastore", &machinelearningservices.DatastoreArgs{
			DatastoreProperties: machinelearningservices.AzureBlobDatastore{
				AccountName:   "string",
				ContainerName: "string",
				Credentials: machinelearningservices.AccountKeyDatastoreCredentials{
					CredentialsType: "AccountKey",
					Secrets: machinelearningservices.AccountKeyDatastoreSecrets{
						Key:         "string",
						SecretsType: "AccountKey",
					},
				},
				DatastoreType: "AzureBlob",
				Description:   "string",
				Endpoint:      "core.windows.net",
				Protocol:      "https",
				Tags: map[string]interface{}{
					"string": "string",
				},
			},
			Name:              pulumi.String("string"),
			ResourceGroupName: pulumi.String("test-rg"),
			SkipValidation:    pulumi.Bool(false),
			WorkspaceName:     pulumi.String("my-aml-workspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
