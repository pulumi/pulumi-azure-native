package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewServerTrustGroup(ctx, "serverTrustGroup", &sql.ServerTrustGroupArgs{
			GroupMembers: []sql.ServerInfoArgs{
				{
					ServerId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/managedInstance-1"),
				},
				{
					ServerId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/managedInstance-2"),
				},
			},
			LocationName:         pulumi.String("Japan East"),
			ResourceGroupName:    pulumi.String("Default"),
			ServerTrustGroupName: pulumi.String("server-trust-group-test"),
			TrustScopes: pulumi.StringArray{
				pulumi.String("GlobalTransactions"),
				pulumi.String("ServiceBroker"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
