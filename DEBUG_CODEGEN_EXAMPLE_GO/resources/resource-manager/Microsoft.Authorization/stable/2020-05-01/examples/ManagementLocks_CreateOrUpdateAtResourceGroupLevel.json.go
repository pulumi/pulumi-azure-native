package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewManagementLockAtResourceGroupLevel(ctx, "managementLockAtResourceGroupLevel", &authorization.ManagementLockAtResourceGroupLevelArgs{
			Level:             pulumi.String("ReadOnly"),
			LockName:          pulumi.String("testlock"),
			ResourceGroupName: pulumi.String("resourcegroupname"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
