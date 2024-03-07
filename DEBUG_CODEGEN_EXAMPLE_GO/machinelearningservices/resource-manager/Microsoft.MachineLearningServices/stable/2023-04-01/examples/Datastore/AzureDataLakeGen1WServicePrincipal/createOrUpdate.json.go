package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewDatastore(ctx, "datastore", &machinelearningservices.DatastoreArgs{
			DatastoreProperties: machinelearningservices.AzureDataLakeGen1Datastore{
				Credentials: machinelearningservices.ServicePrincipalDatastoreCredentials{
					AuthorityUrl:    "string",
					ClientId:        "00000000-1111-2222-3333-444444444444",
					CredentialsType: "ServicePrincipal",
					ResourceUrl:     "string",
					Secrets: machinelearningservices.ServicePrincipalDatastoreSecrets{
						ClientSecret: "string",
						SecretsType:  "ServicePrincipal",
					},
					TenantId: "00000000-1111-2222-3333-444444444444",
				},
				DatastoreType: "AzureDataLakeGen1",
				Description:   "string",
				StoreName:     "string",
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
