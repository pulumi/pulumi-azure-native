package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/search/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := search.NewSharedPrivateLinkResource(ctx, "sharedPrivateLinkResource", &search.SharedPrivateLinkResourceArgs{
			Properties: &search.SharedPrivateLinkResourcePropertiesArgs{
				GroupId:               pulumi.String("blob"),
				PrivateLinkResourceId: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storageAccountName"),
				RequestMessage:        pulumi.String("please approve"),
			},
			ResourceGroupName:             pulumi.String("rg1"),
			SearchServiceName:             pulumi.String("mysearchservice"),
			SharedPrivateLinkResourceName: pulumi.String("testResource"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
