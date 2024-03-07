package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appconfiguration.NewReplica(ctx, "replica", &appconfiguration.ReplicaArgs{
			ConfigStoreName:   pulumi.String("contoso"),
			Location:          pulumi.String("eastus"),
			ReplicaName:       pulumi.String("myReplicaEus"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
