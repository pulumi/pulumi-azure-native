package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/management/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := management.NewManagementGroupSubscription(ctx, "managementGroupSubscription", &management.ManagementGroupSubscriptionArgs{
			GroupId: pulumi.String("Group"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
