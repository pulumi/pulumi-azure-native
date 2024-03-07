package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewConnector(ctx, "connector", &customerinsights.ConnectorArgs{
			ConnectorName: pulumi.String("testConnector"),
			ConnectorProperties: pulumi.Map{
				"connectionKeyVaultUrl": pulumi.Any{
					OrganizationId:  "XXX",
					OrganizationUrl: "https://XXX.crmlivetie.com/",
				},
			},
			ConnectorType:     pulumi.String("AzureBlob"),
			Description:       pulumi.String("Test connector"),
			DisplayName:       pulumi.String("testConnector"),
			HubName:           pulumi.String("sdkTestHub"),
			ResourceGroupName: pulumi.String("TestHubRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
