package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewQueue(ctx, "queue", &storage.QueueArgs{
			AccountName: pulumi.String("sto328"),
			Metadata: pulumi.StringMap{
				"sample1": pulumi.String("meta1"),
				"sample2": pulumi.String("meta2"),
			},
			QueueName:         pulumi.String("queue6185"),
			ResourceGroupName: pulumi.String("res3376"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
