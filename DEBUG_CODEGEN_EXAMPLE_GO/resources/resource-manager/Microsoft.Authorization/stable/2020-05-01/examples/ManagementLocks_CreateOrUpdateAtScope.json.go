package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewManagementLockByScope(ctx, "managementLockByScope", &authorization.ManagementLockByScopeArgs{
			Level:    pulumi.String("ReadOnly"),
			LockName: pulumi.String("testlock"),
			Scope:    pulumi.String("subscriptions/subscriptionId"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
