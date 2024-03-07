package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewManagedPrivateEndpoint(ctx, "managedPrivateEndpoint", &kusto.ManagedPrivateEndpointArgs{
			ClusterName:                pulumi.String("kustoCluster"),
			GroupId:                    pulumi.String("blob"),
			ManagedPrivateEndpointName: pulumi.String("managedPrivateEndpointTest"),
			PrivateLinkResourceId:      pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/storageAccountTest"),
			RequestMessage:             pulumi.String("Please Approve."),
			ResourceGroupName:          pulumi.String("kustorptest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
