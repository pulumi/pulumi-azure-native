package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/management/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := management.NewHierarchySetting(ctx, "hierarchySetting", &management.HierarchySettingArgs{
			DefaultManagementGroup:               pulumi.String("/providers/Microsoft.Management/managementGroups/DefaultGroup"),
			GroupId:                              pulumi.String("root"),
			RequireAuthorizationForGroupCreation: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
