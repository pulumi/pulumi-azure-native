package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewLinkedService(ctx, "linkedService", &operationalinsights.LinkedServiceArgs{
			LinkedServiceName:     pulumi.String("Cluster"),
			ResourceGroupName:     pulumi.String("mms-eus"),
			WorkspaceName:         pulumi.String("TestLinkWS"),
			WriteAccessResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/clusters/testcluster"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
