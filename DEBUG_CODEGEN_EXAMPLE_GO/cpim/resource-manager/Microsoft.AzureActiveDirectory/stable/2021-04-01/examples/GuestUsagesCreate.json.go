package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azureactivedirectory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azureactivedirectory.NewGuestUsage(ctx, "guestUsage", &azureactivedirectory.GuestUsageArgs{
			ResourceGroupName: pulumi.String("contosoResourceGroup"),
			ResourceName:      pulumi.String("contoso.onmicrosoft.com"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
