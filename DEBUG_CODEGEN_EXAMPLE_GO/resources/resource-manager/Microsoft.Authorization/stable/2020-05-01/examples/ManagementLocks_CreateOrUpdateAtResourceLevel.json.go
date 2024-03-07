package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewManagementLockAtResourceLevel(ctx, "managementLockAtResourceLevel", &authorization.ManagementLockAtResourceLevelArgs{
			Level:                     pulumi.String("ReadOnly"),
			LockName:                  pulumi.String("testlock"),
			ParentResourcePath:        pulumi.String("parentResourcePath"),
			ResourceGroupName:         pulumi.String("resourcegroupname"),
			ResourceName:              pulumi.String("teststorageaccount"),
			ResourceProviderNamespace: pulumi.String("Microsoft.Storage"),
			ResourceType:              pulumi.String("storageAccounts"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
