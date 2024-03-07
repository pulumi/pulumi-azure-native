package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/communication/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := communication.NewSenderUsername(ctx, "senderUsername", &communication.SenderUsernameArgs{
			DisplayName:       pulumi.String("Contoso News Alerts"),
			DomainName:        pulumi.String("contoso.com"),
			EmailServiceName:  pulumi.String("contosoEmailService"),
			ResourceGroupName: pulumi.String("contosoResourceGroup"),
			SenderUsername:    pulumi.String("contosoNewsAlerts"),
			Username:          pulumi.String("contosoNewsAlerts"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
