package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewConnectedEnvironment(ctx, "connectedEnvironment", &app.ConnectedEnvironmentArgs{
			ConnectedEnvironmentName: pulumi.String("testenv"),
			CustomDomainConfiguration: &app.CustomDomainConfigurationArgs{
				CertificatePassword: pulumi.String("private key password"),
				CertificateValue:    pulumi.String("Y2VydA=="),
				DnsSuffix:           pulumi.String("www.my-name.com"),
			},
			DaprAIConnectionString: pulumi.String("InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/"),
			ExtendedLocation: &app.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:          pulumi.String("East US"),
			ResourceGroupName: pulumi.String("examplerg"),
			StaticIp:          pulumi.String("1.2.3.4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
