package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewConnectionGateway(ctx, "connectionGateway", &web.ConnectionGatewayArgs{
			ConnectionGatewayName: pulumi.String("test123"),
			Properties: &web.ConnectionGatewayDefinitionPropertiesArgs{
				BackendUri: pulumi.String("https://WABI-WEST-US-redirect.analysis.windows.net"),
				ConnectionGatewayInstallation: &web.ConnectionGatewayReferenceArgs{
					Id: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/locations/westus/connectionGatewayInstallations/865dccd1-5d5c-45fe-b5a0-249d4de4134c"),
				},
				ContactInformation: pulumi.StringArray{
					pulumi.String("test123@microsoft.com"),
				},
				DisplayName: pulumi.String("test123"),
				MachineName: pulumi.String("TEST123"),
				Status:      pulumi.Any("Installed"),
			},
			ResourceGroupName: pulumi.String("testResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
