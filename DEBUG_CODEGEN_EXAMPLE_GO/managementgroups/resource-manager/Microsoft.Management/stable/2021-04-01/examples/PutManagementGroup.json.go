package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/management/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := management.NewManagementGroup(ctx, "managementGroup", &management.ManagementGroupArgs{
			Details: &management.CreateManagementGroupDetailsArgs{
				Parent: &management.CreateParentGroupInfoArgs{
					Id: pulumi.String("/providers/Microsoft.Management/managementGroups/RootGroup"),
				},
			},
			DisplayName: pulumi.String("ChildGroup"),
			GroupId:     pulumi.String("ChildGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
