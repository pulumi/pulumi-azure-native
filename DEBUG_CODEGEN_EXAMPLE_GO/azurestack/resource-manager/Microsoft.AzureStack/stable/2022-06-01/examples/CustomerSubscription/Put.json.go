package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestack/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestack.NewCustomerSubscription(ctx, "customerSubscription", &azurestack.CustomerSubscriptionArgs{
			CustomerSubscriptionName: pulumi.String("E09A4E93-29A7-4EBA-A6D4-76202383F07F"),
			RegistrationName:         pulumi.String("testregistration"),
			ResourceGroup:            pulumi.String("azurestack"),
			TenantId:                 pulumi.String("dbab3982-796f-4d03-9908-044c08aef8a2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
