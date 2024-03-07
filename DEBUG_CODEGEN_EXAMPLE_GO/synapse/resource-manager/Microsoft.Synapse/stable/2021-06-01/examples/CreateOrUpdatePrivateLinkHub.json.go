package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewPrivateLinkHub(ctx, "privateLinkHub", &synapse.PrivateLinkHubArgs{
			Location:           pulumi.String("East US"),
			PrivateLinkHubName: pulumi.String("privateLinkHub1"),
			ResourceGroupName:  pulumi.String("resourceGroup1"),
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
