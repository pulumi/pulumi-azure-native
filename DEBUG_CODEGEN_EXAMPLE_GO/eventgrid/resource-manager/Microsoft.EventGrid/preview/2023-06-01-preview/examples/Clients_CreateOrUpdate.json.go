package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewClient(ctx, "client", &eventgrid.ClientArgs{
			Attributes: pulumi.Any{
				DeviceTypes: []string{
					"Fan",
					"Light",
					"AC",
				},
				Floor: 3,
				Room:  "345",
			},
			ClientCertificateAuthentication: &eventgrid.ClientCertificateAuthenticationArgs{
				ValidationScheme: pulumi.String("SubjectMatchesAuthenticationName"),
			},
			ClientName:        pulumi.String("exampleClientName1"),
			Description:       pulumi.String("This is a test client"),
			NamespaceName:     pulumi.String("exampleNamespaceName1"),
			ResourceGroupName: pulumi.String("examplerg"),
			State:             pulumi.String("Enabled"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
