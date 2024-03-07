package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewGroup(ctx, "group", &migrate.GroupArgs{
			ETag:              pulumi.String("\"1e000c2c-0000-0d00-0000-5cdaa4190000\""),
			GroupName:         pulumi.String("Group2"),
			ProjectName:       pulumi.String("abgoyalWEselfhostb72bproject"),
			Properties:        nil,
			ResourceGroupName: pulumi.String("abgoyal-westEurope"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
