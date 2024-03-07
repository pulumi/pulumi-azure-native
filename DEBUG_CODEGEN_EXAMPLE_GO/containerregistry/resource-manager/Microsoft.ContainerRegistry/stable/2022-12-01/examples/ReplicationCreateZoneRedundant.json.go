package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewReplication(ctx, "replication", &containerregistry.ReplicationArgs{
			Location:              pulumi.String("eastus"),
			RegionEndpointEnabled: pulumi.Bool(true),
			RegistryName:          pulumi.String("myRegistry"),
			ReplicationName:       pulumi.String("myReplication"),
			ResourceGroupName:     pulumi.String("myResourceGroup"),
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
			ZoneRedundancy: pulumi.String("Enabled"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
