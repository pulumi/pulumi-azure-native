package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databasewatcher/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databasewatcher.NewSharedPrivateLinkResource(ctx, "sharedPrivateLinkResource", &databasewatcher.SharedPrivateLinkResourceArgs{
			DnsZone:                       pulumi.String("ec3ae9d410ba"),
			GroupId:                       pulumi.String("vault"),
			PrivateLinkResourceId:         pulumi.String("/subscriptions/6f53185c-ea09-4fc3-9075-318dec805303/resourceGroups/apiTest-ddat4p/providers/Microsoft.KeyVault/vaults/kvmo3ej9ih"),
			RequestMessage:                pulumi.String("request message"),
			ResourceGroupName:             pulumi.String("apiTest-ddat4p"),
			SharedPrivateLinkResourceName: pulumi.String("monitoringh22eed"),
			WatcherName:                   pulumi.String("databasemo3ej9ih"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
