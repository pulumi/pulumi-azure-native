package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dashboard/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dashboard.NewManagedPrivateEndpoint(ctx, "managedPrivateEndpoint", &dashboard.ManagedPrivateEndpointArgs{
			GroupIds: pulumi.StringArray{
				pulumi.String("grafana"),
			},
			Location:                   pulumi.String("West US"),
			ManagedPrivateEndpointName: pulumi.String("myMPEName"),
			PrivateLinkResourceId:      pulumi.String("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-000000000000/resourceGroups/xx-rg/providers/Microsoft.Kusto/Clusters/sampleKustoResource"),
			PrivateLinkResourceRegion:  pulumi.String("West US"),
			PrivateLinkServiceUrl:      pulumi.String("my-self-hosted-influxdb.westus.mydomain.com"),
			RequestMessage:             pulumi.String("Example Request Message"),
			ResourceGroupName:          pulumi.String("myResourceGroup"),
			WorkspaceName:              pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
