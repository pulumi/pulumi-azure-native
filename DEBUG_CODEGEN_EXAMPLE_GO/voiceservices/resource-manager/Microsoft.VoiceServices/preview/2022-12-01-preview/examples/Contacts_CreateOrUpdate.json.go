package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/voiceservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := voiceservices.NewContact(ctx, "contact", &voiceservices.ContactArgs{
			CommunicationsGatewayName: pulumi.String("myname"),
			ContactName:               pulumi.String("John Smith"),
			Email:                     pulumi.String("johnsmith@example.com"),
			Location:                  pulumi.String("useast"),
			PhoneNumber:               pulumi.String("+1-555-1234"),
			ResourceGroupName:         pulumi.String("testrg"),
			Role:                      pulumi.String("Network Manager"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
