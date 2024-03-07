package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicelinker/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicelinker.NewLinker(ctx, "linker", &servicelinker.LinkerArgs{
			AuthInfo: servicelinker.SecretAuthInfo{
				AuthType: "secret",
				Name:     "name",
				SecretInfo: servicelinker.ValueSecretInfo{
					SecretType: "rawValue",
					Value:      "secret",
				},
			},
			LinkerName:  pulumi.String("linkName"),
			ResourceUri: pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app"),
			TargetService: servicelinker.AzureResource{
				Id:   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DBforPostgreSQL/servers/test-pg/databases/test-db",
				Type: "AzureResource",
			},
			VNetSolution: &servicelinker.VNetSolutionArgs{
				Type: pulumi.String("serviceEndpoint"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
