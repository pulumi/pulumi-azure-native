package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewSubscription(ctx, "subscription", &apimanagement.SubscriptionArgs{
			DisplayName:       pulumi.String("testsub"),
			OwnerId:           pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/57127d485157a511ace86ae7"),
			ResourceGroupName: pulumi.String("rg1"),
			Scope:             pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5600b59475ff190048060002"),
			ServiceName:       pulumi.String("apimService1"),
			Sid:               pulumi.String("testsub"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
