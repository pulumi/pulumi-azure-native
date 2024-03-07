package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/communication/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := communication.NewEmailService(ctx, "emailService", &communication.EmailServiceArgs{
			DataLocation:      pulumi.String("United States"),
			EmailServiceName:  pulumi.String("MyEmailServiceResource"),
			Location:          pulumi.String("Global"),
			ResourceGroupName: pulumi.String("MyResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
