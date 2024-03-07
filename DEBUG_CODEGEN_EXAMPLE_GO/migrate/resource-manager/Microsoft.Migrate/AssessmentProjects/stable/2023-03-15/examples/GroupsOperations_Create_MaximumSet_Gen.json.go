package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewGroupsOperation(ctx, "groupsOperation", &migrate.GroupsOperationArgs{
			GroupName:         pulumi.String("kuchatur-test"),
			GroupType:         pulumi.String("Default"),
			ProjectName:       pulumi.String("app18700project"),
			ProvisioningState: pulumi.String("Succeeded"),
			ResourceGroupName: pulumi.String("ayagrawrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
