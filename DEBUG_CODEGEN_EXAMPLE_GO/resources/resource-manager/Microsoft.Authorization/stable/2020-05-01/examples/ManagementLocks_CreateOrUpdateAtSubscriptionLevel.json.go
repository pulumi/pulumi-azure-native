package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewManagementLockAtSubscriptionLevel(ctx, "managementLockAtSubscriptionLevel", &authorization.ManagementLockAtSubscriptionLevelArgs{
			Level:    pulumi.String("ReadOnly"),
			LockName: pulumi.String("testlock"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
