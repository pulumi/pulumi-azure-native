package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/webpubsub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := webpubsub.NewWebPubSubReplica(ctx, "webPubSubReplica", &webpubsub.WebPubSubReplicaArgs{
			Location:          pulumi.String("eastus"),
			ReplicaName:       pulumi.String("myWebPubSubService-eastus"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ResourceName:      pulumi.String("myWebPubSubService"),
			Sku: &webpubsub.ResourceSkuArgs{
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
