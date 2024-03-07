package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicelinker/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicelinker.NewLinkerDryrun(ctx, "linkerDryrun", &servicelinker.LinkerDryrunArgs{
			DryrunName: pulumi.String("dryrunName"),
			Parameters: &servicelinker.CreateOrUpdateDryrunParametersArgs{
				ActionName: pulumi.String("createOrUpdate"),
				AuthInfo: servicelinker.SecretAuthInfo{
					AuthType: "secret",
					Name:     "name",
					SecretInfo: servicelinker.ValueSecretInfo{
						SecretType: "rawValue",
						Value:      "secret",
					},
				},
				TargetService: servicelinker.AzureResource{
					Id:   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db",
					Type: "AzureResource",
				},
			},
			ResourceUri: pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
