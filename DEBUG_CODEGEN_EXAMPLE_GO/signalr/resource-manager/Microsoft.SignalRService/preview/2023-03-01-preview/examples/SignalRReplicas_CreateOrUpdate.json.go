package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/signalrservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := signalrservice.NewSignalRReplica(ctx, "signalRReplica", &signalrservice.SignalRReplicaArgs{
			Location:          pulumi.String("eastus"),
			ReplicaName:       pulumi.String("mySignalRService-eastus"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ResourceName:      pulumi.String("mySignalRService"),
			Sku: &signalrservice.ResourceSkuArgs{
				Capacity: pulumi.Int(1),
				Name:     pulumi.String("Premium_P1"),
				Tier:     pulumi.String("Premium"),
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
